package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	fmt.Println("begin listen", ":8081")
	if err != nil {
		log.Print(err)
		return
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handle_tcp(conn)
	}
}
func handle_tcp(conn net.Conn) {
	fmt.Println("handle_tcp", conn)
	defer conn.Close()
	data := make([]byte, 2014)
	for {
		n, err := conn.Read(data)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Write(data[:n])
	}

}
