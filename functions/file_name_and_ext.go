package functions

import (
	"path/filepath"
)

func FileNameAndExt(fileName string) (string, string) {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))], filepath.Ext(fileName)
}
