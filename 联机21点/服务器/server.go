package main

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

type Client struct {
	ID    string
	Conn  net.Conn
	hand  []string
	score int
}

type ClientInfo struct {
	Clients map[string]*Client
	mutex   sync.Mutex
	wg      sync.WaitGroup
}

var clientInfo = ClientInfo{
	Clients: make(map[string]*Client),
}

var hands [2][]string
var clientCount int
var maxClients int = 2

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

func broadcastMessage(message string) {
	for _, client := range clientInfo.Clients {
		_, err := client.Conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("无法发送消息给客户端[%s]：%s\n", client.ID, err)
		}
	}
}

func handleClientConnection(client Client) {
	defer clientInfo.wg.Done()

	reader := bufio.NewReader(client.Conn)
	writer := bufio.NewWriter(client.Conn)

	// 发牌
	Xipai(Deck)
	hands := Fapai(Deck, 2, 2) // 两名玩家，每人发2张牌
	Paixu(hands[0])            // 玩家1的手牌排序
	Paixu(hands[1])            // 玩家2的手牌排序

	// 发送手牌给客户端
	if client.ID == "1" {
		client.hand = hands[0]
		writer.WriteString(fmt.Sprintf("玩家1的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", client.hand))
		writer.Flush()

	} else if client.ID == "2" {
		client.hand = hands[1]
		writer.WriteString(fmt.Sprintf("玩家2的手牌：%s\n-----欢迎来到纸牌游戏21点-----\n", client.hand))
		writer.Flush()
	}

	// 调用 PlayerGameStart 函数，传递玩家手牌
	score, lhand := PlayerGameStart(client.hand, writer, reader)
	clientInfo.mutex.Lock()
	clientInfo.Clients[client.ID].score = score
	clientInfo.Clients[client.ID].hand = lhand
	clientInfo.mutex.Unlock()

	// 发送玩家得分信息给客户端
	response := fmt.Sprintf("玩家[%s]的得分是：%v\n", client.ID, score)
	_, err := writer.WriteString(response)
	if err != nil {
		fmt.Printf("无法发送消息给客户端[%s]：%s\n", client.ID, err)
	}
	writer.Flush()

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
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

		clientInfo.mutex.Lock()
		clientID := "" // 初始化客户端ID
		clientCount++
		clientInfo.mutex.Unlock()

		// 读取客户端发送的标识符
		reader := bufio.NewReader(conn)
		clientID, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("无法读取客户端标识符:", err)
			conn.Close()
			continue
		}

		// 去除标识符中的换行符和空格
		clientID = strings.TrimSpace(clientID)
		clientInfo.mutex.Lock()
		client := Client{
			ID:   clientID,
			Conn: conn,
		}
		clientInfo.Clients[clientID] = &client
		clientInfo.mutex.Unlock()

		fmt.Printf("客户端[%s]已连接\n", clientID)
		clientInfo.wg.Add(1)
		go handleClientConnection(client)
		clientInfo.wg.Wait()
		// 等待所有客户端处理完成
		clientInfo.mutex.Lock()
		if clientCount == maxClients {
			clientInfo.wg.Wait()

			// 广播消息
			broadcastMessage("————————————游戏结束，正在生成最终结果————————————")
			time.Sleep(1 * time.Second)
			// 输出客户端信息
			broadcastMessage("玩家1的最终得分是：" + (strconv.Itoa(clientInfo.Clients["1"].score)))
			broadcastMessage("玩家2的最终得分是：" + (strconv.Itoa(clientInfo.Clients["2"].score)))
			time.Sleep(1 * time.Second)
			if clientInfo.Clients["1"].score > 21 {
				broadcastMessage("玩家" + clientInfo.Clients["2"].ID + "胜利!" + "玩家" + clientInfo.Clients["1"].ID + "爆牌了")
			} else if clientInfo.Clients["2"].score > 21 {
				broadcastMessage("玩家" + clientInfo.Clients["1"].ID + "胜利!" + "玩家" + clientInfo.Clients["2"].ID + "爆牌了")
			} else if clientInfo.Clients["1"].score == 21 {
				broadcastMessage("玩家" + clientInfo.Clients["1"].ID + "以王者之姿胜利")
			} else if clientInfo.Clients["2"].score == 21 {
				broadcastMessage("玩家" + clientInfo.Clients["2"].ID + "以王者之姿胜利")
			} else if clientInfo.Clients["1"].score > clientInfo.Clients["2"].score {
				broadcastMessage("玩家" + clientInfo.Clients["1"].ID + "以较大的点数取得胜利")
			} else if clientInfo.Clients["2"].score > clientInfo.Clients["1"].score {
				broadcastMessage("玩家" + clientInfo.Clients["2"].ID + "以较大的点数取得胜利")
			} else {
				broadcastMessage("点数相同，平局了")
			}
			// 重置计数器和客户端列表
			clientCount = 0
			clientInfo.mutex.Lock()
			clientInfo.Clients = make(map[string]*Client)
			clientInfo.mutex.Unlock()
		}
		clientInfo.mutex.Unlock()
	}
}
