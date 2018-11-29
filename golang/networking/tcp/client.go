package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Print(err)
		return
	}
	defer conn.Close()
	fmt.Println("connect to ", ":8081")

	go sendMsg(conn)
	data := make([]byte, 2048)
	for {
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Print("data", string(data[:n]))
	}
}
func sendMsg(conn net.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		msg, err := r.ReadString('\n')
		if err != nil {
			fmt.Print("err2", err)
			return
		} else {
			fmt.Print("msg", msg)
			conn.Write([]byte(msg))
		}
	}
}
