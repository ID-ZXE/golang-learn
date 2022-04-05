package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// nc localhost 8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// goroutine是Go语言程序的并发体
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:01:01\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
