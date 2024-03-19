package main

import (
	"bufio"
	"fmt"
	"github.com/Eve-15/GoProjects/rebuild/server/controller"
	"github.com/Eve-15/GoProjects/rebuild/server/model"
	"net"
	"strings"
)

var userCount int
var maxUsers int = 2

func main() {
	userInfo := &model.UserInfo{
		Users: make(map[string]*model.User),
	}
	//实例化控制器
	gameController := controller.NewGameController(userInfo)

	listener, err := net.Listen("tcp", "192.168.111.1:8080")
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
		userID := ""
		userInfo.Wg.Add(1)
		userCount++
		userInfo.Mutex.Unlock()

		reader := bufio.NewReader(conn)
		userID, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("无法读取客户端标识符:", err)
			conn.Close()
			continue
		}

		userID = strings.TrimSpace(userID)
		userInfo.Mutex.Lock()
		user := &model.User{
			ID:   userID,
			Conn: conn,
		}
		userInfo.Users[userID] = user
		userInfo.Mutex.Unlock()

		fmt.Printf("user[%s]已连接\n", userID)

		go gameController.HandleUserConnection(*user)

		userInfo.Wg.Wait()

		userInfo.Mutex.Lock()
		if userCount == maxUsers {
			gameController.EndGame()
		}
		userInfo.Mutex.Unlock()

	}

}
