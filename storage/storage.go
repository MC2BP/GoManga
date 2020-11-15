package storage

import (
	"GoManga/collector"
	"fmt"
	"os"
)

func PrepareMangaFolder(basePath string, collectors map[string]collector.Collector) error {
	exists, err := ExistsFolder(basePath)
	if !exists {
		err = CreateFolder(basePath)
		if err != nil {
			return err
		}
	}
	for source, _ := range collectors {
		path := fmt.Sprint(basePath, "/", source)
		exists, err := ExistsFolder(path)
		if !exists {
			err = CreateFolder(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateFolder(path string) error {
	return os.Mkdir(path, 0755)
}

func ExistsFolder(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
