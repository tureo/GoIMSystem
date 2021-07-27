package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前client的模式
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>请输入合法范围内的数字<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 根据不通的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			fmt.Println("公聊模式选择...")
		case 2:
			// 私聊模式
			fmt.Println("私聊模式选择...")
		case 3:
			// 更新用户名
			fmt.Println("更新用户名选择...")

		}
	}
}

var serverIP string
var serverPort int

//./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>连接服务器失败...")
		return
	}
	fmt.Println(">>>连接服务器成功...")

	// 启动客户端的业务
	client.Run()
}
