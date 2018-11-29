package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	lis,err:=net.Listen("tcp",":8081")
	if err != nil{
		panic(err)
	}
	fmt.Println("server1 Listened")

	for{
		conn ,err:= lis.Accept()
		fmt.Println("server1 Accepted")
		if err!= nil{
			fmt.Println(err)
			break
		}
		go handler(conn)
	}
}

func handler(conn net.Conn)  {
	defer close(conn)
	fmt.Println("server1 handler()")
	//conn.Write([]byte(time.Now().String()))
	time.AfterFunc(time.Second*5, func() {
		conn.Close()
	})

	for{
		conn.Write([]byte(time.Now().String()))
		time.Sleep(time.Second)
	}

}
func close(conn net.Conn)  {
	conn.Close()
	fmt.Println("server1 Close()")
}
