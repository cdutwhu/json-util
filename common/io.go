package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	FailOnErr("%v", ioutil.WriteFile(filename, data, 0666))
}
