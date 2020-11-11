package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gitlab.xinghuolive.com/Backend-Go/StructParser/config"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/parser"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/record"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/walkfile"
	"go.uber.org/zap"
)

var logger *zap.Logger

// InitLogger ...
func InitLogger() {
	logger, _ = zap.NewDevelopment()
}

type configStruct struct {
	Conf []config.Config `json:"config"`
}

func main() {
	InitLogger()
	defer logger.Sync()
	startTime := time.Now()
	JsonParse := config.NewJsonStruct()
	config := new(configStruct)
	// 从本地加载配置文件
	load, err := JsonParse.Load("../config.json", &config)
	if err != nil {
		logger.Error("load config failed, returned ", zap.Error(err))
		return
	}
	// 反序列化json文件
	err = json.Unmarshal(load, config)
	if err != nil {
		logger.Error("json unmarshal failed, returned ", zap.Error(err))
		return
	}
	for index := range config.Conf {
		PGListFile, err := walkfile.GetFileList(config.Conf[index].ScanItemName, config.Conf[index].FileScanDir)
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
			record.Record2MarkdownFile(config.Conf[index].DestDir, parseStruct)
		}
		PGListFile = nil
	}
	logger.Info("All exec success")
	logger.Info(fmt.Sprintf("Total time is %f", time.Since(startTime).Seconds()))
}
