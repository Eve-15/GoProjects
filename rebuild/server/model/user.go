package model

import (
	"bufio"
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
	Hand  []string
	Score int
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
	for i := 0; i < numCardsPerPlayer; i++ { //玩家数量控制轮次
		for j := 0; j < numPlayers; j++ { //每人发几张牌
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
func Score(hand []string) int {
	var score int
	for i := 0; i < len(hand); i++ {
		score += cardValue[hand[i]]
	}
	return score
}

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
