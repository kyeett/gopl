package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Usage:
//  go run netcat2.go
//  nc localhost 8000

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// Catch ctrl-c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	var i int
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		i++
		go handleConn(conn, i)
	}
}

func cleanup() {
	fmt.Println("\nUser interrupt. Shut down connections...")
}

func handleConn(c net.Conn, i int) {
	defer c.Close()
	fmt.Printf("Connection #%d accepted\n", i)
	defer fmt.Printf("Connection #%d shutdown\n", i)

	for {
		_, err := io.WriteString(c, time.Now().Format(time.StampMilli)+"\n")
		if err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
