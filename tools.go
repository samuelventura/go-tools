package tools

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetEnviron(name string, defval string) string {
	value := os.Getenv(name)
	trimmed := strings.TrimSpace(value)
	if len(trimmed) > 0 {
		log.Println(name, value)
		return value
	}
	log.Println(name, defval)
	return defval
}

func WithExtension(ext string) string {
	exe, err := os.Executable()
	PanicIfError(err)
	dir := filepath.Dir(exe)
	base := filepath.Base(exe)
	file := base + "." + ext
	return filepath.Join(dir, file)
}

func GetHostname() string {
	host, err := os.Hostname()
	PanicIfError(err)
	return host
}

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
