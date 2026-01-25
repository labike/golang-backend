package util

import "os"

func CreateDir(filePath string) error {
	if !IsExists(filePath) {
		err := os.Mkdir(filePath, os.ModePerm)
		return err
	}
	return nil
}

func IsExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
