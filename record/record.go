package record

import (
	"fmt"
	"os"
	"strings"

	"gitlab.xinghuolive.com/Backend-Go/StructParser/parser"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Record2MarkdownFile(path string, schema *parser.Schema) {
	if schema.TableName == "" {
		return
	}
	outputFileName := path + schema.TableName + ".md"
	pgName := strings.Split(schema.TableName, ".")[0]
	if pgName == "multi" || pgName == "public" || pgName == "common" {
		outputFileName = path + strings.Split(schema.TableName, ".")[0] + "\\" + schema.TableName + ".md"
	}

	mdFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Error("\033[31mcreate and open markdown file error \033[0m \n%v\n", zap.Error(err))
		return
	}
	// markdown format语法
	tableContent := "## 介绍" + "\n"
	tableContent += "#### 本文档是自动生成的，请勿手改" + "\n"
	tableContent += "\n"
	tableContent += "## " + schema.TableName + "表" + "\n"
	tableContent += "\n" +
		"| 命名 | 类型  | 注释 |\n" +
		"| :--: | :--: | :--: |\n"
	tableContent += ModelPlus(schema.ModelType)

	for _, valMap := range schema.FieldMap {
		tableContent += fmt.Sprintf(
			"| `%s` | `%s` | `%s` |\n",
			valMap.Name,
			valMap.Type,
			valMap.Comment,
		)
	}
	tableContent += "\n\n"
	mdFile.WriteString(tableContent)
	err = mdFile.Close()
	fmt.Printf("\033[32m record %s markdown finished ... \033[0m \n", schema.TableName)
}

func ModelPlus(modelType string) string {
	var addPlus string
	if modelType == "" {
		return addPlus
	}
	addPlus += fmt.Sprintf("| `Id` | int64 | 主键ID |\n")
	addPlus += fmt.Sprintf("| `CreatedAt` | Time | 创建时间 |\n")
	addPlus += fmt.Sprintf("| `UpdatedAt` | Time | 更新时间 |\n")
	switch {
	case modelType == "Model" || modelType == "ModelByID":
		addPlus += fmt.Sprintf("| `DeletedAt` | Time | 删除时间 |\n")
		addPlus += fmt.Sprintf("| `IsDelete`  | bool | 是否删除 |\n")
	default:
		break
	}
	return addPlus
}
