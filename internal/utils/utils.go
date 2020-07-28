package utils

import (
	"fmt"
	"os"
)

// MaxInt returns the largest integer in a slice of integers
func MaxInt(s []int) (int, error) {
	var n int

	if len(s) == 0 {
		return 0, fmt.Errorf("utils.Max: slice cannot be empty")
	}
	for _, i := range s {
		if n == 0 || i > n {
			n = i
		}
	}
	return n, nil
}

func CheckFileExists(filepath string) (bool, error) { // Should this return (bool, error)?
	stat, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	} else if stat.IsDir() {
		return false, fmt.Errorf("'%s' is a directory", filepath)
	}
	return true, nil
}

func CheckDirExists(dirpath string) (bool, error) {
	stat, err := os.Stat(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !stat.IsDir() {
		return false, fmt.Errorf("'%s' is not a directory", dirpath)
	}
	return true, nil
}
