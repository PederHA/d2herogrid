package cli

import (
	"fmt"
	"reflect"
)

type Config struct {
	Brackets      []int  `json:"brackets" yaml:"brackets"`
	DefaultName   string `json:"default_name" yaml:"default_name"`
	Layout        int    `json:"layout" yaml:"layout"`
	Path          string `json:"path" yaml:"path"`
	SortAscending bool   `json:"sort_ascending" yaml:"sort_ascending"`
}

func (c *Config) Get(key string) (interface{}, error) {
	val := c.getFieldValue(key)
	if val == reflect.ValueOf(nil) {
		return nil, fmt.Errorf("Config.Get: no key named '%s'", key)
	}
	switch val.Kind() {
	case reflect.String:
		return val.String(), nil
	case reflect.Int, reflect.Int32, reflect.Int64:
		return val.Int(), nil
	case reflect.Slice:
		return val.Interface(), nil
	default:
		return nil, fmt.Errorf("Config.Get: Unimplemented type")
	}
}

func (c *Config) Set(key string, val interface{}) error {
	err := c.setFieldValue(key, val)
	return err
}

func (c *Config) getFieldValue(n string) reflect.Value {
	val := reflect.ValueOf(c).Elem()
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//fmt.Printf("field %d:\n", i)
		//fmt.Printf("\tName: %s\n", field.Name)
		//fmt.Printf("\tType: %v\n", field.Type)
		//fmt.Printf("\tTag: %v\n", field.Tag)
		//fmt.Printf("\tOffset: %v\n", field.Offset)
		//fmt.Printf("\tIndex: %v\n", field.Index)
		//fmt.Printf("\tAnonymous: %v\n", field.Anonymous)
		if field.Name == n {
			return val.Field(i)
		}
		//fmt.Printf("field %d: %v\n", i, field)
	}
	return reflect.ValueOf(nil)
}

func (c *Config) setFieldValue(key string, val interface{}) error {
	field := c.getFieldValue(key)
	switch val.(type) {
	case string:
		if field.Kind() == reflect.String {
			field.SetString(val.(string))
		} else {
			return fmt.Errorf("Config.Set: invalid type for field. Should be %v, got string", field.Kind())
		}
	case int, int32, int64:
		if field.Kind() == reflect.Int {
			field.SetInt(val.(int64))
		} else {
			return fmt.Errorf("Config.Set: invalid type for field. Should be %v, got int", field.Kind())
		}
	case float32:
		if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
			field.SetFloat(val.(float64))
		} else {
			return fmt.Errorf("Config.Set: invalid type for field. Should be %v, got float", field.Kind())
		}
	case []int:
		if field.Kind() == reflect.Slice {
			//if len(val.([]int)) > 0 || reflect.ValueOf(val.([]int)[0]).Kind() == reflect.Int {}
			s := reflect.ValueOf(val)
			field.Set(s)
		} else {
			return fmt.Errorf("Config.Set: invalid type for field. Should be %v, got slice", field.Kind())
		}
	// Other types are unhandled
	default:
		return fmt.Errorf("Config.Set: type is unimplemented")
	}

	return nil
}
