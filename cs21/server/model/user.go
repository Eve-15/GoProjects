package model

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID    string
	Conn  net.Conn
	hand  []string
	score int
}

type UserInfo struct {
	Users map[string]*User
	Mutex sync.Mutex
	Wg    sync.WaitGroup
}

var Deck = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

var cardValue = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10,
}

func Xipai(cards []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func Fapai(cards []string, numPlayers int, numCardsPerPlayer int) [][]string {
	hands := make([][]string, numPlayers)
	for i := 0; i < numCardsPerPlayer; i++ {
		for j := 0; j < numPlayers; j++ {
			hands[j] = append(hands[j], cards[i*numPlayers+j])
			Deck = append(Deck[:i*numPlayers+j], Deck[i*numPlayers+j+1:]...)
		}
	}
	return hands
}

func Paixu(hand []string) {
	sort.Slice(hand, func(i, j int) bool {
		score1 := cardValue[hand[i]]
		score2 := cardValue[hand[j]]
		return score1 < score2
	})
}

// Score 每名玩家的点数
func Score(hand []string) int { //此处注意返回值类型，若不加则产生问题返回实参过多
	var score int
	for i := 0; i < len(hand); i++ {
		score += cardValue[hand[i]] //通过地图将牌面对应点数相加
	}
	return score
}

// Getcard 再次摸牌,在此处需要注意切片索引和追加的语法规范,将新摸得的牌加入到手牌中
func Getcard() string {
	var newcard string
	newcard = Deck[len(Deck)-1]
	Deck = Deck[:len(Deck)-1]
	return newcard
}

func PlayerGameStart(hand []string, writer *bufio.Writer, reader *bufio.Reader) (int, []string) {
	var score = Score(hand)
	var newcard string
	var choice string

	for score < 21 {
		writer.WriteString("是否继续摸牌？如果摸牌请输入y，若不继续则输入n:\n")
		writer.Flush()

		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "n" {
			writer.WriteString("适当的收手也是明智的选择，但你要准备接受豪赌怪人的挑战了\n")
			writer.WriteString("____________________你的回合结束了_______________________\n")
			writer.Flush()
			break
		}

		time.Sleep(1 * time.Second)

		if choice == "y" {
			newcard = Getcard()
			writer.WriteString("你摸到了: " + newcard + "\n")
			writer.Flush()

			hand = append(hand, newcard)
			time.Sleep(1 * time.Second)

			writer.WriteString("目前你的手牌是: " + strings.Join(hand, ", ") + "\n")
			writer.Flush()

			score = Score(hand)
			writer.WriteString("目前你的点数是: " + strconv.Itoa(score) + "\n")
			writer.Flush()

			time.Sleep(1 * time.Second)
			writer.WriteString("_________________________________________________________\n")
			writer.Flush()

			if score < 21 {
				continue
			}

			if score == 21 {
				time.Sleep(1 * time.Second)
				writer.WriteString("你现在是21点，多么酣畅淋漓的一场豪赌呀，是吧豪赌怪人\n")
				writer.Flush()
				break
			}

			if score > 21 {
				time.Sleep(1 * time.Second)
				writer.WriteString("豪赌怪人你勇气可嘉但是运气不好，爆牌了\n")
				writer.Flush()
				break
			}
		}
	}

	time.Sleep(1 * time.Second)
	writer.WriteString("最终你的手牌是: " + strings.Join(hand, ", ") + "\n")
	writer.WriteString("你的点数是: " + strconv.Itoa(score) + "\n")
	writer.Flush()

	return score, hand
}

func (userInfo *UserInfo) BroadcastMessage(message string) {
	for _, user := range userInfo.Users {
		_, err := user.Conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("无法发送消息给客户端[%s]：%s\n", user.ID, err)
		}
	}
}

func (userInfo *UserInfo) EndGame() {
	// 广播消息
	userInfo.BroadcastMessage("————————————游戏结束，正在生成最终结果————————————")
	time.Sleep(1 * time.Second)
	// 输出客户端信息
	userInfo.BroadcastMessage("玩家1的最终得分是：" + (strconv.Itoa(userInfo.Users["1"].score)))
	userInfo.BroadcastMessage("玩家2的最终得分是：" + (strconv.Itoa(userInfo.Users["2"].score)))
	time.Sleep(1 * time.Second)
	if userInfo.Users["1"].score > 21 {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["2"].ID + "胜利!" + "玩家" + userInfo.Users["1"].ID + "爆牌了")
	} else if userInfo.Users["2"].score > 21 {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["1"].ID + "胜利!" + "玩家" + userInfo.Users["2"].ID + "爆牌了")
	} else if userInfo.Users["1"].score == 21 {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["1"].ID + "以王者之姿胜利")
	} else if userInfo.Users["2"].score == 21 {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["2"].ID + "以王者之姿胜利")
	} else if userInfo.Users["1"].score > userInfo.Users["2"].score {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["1"].ID + "以较大的点数取得胜利")
	} else if userInfo.Users["2"].score > userInfo.Users["1"].score {
		userInfo.BroadcastMessage("玩家" + userInfo.Users["2"].ID + "以较大的点数取得胜利")
	} else {
		userInfo.BroadcastMessage("点数相同，平局了")
	}
}

func (userInfo *UserInfo) HandleUserConnection(user User) {
	defer userInfo.Wg.Done() //当客户端连接后并执行一次逻辑后减少等待组

	reader := bufio.NewReader(user.Conn)
	writer := bufio.NewWriter(user.Conn)

	// 发牌
	Xipai(Deck)
	hands := Fapai(Deck, 2, 2) // 两名玩家，每人发2张牌
	Paixu(hands[0])            // 玩家1的手牌排序
	Paixu(hands[1])            // 玩家2的手牌排序

	// 发送基本手牌给客户端
	if user.ID == "1" {
		user.hand = hands[0]
		writer.WriteString(fmt.Sprintf("玩家1的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", user.hand))
		writer.Flush()

	} else if user.ID == "2" {
		user.hand = hands[1]
		writer.WriteString(fmt.Sprintf("玩家2的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", user.hand))
		writer.Flush()
	}

	// 调用 PlayerGameStart 函数，传递玩家手牌
	score, lhand := PlayerGameStart(user.hand, writer, reader) //求得玩家最终手牌以及得分
	userInfo.Mutex.Lock()
	userInfo.Users[user.ID].score = score
	userInfo.Users[user.ID].hand = lhand
	userInfo.Mutex.Unlock()

	// 发送玩家得分信息给客户端
	response := fmt.Sprintf("玩家[%s]的得分是：%v\n", user.ID, score)
	_, err := writer.WriteString(response)
	if err != nil {
		fmt.Printf("无法发送消息给客户端[%s]：%s\n", user.ID, err)
	}
	writer.Flush()

}
