package tools

import (
	"fmt"
	"log"
)

type argsDso struct {
	values map[string]interface{}
}

type Args interface {
	Clone() Args
	SetValue(name string, value interface{})
	GetValue(name string) interface{}
}

func NewArgs() Args {
	dso := &argsDso{}
	dso.values = make(map[string]interface{})
	return dso
}

func (dso *argsDso) Clone() Args {
	ndso := &argsDso{}
	ndso.values = make(map[string]interface{})
	for k, v := range dso.values {
		ndso.values[k] = v
	}
	return ndso
}

func (dso *argsDso) GetValue(name string) interface{} {
	value, ok := dso.values[name]
	if !ok {
		err := fmt.Errorf("value not found %s", name)
		log.Panic(err)
	}
	return value
}

func (dso *argsDso) SetValue(name string, value interface{}) {
	dso.values[name] = value
}
