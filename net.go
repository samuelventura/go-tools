package tools

import (
	"net"
	"time"

	"github.com/felixge/tcpkeepalive"
)

func KeepAlive(conn net.Conn, seconds int64) {
	count := 3
	interval := time.Duration(1) * time.Second
	idleTime := time.Duration(seconds) * time.Second
	err := tcpkeepalive.SetKeepAlive(
		conn, idleTime, count, interval)
	PanicIfError(err)
}
