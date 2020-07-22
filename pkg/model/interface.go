package model

type CFG interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
}
