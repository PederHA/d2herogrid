package utils

import (
	"encoding/json"
	"io/ioutil"
)

func MarshalJSON(path string, i interface{}) error {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
