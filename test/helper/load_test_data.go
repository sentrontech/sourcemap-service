package helper

import (
	"io/ioutil"
	"path/filepath"
)

// LoadTestData loads data for tests
func LoadTestData(testDataPath string) []byte {
	path, err := filepath.Abs("../../../test/data/" + testDataPath)
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}
