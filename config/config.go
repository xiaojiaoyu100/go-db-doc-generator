package config

import (
	"io/ioutil"
)

type (
	PgConfig         Config
	MongoDBConfig    Config
	TableStoreConfig Config
	Config           struct {
		FileScanDir  string `json:"file_scan_dir"`  // 文件夹扫描路径
		ScanItemName string `json:"scan_item_name"` // 需要扫描的文件名
		DestDir      string `json:"dest_dir"`       // 最终保存的文件路径（相对路径）
	}
)

func (jst *JsonStruct) Load(filename string, v interface{}) ([]byte, error) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
