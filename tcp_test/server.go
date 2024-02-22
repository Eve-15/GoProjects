package main

import (
	"fmt"
	"net" //做socket开发时，此包含有所需的所有方法和函数
)

func process(conn net.Conn) {
	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端没有write[发送]，那么协程就阻塞在这里
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		recv := string(buf[:n])
		fmt.Printf("接收到的数据：%v", recv)
		//向客户端发送数据
		_, _ = conn.Write([]byte("ok"))
	}
}

// 监听8888号端口
// 需要协程服务
func main() {
	fmt.Println("服务器开始监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待客户端来链接
	//noinspection GoInfiniteLoop
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备其一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen suc=%v\n", listen)
}
