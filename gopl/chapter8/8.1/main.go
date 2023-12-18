package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// Parse the port from the command line
	portPtr := flag.String("port", "8000", "Port")
	flag.Parse()
	port := fmt.Sprintf("localhost:%v", *portPtr)

	// Parse the Time Zone from the environment variable
	timezone := os.Getenv("TZ")

	if timezone == "" {
		log.Fatalln("Invalid timezone")
	}

	startServer(port, timezone)
}

func startServer(port string, timezone string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handleConn(conn, timezone)
	}
}

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	for {

		loc, err := time.LoadLocation(timezone)
		if err != nil {
			fmt.Println("Error loading location:", err)
			return
		}

		now := fmt.Sprintf("%s\n", time.Now().In(loc).Format(time.UnixDate))

		_, err = io.WriteString(c, now)
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
