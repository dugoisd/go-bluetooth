package gen

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Mkdir(dirpath string) error {
	err := os.Mkdir(dirpath, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func ListFiles(dir string) []string {

	list := make([]string, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "mgmt-api.txt") {
			return nil
		}

		if !strings.HasSuffix(path, "-api.txt") {
			return nil
		}

		list = append(list, path)
		return nil
	})

	if err != nil {
		log.Errorf("Failed to list files: %s", err)
	}

	return list
}

func ReadFile(srcFile string) ([]byte, error) {
	file, err := os.Open(srcFile)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
