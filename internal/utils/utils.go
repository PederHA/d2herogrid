package utils

import (
	"fmt"
	"os"
)

// Max returns the largest number in a slice of integers
func Max(s []int) (int, error) {
	var n int

	if len(s) == 0 {
		return 0, fmt.Errorf("utils.max: slice cannot be empty")
	}
	for _, i := range s {
		if n == 0 || i > n {
			n = i
		}
	}
	return n, nil
}

func CheckFileExists(filepath string) error { // Should this return (bool, error)?
	stat, err := os.Stat(filepath)
	if err != nil {
		return err
	} else if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory", filepath)
	}
	return nil
}
