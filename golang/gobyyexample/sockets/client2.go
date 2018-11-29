package main

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"time"
)

func main()  {
	for {
		time.Sleep(time.Millisecond * 2)
		go dialBaidu()
	}
	fmt.Println("die")
}

func dialBaidu()  {
	addr := "www.baidu.com:80"
	conn,err:= net.Dial("tcp",addr)
	defer conn.Close()
	if err!= nil{
		panic(err)
	}
	fmt.Println("访问的公司IP是",conn.RemoteAddr().String(),"本地IP是",conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.RemoteAddr()),reflect.TypeOf(conn.RemoteAddr().String()))
	n,err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err!= nil{
		panic(err)
	}
	fmt.Println("向服务端发送的数据大小是:",n)
	buf := make([]byte,1024)
	//n,err=conn.Read(buf)
	//if err!= nil && err != io.EOF{
	//	panic(err)
	//}

	for{
		n,err = conn.Read(buf)
		if err  == io.EOF{
			break
		}
		fmt.Println(string(buf[:n]))
	}
}