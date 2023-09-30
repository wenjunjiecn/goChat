package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net Dial error:", err)
		return nil
	}

	client.conn = conn
	return client
}

func (client *Client) menu() bool {
	var f int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&f)

	if f >= 0 && f <= 3 {
		client.flag = f
		return true
	} else {
		fmt.Println(">>>>>>>输入合法范围内的数字>>>>>>>>")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			fmt.Println("进入公聊模式")
			break
		case 2:
			fmt.Println("进入私聊模式")
			break
		case 3:
			fmt.Println("更改用户名模式")
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器ip，默认127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "服务端接口，默认8888")
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>连接服务器失败>>>>>>>>>>")
		return
	}

	fmt.Println(">>>>>>>>>连接服务器成功>>>>>>>>>>>")

	client.Run()
}
