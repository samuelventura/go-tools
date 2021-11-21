package tools

import (
	"bufio"
	"log"
	"os"
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

func ReadEnviron(path string) []string {
	lines := []string{}
	_, err := os.Stat(path)
	if err != nil {
		return lines
	}
	file, err := os.Open(path)
	PanicIfError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if len(trimmed) > 0 {
			lines = append(lines, trimmed)
		}
	}
	return lines
}
