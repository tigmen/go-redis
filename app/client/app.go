package main

import (
	"context"
	"log"
	"net"
)

func main() {
	var d net.Dialer

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello, World"))
	if err != nil {
		log.Fatalln(err)
	}
}
