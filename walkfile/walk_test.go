package walkfile

import (
	"testing"

	"gitlab.xinghuolive.com/Backend-Go/StructParser/config"
)

func TestParse(t *testing.T) {
	_ = parsePackageFromDir(config.PgFileScanDir)
}
