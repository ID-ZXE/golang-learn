package main

import (
	"io"
	"log"
	"net"
	"time"
)

// GOMAXPROCS 指定操作系统线程数量
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
		// goroutine的调度方式不需要进入内核的上下文，所以重新调度一个goroutine比调度 一个线程代价要低得多。
		// goroutine没有ID号
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
