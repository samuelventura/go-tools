package tools

import (
	"os"
	"os/signal"
)

func SetupCtrlc() chan os.Signal {
	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)
	return ctrlc
}
