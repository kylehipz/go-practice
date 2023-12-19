package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	new_york := os.Getenv("NewYork")
	tokyo := os.Getenv("Tokyo")
	london := os.Getenv("London")

	go connect("New York", new_york)
	go connect("Tokyo", tokyo)
	connect("London", london)
}

func connect(location string, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		from_tcp, _ := reader.ReadString('\n')
		message := fmt.Sprintf("%s: %s", location, from_tcp)
		fmt.Print(message)
	}
}
