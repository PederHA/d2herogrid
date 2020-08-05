package utils

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func MarshalJSON(path string, i interface{}) error {
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0644)
}

// I'm not sure how robust this is? Param i MUST be a pointer.
func UnmarshalJSON(path string, i interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}

func MarshalYAML(path string, i interface{}) error {
	b, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0644)
}

// I'm not sure how robust this is? Param i MUST be a pointer.
func UnmarshalYAML(path string, i interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, i)
}
