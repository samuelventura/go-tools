package tools

import (
	"os"
)

func GetHostname() string {
	host, err := os.Hostname()
	PanicIfError(err)
	return host
}
