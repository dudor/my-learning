package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main()  {
	addr := "www.baidu.com:80"
	conn,err:= net.Dial("tcp",addr)
	defer conn.Close()
	if err!= nil{
		panic(err)
	}
	n,err:= conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil{
		panic(err)
	}
	fmt.Println("写入的大小是:",n)
	f,err:=os.Open("./test1.txt")
	io.Copy(f,conn)
	f.Close()
}