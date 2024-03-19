package controller

import (
	"bufio"
	"fmt"
	"github.com/Eve-15/GoProjects/rebuild/server/model"
	"strconv"
	"time"
)

type GameController struct {
	UserInfo *model.UserInfo
}

// NewGameController 函数用于创建一个新的游戏控制器实例，并初始化其中的用户信息。
func NewGameController(userInfo *model.UserInfo) *GameController {
	return &GameController{
		UserInfo: userInfo, //将传入的 userInfo 参数赋值给新创建的游戏控制器的 UserInfo 字段，以便控制器能够访问和操作用户信息。
	}
}

// BroadcastMessage 将方法定义为func (gc *GameController) BroadcastMessage(message string)，表示将BroadcastMessage方法与GameController结构体关联起来。这样定义后，BroadcastMessage方法就成为GameController结构体的一个成员方法，可以通过gc.BroadcastMessage(message)来调用
func (gc *GameController) BroadcastMessage(message string) {
	for _, user := range gc.UserInfo.Users {
		writer := bufio.NewWriter(user.Conn)
		writer.WriteString(message + "\n")
		writer.Flush()
	}
}

func (gc *GameController) HandleUserConnection(user model.User) {
	defer gc.UserInfo.Wg.Done()

	reader := bufio.NewReader(user.Conn) //对连接的数据读取
	writer := bufio.NewWriter(user.Conn) //向连接写入数据

	// 发牌
	model.Xipai(model.Deck)
	hands := model.Fapai(model.Deck, 2, 2)
	model.Paixu(hands[0])
	model.Paixu(hands[1])

	if user.ID == "1" {
		user.Hand = hands[0]
		writer.WriteString(fmt.Sprintf("玩家1的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", user.Hand))
		writer.Flush()
	} else if user.ID == "2" {
		user.Hand = hands[1]
		writer.WriteString(fmt.Sprintf("玩家2的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", user.Hand))
		writer.Flush()
	}

	// 调用 PlayerGameStart 函数，传递玩家手牌
	score, lhand := model.PlayerGameStart(user.Hand, writer, reader)
	gc.UserInfo.Mutex.Lock()
	gc.UserInfo.Users[user.ID].Score = score
	gc.UserInfo.Users[user.ID].Hand = lhand
	gc.UserInfo.Mutex.Unlock()

	// 发送玩家得分信息给客户端
	response := fmt.Sprintf("玩家[%s]的得分是：%v\n", user.ID, score)
	_, err := writer.WriteString(response)
	if err != nil {
		fmt.Printf("无法发送消息给客户端[%s]：%s\n", user.ID, err)
	}
	writer.Flush()
}

func (gc *GameController) EndGame() {
	gc.BroadcastMessage("————————————游戏结束，正在生成最终结果————————————")
	time.Sleep(1 * time.Second)
	gc.BroadcastMessage("玩家1的最终得分是：" + strconv.Itoa(gc.UserInfo.Users["1"].Score))
	gc.BroadcastMessage("玩家2的最终得分是：" + strconv.Itoa(gc.UserInfo.Users["2"].Score))
	time.Sleep(1 * time.Second)
	if gc.UserInfo.Users["1"].Score > 21 {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["2"].ID + "胜利!" + "玩家" + gc.UserInfo.Users["1"].ID + "爆牌了")
	} else if gc.UserInfo.Users["2"].Score > 21 {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["1"].ID + "胜利!" + "玩家" + gc.UserInfo.Users["2"].ID + "爆牌了")
	} else if gc.UserInfo.Users["1"].Score == 21 {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["1"].ID + "以王者之姿胜利")
	} else if gc.UserInfo.Users["2"].Score == 21 {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["2"].ID + "以王者之姿胜利")
	} else if gc.UserInfo.Users["1"].Score > gc.UserInfo.Users["2"].Score {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["1"].ID + "以较大的点数取得胜利")
	} else if gc.UserInfo.Users["2"].Score > gc.UserInfo.Users["1"].Score {
		gc.BroadcastMessage("玩家" + gc.UserInfo.Users["2"].ID + "以较大的点数取得胜利")
	} else {
		gc.BroadcastMessage("点数相同，平局了")
	}
}
