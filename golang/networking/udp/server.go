package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		panic(err)
	}

	fmt.Print("listen udp", laddr)
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		n, raddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Print("err1 ", err, raddr)
			continue
		}
		fmt.Print("recv ", raddr.String(), string(data[:n]))

	}
}
