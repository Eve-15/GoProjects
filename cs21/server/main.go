package main

import (
	"bufio"
	"fmt"
	"github.com/Eve-15/GoProjects/cs21/server/model"
	"net"
	"strings"
	"sync"
)

var userCount int
var maxUsers int = 2

func main() {
	userInfo := model.UserInfo{
		Users: make(map[string]*model.User),
		Mutex: sync.Mutex{},
		Wg:    sync.WaitGroup{},
	}
	listener, err := net.Listen("tcp", "192.168.48.167:8080")
	if err != nil {
		fmt.Println("无法监听端口:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器已启动，等待客户端连接...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("无法接受客户端连接:", err)
			continue
		}

		userInfo.Mutex.Lock()
		userID := "" // 初始化客户端ID
		userCount++  //用来实现计数器
		userInfo.Mutex.Unlock()

		// 读取客户端发送的标识符
		reader := bufio.NewReader(conn)
		userID, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("无法读取客户端标识符:", err)
			conn.Close()
			continue
		}

		// 去除标识符中的换行符和空格
		userID = strings.TrimSpace(userID)
		//初始化user
		userInfo.Mutex.Lock()
		user := model.User{
			ID:   userID,
			Conn: conn,
		}
		userInfo.Users[userID] = &user
		userInfo.Mutex.Unlock()

		fmt.Printf("user[%s]已连接\n", userID)
		userInfo.Wg.Add(1)
		go userInfo.HandleUserConnection(user)
		userInfo.Wg.Wait()
		// 等待所有客户端处理完成
		userInfo.Mutex.Lock()
		if userCount == maxUsers {
			userInfo.EndGame()
		}
		userInfo.Mutex.Unlock()
	}
}
