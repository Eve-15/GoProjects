package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.111.1:8080")
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	// 发送客户端标识符给服务器
	fmt.Print("请输入客户端标识符（1或2）：")
	clientID, _ := reader.ReadString('\n')
	clientID = strings.TrimSpace(clientID)
	writer.WriteString(clientID + "\n")
	writer.Flush()

	// 读取服务器发送的消息
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("无法读取消息:", err)
				break
			}
			fmt.Println(msg)
		}
	}()

	// 从标准输入读取并发送给服务器
	for {
		fmt.Print("server（exit退出）：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		writer.WriteString(msg + "\n")
		writer.Flush()

		if msg == "exit" {
			break
		}
	}
}
