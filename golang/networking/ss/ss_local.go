package main

import (
	"bufio"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rc4"
	"io"
	"log"
	"net"
)

type config struct {
	Local    string
	Server   string
	Password string
}

var cfg *config

func main() {
	cfg = &config{
		Local:    ":8081",
		Server:   "node1.us.to:8082",
		Password: "123456",
	}
	lis, err := net.Listen("tcp", cfg.Local)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	log.Print("local listening at ", cfg.Local)

	for {
		lconn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go handle_conn(lconn)
	}
}

func handle_conn(lconn net.Conn) {
	defer lconn.Close()
	log.Print("handle_conn ", lconn)
	rconn, err := net.Dial("tcp", cfg.Server)
	if err != nil {
		panic(err)
	}
	defer rconn.Close()
	go io.Copy(rconn, lconn)
	io.Copy(lconn, rconn)
	c := Conn{}
	c.Read()

}
