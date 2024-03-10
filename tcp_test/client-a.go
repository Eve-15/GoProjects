package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.111.1:8888")
	if err != nil {
		fmt.Println("server dial err=", err)
		return
	}
	fmt.Println("conn 成功=", conn)
	//客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入[终端]
	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		//如果用户输入的是exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}
		//再将line 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
		// 从服务器接收信息
		buf := make([]byte, 1024)
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read fasiled err=%v", err)
			return
		}
		fmt.Println("收到服务器回复：", string(buf[:n]))
	}

}
