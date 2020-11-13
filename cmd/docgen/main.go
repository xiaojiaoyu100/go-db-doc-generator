package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/xiaojiaoyu100/go-db-doc-generator/config"
	"github.com/xiaojiaoyu100/go-db-doc-generator/parser"
	"github.com/xiaojiaoyu100/go-db-doc-generator/record"
	"github.com/xiaojiaoyu100/go-db-doc-generator/walkfile"
	"go.uber.org/zap"
)

func init() {
	newCmd.Flags().StringVar(&conf, "config", "sss", `Set config`)
}

var (
	logger *zap.Logger
	conf   string // 记录指定的config的名称
	dir    string // 记录当前路径
	newCmd = &cobra.Command{
		Use:   "set",
		Short: "set config for function",
		Long:  `set config`,
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ = os.Getwd()
		},
	}
)

// InitLogger ...
func InitLogger() {
	logger, _ = zap.NewDevelopment()
}

type configStruct struct {
	Conf []config.Config `json:"config"`
}

func main() {
	err := newCmd.Execute()
	InitLogger()
	defer logger.Sync()
	startTime := time.Now()
	JsonParse := config.NewJsonStruct()
	confStruct := new(configStruct)
	// 从本地加载配置文件
	load, err := JsonParse.Load(dir+conf, &confStruct)
	if err != nil {
		logger.Error("load config failed, returned ", zap.Error(err))
		return
	}
	// 反序列化json文件
	err = json.Unmarshal(load, confStruct)
	if err != nil {
		logger.Error("json unmarshal failed, returned ", zap.Error(err))
		return
	}
	for index := range confStruct.Conf {
		PGListFile, err := walkfile.GetFileList(confStruct.Conf[index].ScanItemName, confStruct.Conf[index].FileScanDir)
		if err != nil {
			logger.Error("filepath.Walk() failed, returned ", zap.Error(err))
			return
		}
		for _, filePath := range PGListFile {
			// 进行parse
			parseStruct, err := parser.ParseStruct(filePath)
			if err != nil {
				logger.Error("Error while parsing struct: ", zap.Error(err))
				return
			}
			// 记录到Markdown中
			record.Record2MarkdownFile(dir+confStruct.Conf[index].DestDir, parseStruct)
		}
		PGListFile = nil
	}
	logger.Info("All exec success")
	logger.Info(fmt.Sprintf("Total time is %f", time.Since(startTime).Seconds()))
}
