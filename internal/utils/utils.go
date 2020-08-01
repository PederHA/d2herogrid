package utils

import (
	"fmt"
	"os"
)

func CheckFileExists(filepath string) error { // Should this return (bool, error)?
	stat, err := os.Stat(filepath)
	if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory", filepath)
	}
	return nil
}

func CheckDirExists(dirpath string) error {
	stat, err := os.Stat(dirpath)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("'%s' is not a directory", dirpath)
	}
	return nil
}
