package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func GetEnvironInt(name string, base int, bitSize int, defval int64) int64 {
	value := os.Getenv(name)
	trimmed := strings.TrimSpace(value)
	if len(trimmed) > 0 {
		log.Println(name, value)
		parsed, err := strconv.ParseInt(value, base, bitSize)
		PanicIfError(err)
		return parsed
	}
	log.Println(name, defval)
	return defval
}

func GetEnvironUint(name string, base int, bitSize int, defval uint64) uint64 {
	value := os.Getenv(name)
	trimmed := strings.TrimSpace(value)
	if len(trimmed) > 0 {
		log.Println(name, value)
		parsed, err := strconv.ParseUint(value, base, bitSize)
		PanicIfError(err)
		return parsed
	}
	log.Println(name, defval)
	return defval
}

func GetEnvironFloat(name string, bitSize int, defval float64) float64 {
	value := os.Getenv(name)
	trimmed := strings.TrimSpace(value)
	if len(trimmed) > 0 {
		log.Println(name, value)
		parsed, err := strconv.ParseFloat(value, bitSize)
		PanicIfError(err)
		return parsed
	}
	log.Println(name, defval)
	return defval
}

func GetEnvironBool(name string, defval bool) bool {
	value := os.Getenv(name)
	trimmed := strings.TrimSpace(value)
	if len(trimmed) > 0 {
		log.Println(name, value)
		parsed, err := strconv.ParseBool(value)
		PanicIfError(err)
		return parsed
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

func LoadEnviron(path string) {
	env := ReadEnviron(path)
	for n, line := range env {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			err := fmt.Errorf("invalid entry '%s' at %s:%d", line, path, n)
			log.Panic(err)
		}
		os.Setenv(parts[0], parts[1])
	}
}

func ReadDefaultEnviron() []string {
	path := WithExtension("env")
	return ReadEnviron(path)
}

func LoadDefaultEnviron() {
	path := WithExtension("env")
	LoadEnviron(path)
}
