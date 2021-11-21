package tools

import (
	"log"
)

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func TraceRecover() {
	r := recover()
	if r != nil {
		log.Println("recover", r)
	}
}
