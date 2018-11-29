package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 8083} // 注意端口必须固定
	dstAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 8082}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {

		r := bufio.NewReader(os.Stdin)
		msg, _ := r.ReadBytes('\n')
		conn.Write(msg)
	}
}
