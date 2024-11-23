package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: go.dev\n"
	fmt.Println("Request:\n", httpRequest)
	conn, err := net.Dial("tcp", "go.dev:443")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
