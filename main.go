package main

import (
	"fmt"
	"time"

	"gitlab.xinghuolive.com/Backend-Go/StructParser/config"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/parser"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/record"
	"gitlab.xinghuolive.com/Backend-Go/StructParser/walkfile"
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewDevelopment()
}

func main() {
	InitLogger()
	defer logger.Sync()
	startTime := time.Now()

	for index, _ := range config.SourceConfig {
		PGListFile, err := walkfile.GetFileList(config.SourceConfig[index].ScanItemName, config.SourceConfig[index].FileScanDir)
		if err != nil {
			logger.Error("pg filepath.Walk() failed, returned ", zap.Error(err))
			return
		}
		for _, filePath := range PGListFile {
			// 进行parse
			parseStruct, err := parser.ParseStruct(filePath)
			if err != nil {
				logger.Error("Error while parsing pg router.go:", zap.Error(err))
				return
			}
			// 记录到Markdown中
			record.Record2MarkdownFile(config.SourceConfig[index].DestDir, parseStruct)
		}
		PGListFile = nil
	}
	logger.Info("all exec success")
	logger.Info(fmt.Sprintf("Total time is %f", time.Since(startTime).Seconds()))
}
