package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WithExtension(ext string) string {
	exe, err := os.Executable()
	PanicIfError(err)
	dir := filepath.Dir(exe)
	base := filepath.Base(exe)
	file := base + "." + ext
	return filepath.Join(dir, file)
}

func ChangeExtension(cpath string, next string) string {
	cext := filepath.Ext(cpath) //includes .
	npath := strings.TrimSuffix(cpath, cext)
	return fmt.Sprintf("%s.%s", npath, next)
}
