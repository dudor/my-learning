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

type config struct {
	Local    string
	Server   string
	Password string
}

var cfg *config

func main() {
	cfg = &config{
		Server:   ":8082",
		Password: "123456",
	}
	lis, err := net.Listen("tcp", cfg.Server)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	log.Print("server listening at ", cfg.Server)
	for {
		lconn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go handle_conn(lconn)
	}
}

func handle_conn(conn net.Conn) {
	defer conn.Close()
	log.Print("handle_conn ", conn)
	reader := bufio.NewReader(conn)
	err := shakeHand(reader, conn)
	if err != nil {
		log.Print(err)
		return
	}
	addr, err := getAddress(reader)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("REMOTE ADDRESS IS ", addr)

	rconn, err := net.Dial("tcp", addr)
	if err != nil {
		conn.Write([]byte{0x05, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		log.Print(err)
		return
	}
	defer rconn.Close()

	// return a reply
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	log.Print("RETURN A REPLY")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(rconn, reader)
		log.Print("COPY reader TO rconn")
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, rconn)
		log.Print("COPY rconn TO conn")
	}()

	wg.Wait()
	log.Println("handle_conn() END")

}

func shakeHand(reader *bufio.Reader, conn net.Conn) error {
	VER, err := reader.ReadByte()
	if err != nil {
		log.Print(err)
		return err
	}
	if VER != 0x05 {
		log.Print("not socks5 protocol")
		return errors.New("not socks5 protocol")
	}
	log.Print("VER ", VER)
	NMETHODS, _ := reader.ReadByte()
	log.Print("NMETHODS ", NMETHODS)

	METHODS := make([]byte, NMETHODS)
	n, _ := reader.Read(METHODS)
	log.Print("METHODS ", METHODS)
	if n != 1 && METHODS[0] != 0x00 {
		log.Print("METHODS ONLY SUPPORT NO AUTHENTICATION REQUIRED")
		return errors.New("METHODS ONLY SUPPORT NO AUTHENTICATION REQUIRED")
	}
	conn.Write([]byte{5, 0})
	return nil
}

func getAddress(r *bufio.Reader) (string, error) {
	VER, _ := r.ReadByte()
	if VER != 0x05 {
		log.Print("not socks5 protocol")
		return "", errors.New("not socks5 protocol")
	}
	CMD, _ := r.ReadByte()
	if CMD != 0x01 {
		log.Print("CMD ONLY SUPPORT CONNECT")
		return "", errors.New("CMD ONLY SUPPORT CONNECT")
	}
	r.ReadByte()
	ATYP, _ := r.ReadByte()

	var ip_str string

	if ATYP == 0x01 {
		ip_byte := make([]byte, net.IPv4len)
		r.Read(ip_byte)
		ip_str = fmt.Sprintf("%d.%d.%d.%d", ip_byte[0], ip_byte[1], ip_byte[2], ip_byte[3])
	} else if ATYP == 0x03 {
		addr_len, _ := r.ReadByte()
		ip_byte := make([]byte, addr_len)
		n, _ := r.Read(ip_byte)
		log.Print("domain = ", string(ip_byte[:n]))
		ipaddr, err := net.ResolveIPAddr("tcp", string(ip_byte[:n]))
		if err != nil {
			log.Print(err)
			return "", errors.New("ResolveIPAddr err")
		}
		ip_str = ipaddr.IP.String()

	} else if ATYP == 0x04 {
		ip_byte := make([]byte, net.IPv6len)
		n, _ := r.Read(ip_byte)
		ip_str = net.IP(ip_byte[:n]).String()

	} else {
		log.Print("ATYP NOT SUPPORT ", ATYP)
		return "", errors.New("ATYP NOT SUPPORT ")
	}

	var port int16
	err := binary.Read(r, binary.BigEndian, &port)
	if err != nil {
		log.Print("PORT IS ERR")
		return "", errors.New("PORT IS ERR")
	}

	return fmt.Sprintf("%s:%d", ip_str, port), nil
}
