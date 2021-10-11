// 在Go语言中，每一个并发的执行单元叫作一个goroutine
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	// 可以简单地把goroutine类比作一个线程，如
	// f()    // call f(); wait for it to return
	// go f() // create a new goroutine that calls f(); don't wait

	// 时钟服务器
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
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
