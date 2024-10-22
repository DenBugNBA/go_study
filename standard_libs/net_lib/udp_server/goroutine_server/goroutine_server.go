package main

import (
	"log"
	"net"
	"time"
)

func runServer(closeServer <-chan struct{}) error {
	serverConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 10001})
	if err != nil {
		log.Fatalf("err: %v\n", err)
		return err
	}

	go func() {
		defer log.Println("server closed")
		defer serverConn.Close()

		for {
			select {
			case <-closeServer:
				log.Println("Return")
				return
			default:
				log.Println("Try to listen")
				buf := make([]byte, 1024)
				serverConn.SetReadDeadline(time.Now().Add(3 * time.Second))
				n, _, err := serverConn.ReadFromUDP(buf)
				if err != nil {
					log.Printf("Nothing: %v\n", err)
				} else {
					log.Printf("Received: %s\n", string(buf[:n]))
				}
			}
		}
	}()

	return nil
}

func runClient() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 10001})
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	message := "Привет, сервер!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	ch := make(chan struct{})
	if err := runServer(ch); err != nil {
		return
	}

	log.Println("Server established")

	for i := 0; i < 5; i++ {
		go runClient()
	}

	time.Sleep(time.Second * 10)
	close(ch)
	time.Sleep(time.Second * 5)
}
