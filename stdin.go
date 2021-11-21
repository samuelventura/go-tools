package tools

import (
	"bufio"
	"io/ioutil"
	"os"
)

func SetupStdinAll() chan interface{} {
	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		ioutil.ReadAll(os.Stdin)
	}()
	return stdin
}

func SetupStdinLine() chan interface{} {
	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	}()
	return stdin
}

func SetupStdinWord(word string) chan interface{} {
	stdin := make(chan interface{})
	go func() {
		defer close(stdin)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == word {
				return
			}
		}
	}()
	return stdin
}
