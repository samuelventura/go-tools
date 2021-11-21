package tools

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logWriter struct {
	pid int
}

func (w logWriter) Write(bytes []byte) (int, error) {
	ts := time.Now().Format("20060102T150405.000")
	line := fmt.Sprintf("%s %d %s", ts, w.pid, string(bytes))
	return fmt.Print(line)
}

func SetupLog() {
	w := &logWriter{}
	w.pid = os.Getpid()
	log.SetFlags(0)
	log.SetOutput(w)
}
