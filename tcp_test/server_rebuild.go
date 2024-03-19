package main

import (
	"bufio"
	"math/rand"
	"net"
	"sort"
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
	mutex sync.Mutex
	wg    sync.WaitGroup
}

type PokerCards struct {
	Deck      []string
	cardValue map[string]int
}

// PokerGame 游戏底层逻辑的接口
type PokerGame interface {
	XiPai(cards []string)
	FaPai(cards []string, numPlayers int, numCardsPerPlayer int) [][]string
	PaiXu(hand []string)
	Score(hand []string) int
	GetCard() string
	PlayerGameStart(hand []string, writer *bufio.Writer, reader *bufio.Reader) (int, []string)
}

// CnS C/S架构的接口
type CnS interface {
	handleClientConnection(client User)
	broadcastMessage(message string)
}

// UserRepository CRUD接口
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

var userInfo = UserInfo{
	Users: make(map[string]*User),
}

var hands [2][]string
var userCount int
var maxUsers int = 2

var Deck = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

var cardValue = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10,
}

func (r *PokerCards) XiPai(cards []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func (r *User) FaPai(cards []string, numPlayers int, numCardsPerPlayer int) [][]string {
	hands := make([][]string, numPlayers)
	for i := 0; i < numCardsPerPlayer; i++ {
		for j := 0; j < numPlayers; j++ {
			hands[j] = append(hands[j], cards[i*numPlayers+j])
			Deck = append(Deck[:i*numPlayers+j], Deck[i*numPlayers+j+1:]...)
		}
	}
	return hands
}

func (r *User) PaiXu(hand []string) {
	sort.Slice(hand, func(i, j int) bool {
		score1 := cardValue[hand[i]]
		score2 := cardValue[hand[j]]
		return score1 < score2
	})
}

func (r *User) Score(hand []string) int { //此处注意返回值类型，若不加则产生问题返回实参过多
	var score int
	for i := 0; i < len(hand); i++ {
		score += cardValue[hand[i]] //通过地图将牌面对应点数相加
	}
	return score
}

func (r *User) GetCard() string {
	var newcard string
	newcard = Deck[len(Deck)-1]
	Deck = Deck[:len(Deck)-1]
	return newcard
}

func main() {

}
