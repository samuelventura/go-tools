package tools

import (
	"fmt"
	"log"
)

type argmDso struct {
	values map[string]interface{}
}

type Args interface {
	SetValue(name string, value interface{})
	GetValue(name string) interface{}
}

type Argm interface {
	Clone() Argm
	SetValue(name string, value interface{})
	GetValue(name string) interface{}
}

func NewArgm() Argm {
	dso := &argmDso{}
	dso.values = make(map[string]interface{})
	return dso
}

func (dso *argmDso) Clone() Argm {
	ndso := &argmDso{}
	ndso.values = make(map[string]interface{})
	for k, v := range dso.values {
		ndso.values[k] = v
	}
	return ndso
}

func (dso *argmDso) GetValue(name string) interface{} {
	value, ok := dso.values[name]
	if !ok {
		err := fmt.Errorf("value not found %s", name)
		log.Panic(err)
	}
	return value
}

func (dso *argmDso) SetValue(name string, value interface{}) {
	dso.values[name] = value
}
