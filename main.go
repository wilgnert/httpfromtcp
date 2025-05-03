package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"net"
)



func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		os.Exit(1)
	}
	defer listener.Close()
	for {
		f, err := listener.Accept()
		if err != nil {
			os.Exit(1)
		}
		ch := getLinesChannel(f)
		for s := range ch {
			fmt.Fprintln(os.Stdout, s)
		}
		f.Close()
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		current_line := ""
		for {
			var s [8]byte
			_, err := f.Read(s[:])
			if err != nil { break }
			parts := strings.Split(fmt.Sprintf("%s", s), "\n")
			for i := range len(parts) - 1 {
				if parts[i] != "" {
					current_line = strings.Join([]string{current_line, parts[i]}, "")
				}
				ch <- current_line
				current_line = ""
			}
			current_line = strings.Join([]string{current_line, parts[len(parts) - 1]}, "")
		}
		ch <- current_line
		close(ch)
	}()
	return ch
}