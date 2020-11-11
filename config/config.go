package config

type (
	PgConfig         Config
	MongoDBConfig    Config
	TableStoreConfig Config
	Config           struct {
		FileScanDir  string // 文件夹扫描路径
		ScanItemName string // 需要扫描的文件名
		DestDir      string // 最终保存的文件路径
	}
)

var SourceConfig = []Config{
	Config{
		"",
		"",
		"",
	},
}
