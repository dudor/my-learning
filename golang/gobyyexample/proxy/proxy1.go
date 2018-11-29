package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	err := handShake(r, conn)
	if err != nil {
		log.Println(err)
	}
	addr, err := readAddress(r)
	if err != nil {
		log.Println(err)
	}
	log.Print("得到的完整的地址是：", addr) //注意：HTTP对应的是80端口，HTTPS对应的是443端口。

	remote, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	//如果成功连接远程主机，给客户端发送如下通知
	resp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	conn.Write(resp)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		io.Copy(remote, r)
		remote.Close()
		log.Print("copy r to remote")
	}()

	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
		log.Print("copy remote to conn")
	}()
	wg.Wait()
	log.Print("handler end")

}

func handShake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Println("版本号是", version)
	if version != 5 {
		return errors.New("该协议不是socks5")
	}
	nmethods, _ := r.ReadByte()
	log.Println("methods部分的长度是", nmethods)
	methods := make([]byte, nmethods)
	io.ReadFull(r, methods)
	log.Println("methods是", methods)
	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}
func readAddress(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("客户端协议版本：%d", version)
	if version != 5 {
		return "", errors.New("该协议不是socks5协议")
	}
	cmd, _ := r.ReadByte()
	log.Printf("客户端请求的类型是：%d", cmd)
	if cmd != 1 {
		return "", errors.New("客户端请求类型不为“1”，即请求类型必须是代理连接！.")
	}
	rsv, _ := r.ReadByte()
	log.Printf("客户端请求的RSV是：%d", rsv)
	atyp, _ := r.ReadByte()
	log.Printf("客户端请求的ATYP是：%d", atyp)

	var addr_len byte
	if atyp == 1 { //Ip4地址
		addr_len = 4

	} else if atyp == 3 { //域名
		addr_len, _ = r.ReadByte()
	} else { //ipv6
		addr_len = 16
	}

	addr := make([]byte, addr_len)
	io.ReadFull(r, addr)
	addr_str := ""
	if atyp == 1 {
		addr_str = fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
	} else if atyp == 3 {
		addr_str = string(addr)

	} else {
		log.Print("不支持	IPV6")
	}
	log.Printf("客户端请求的ADDR是:%s", addr_str)

	var port int16
	binary.Read(r, binary.BigEndian, &port)
	log.Printf("客户端请求的PORT是:%d", port)
	return fmt.Sprintf("%s:%d", addr_str, port), nil

}
