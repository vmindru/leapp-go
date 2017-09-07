package main

import (
	"io"
	"net"
	"os"
)

func main() {
	stdoutSockPath := os.Getenv("LEAPP_ACTOR_STDOUT_SOCK")
	if len(os.Args) > 1 && os.Args[1] == "server" {
		if sock, err := net.ListenUnix("unix", stdoutSockPath); err == nil {
			defer sock.Close()
			io.Copy(os.Stdout, sock)
		}
	} else {
		if sock, err := net.Dial("unix", stdoutSockPath); err == nil {
			defer sock.Close()
			io.Copy(sock, os.Stdin)
		}
	}
}
