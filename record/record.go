package record

import (
	"fmt"
	"os"

	"github.com/Howie59/go-db-doc-generator/parser"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Record2MarkdownFile(path string, schema *parser.Schema) {
	if schema.TableName == "" {
		return
	}
	err := createFile(path)
	if err != nil {
		logger.Error("create file failed ", zap.Error(err))
		return
	}
	outputFileName := path + schema.TableName + ".md"

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

func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
