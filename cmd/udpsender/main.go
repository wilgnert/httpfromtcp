package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := "localhost:42069"
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v\n", err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Failed to dial udp connection: %v\n", err)
	}
	defer conn.Close()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := r.ReadString('\n')
		if err != nil {
			log.Printf("error reading line: %v\n", err)
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("error writing through connection: %v\n", err)
		}
	}

}