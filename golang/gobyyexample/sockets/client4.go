package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8081"
	list, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("server Accept")
		go Handler(conn)
	}

}

var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn, 1),
	}
}
func (r *Room) Join(username string, conn net.Conn) {
	_, ok := r.users[username]
	if ok {
		r.Leave(username)
	}
	r.users[username] = conn
	fmt.Printf("%s登陆成功", username)
	conn.Write([]byte(username + "加入聊天室！"))
}
func (r *Room) Leave(username string) {
	conn, ok := r.users[username]
	if !ok {
		fmt.Printf("%v不存在", username)
	}
	conn.Close()
	delete(r.users, username)
	fmt.Printf("%v已离开", username)
}
func (r *Room) Broadcast(who string, msg string) {
	_time := time.Now().Format("2006年01月02日 15时04分05秒")
	_msg := fmt.Sprintf("%v %s %s\n", _time, who, msg)
	for k, v := range r.users {
		if k == who {
			continue
		}
		v.Write([]byte(_msg))
	}
}
func Handler(conn net.Conn) {
	defer conn.Close()
	fmt.Println("server Handler")
	r := bufio.NewReader(conn)
	username := ""
	for {
		input, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(input)
		fields := strings.Fields(input)
		if len(fields) != 2 {
			fmt.Println("请输入正确的用户名和密码")
			continue
		}
		username = fields[0]
		pwd := fields[1]
		if pwd != "123" {
			fmt.Println("密码不正确")
			continue
		}
		break
	}

	globalRoom.Join(username, conn)
	globalRoom.Broadcast("system", username+"加入房间")
	for {
		conn.Write([]byte("按回车键发送消息："))
		input, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(input, err)
			//panic(err)
			break
		}
		fmt.Println(username, input)
		globalRoom.Broadcast(username, strings.TrimSpace(input))
	}
	globalRoom.Broadcast("system", username+"leave room")
	globalRoom.Leave(username)

}
