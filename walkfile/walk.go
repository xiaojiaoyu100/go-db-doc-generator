package walkfile

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	model string
)

var listFile []string //获取文件列表

func ListFunc(path string, f os.FileInfo, err error) error {
	var strRet string

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path //+ "\r\n"

	//用strings.HasSuffix(src, suffix)//判断src中是否包含以suffix结尾的文件
	ok := strings.HasSuffix(strRet, model)
	if ok {
		listFile = append(listFile, strRet)
	}
	return nil
}

func GetFileList(scanModel, path string) ([]string, error) {
	listFile = nil
	model = scanModel
	err := filepath.Walk(path, ListFunc)
	if err != nil {
		return nil, err
	}
	return listFile, nil
}
