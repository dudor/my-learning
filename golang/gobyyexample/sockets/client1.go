package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入用户名和密码")
	go Send(conn)
	//	input, err := reader.ReadString('\n')
	//	if err != nil {
	//		panic(err)
	//	}
	//	conn.Write([]byte(input))
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data[:n]))
	}
}
func Send(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(input))
	}

}
