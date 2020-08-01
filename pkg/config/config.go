// I REALLY should've read the chapter on reflection
// in The Go Programming Language before writing this

package config

import (
	"fmt"
	"reflect"
)

// Config is a type intended to be embedded in other types to add Set/Get methods
// to retrieve and modify struct member fields
type Config struct{}

// Get retrieves the value of a specific struct field
func (c *Config) Get(key string) (interface{}, error) {
	val := c.getFieldValue(key)
	if val == reflect.ValueOf(nil) {
		return nil, fmt.Errorf("config.Config.Get: no key named '%s'", key)
	}
	switch val.Kind() {
	case reflect.String:
		return val.String(), nil
	case reflect.Int, reflect.Int32, reflect.Int64:
		return val.Int(), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Slice:
		return val.Interface(), nil
	default:
		return nil, fmt.Errorf("config.Config.Set: type %v is not implemented", val.Kind())
	}
}

// Set modifies the value of a specific struct field
func (c *Config) Set(key string, val interface{}) error {
	return c.setFieldValue(key, val)
}

func (c *Config) getFieldValue(n string) reflect.Value {
	val := reflect.ValueOf(c).Elem()
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == n {
			return val.Field(i)
		}
	}
	return reflect.ValueOf(nil)
}

func (c *Config) setFieldValue(key string, val interface{}) error {
	invalidType := "config.Config.Set: invalid type for field. Should be %v, got %s"

	field := c.getFieldValue(key)
	switch val.(type) {
	case string:
		if field.Kind() == reflect.String {
			field.SetString(val.(string))
		} else {
			return fmt.Errorf(invalidType, field.Kind(), "string")
		}
	case int, int32, int64:
		if field.Kind() == reflect.Int {
			field.SetInt(val.(int64))
		} else {
			return fmt.Errorf(invalidType, field.Kind(), "int")
		}
	case float32, float64:
		if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
			field.SetFloat(val.(float64))
		} else {
			return fmt.Errorf(invalidType, field.Kind(), "float")
		}
	case []int:
		if field.Kind() == reflect.Slice {
			//if len(val.([]int)) > 0 || reflect.ValueOf(val.([]int)[0]).Kind() == reflect.Int {}
			s := reflect.ValueOf(val)
			field.Set(s)
		} else {
			return fmt.Errorf(invalidType, field.Kind(), "[]int")
		}
	// Other types are unhandled
	default:
		return fmt.Errorf("config.Config.Set: type %v is not implemented", field.Kind())
	}

	return nil
}
