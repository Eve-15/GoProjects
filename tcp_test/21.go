package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Deck 定义一个切片作为扑克牌数量除大小王共52张
var Deck = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

// 根据21点规则定义每张牌对应点数，2~10为2~10分，J、Q、K为10点A为1点（A按规则应为1或11点）
var cardValue = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10,
}

// Xipai 洗牌
func Xipai(cards []string) {
	rand.Seed(time.Now().UnixNano()) //设置随机数生成器的种子。通过使用当前时间的纳秒级别作为种子，可以确保每次运行程序时都会得到不同的随机数序列。
	//rand.Shuffle用于随机打乱切片中的元素顺序。它接受两个参数：切片的长度和一个用于交换元素的函数。
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i] //每次调用会传递两个参数 i 和 j，表示切片中要交换的两个元素的索引。
	})
}

// Fapai  发牌：传入3个变量分别是纸牌、玩家数量、手牌数量，通过双重循环来实现发牌
func Fapai(cards []string, numPlayers int, numCardsPerPlayer int) [][]string {
	hands := make([][]string, numPlayers)    //二维数组：第几个人和其对应手牌
	for i := 0; i < numCardsPerPlayer; i++ { //每个人的手牌数量
		for j := 0; j < numPlayers; j++ { //发牌的操作是比如总共3个玩家每人2张牌此函数采取先每人发一张每人发完一张后再每人发一张
			hands[j] = append(hands[j], cards[i*numPlayers+j]) //此处追加的元素需要注意防止重复分发纸牌，假如两名玩家玩家1取索引：0，2玩家2取索引：1，3
			Deck = append(Deck[:i*numPlayers+j], Deck[i*numPlayers+j+1:]...)
			//fmt.Println(hands[j])
		}

	}
	return hands //返回二维数组切片
}

// Paixu 排序手牌按照地图中不同牌面对应点数大小
func Paixu(hand []string) {
	sort.Slice(hand, func(i, j int) bool { //闭包的理解仍需要提高
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

// PlayerGameStart 游戏规则：分别传入玩家手牌，若点数小于21点通过循环来判断是否需要摸牌以及摸牌后的点数
func PlayerGameStart(hand []string) int {
	var score = Score(hand)
	var newcard string
	var choice string

	for score < 21 { //可以选择是否继续摸牌
		fmt.Printf("是否继续摸牌，如果摸牌请输入y,若不继续则输入n:\n")
		fmt.Scanf("%v", &choice)
		if choice == "n" {
			fmt.Println("适当的收手也是明智的选择，但你要准备接受豪赌怪人的挑战了")
			fmt.Println("____________________你的回合结束了_______________________")
			break
		}
		time.Sleep(1 * time.Second)
		if choice == "y" {
			newcard = Getcard()
			fmt.Println("你摸到了:", newcard)
			hand = append(hand, newcard)
			time.Sleep(1 * time.Second)
			fmt.Println("目前你的手牌是:", hand)
			score = Score(hand)
			fmt.Println("目前你的点数是:", score)
			time.Sleep(1 * time.Second)
			//fmt.Println("牌堆里还有:\n", Deck)
			fmt.Println("_________________________________________________________")
			if score < 21 {
				//fmt.Scanf("%v", &choice),也可以算一种解决方法
				fmt.Scanln() // 读取并丢弃换行符，在此处遇到了问题：连续多次询问是否继续摸牌
				//错误原因没有理解好go语言中的输入
				//fmt.Scanf 会在读取输入时遇到空格、制表符或换行符时停止读取。
				//当输入 "y" 或 "n" 后按下回车键，回车键会被视为换行符，而 fmt.Scanf 会将换行符留在输入缓冲区中。
				//这样，当下一次询问是否继续摸牌时，fmt.Scanf 会立即读取到换行符，
				//而不会等待用户输入。因此，会出现连续询问是否继续摸牌的情况
				continue

			}
			if score == 21 {
				time.Sleep(1 * time.Second)
				fmt.Println("你现在是21点，多么酣畅淋漓的一场豪赌呀，是吧豪赌怪人")
				break
			}
			if score > 21 { //此处有所不足没有考虑21点规则中的A可以看作1或10点
				time.Sleep(1 * time.Second)
				fmt.Println("豪赌怪人你勇气可嘉但是运气不好，爆牌了")
				break
			}
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("最终你的手牌是:", hand)
	fmt.Println("你的点数是:", score)
	return score
}

// ComputerGameStart 传入电脑手牌来判断是否摸牌，此处只是简单逻辑， 个人觉得如果需要电脑更加真实可能需要更多的判断让其更加智能
func ComputerGameStart(hand []string) int {
	var score = Score(hand)
	var newcard string
	var choice string
	for score < 17 { //此处定电脑为庄家，庄家点数达到或超过17点则不能再摸牌
		choice = "y"
		if choice == "y" {
			newcard = Getcard()
			hand = append(hand, newcard)
			score = Score(hand)
		}
	}
	fmt.Println("庄家最终的手牌是:", hand)
	fmt.Println("庄家的点数是:", score)
	//fmt.Println("牌堆里还有:", Deck)

	return score
}

// EndGame 先进行玩家回合再进行庄家回合，最后双方明牌，比较点数
func EndGame(hand1, hand2 []string) {

	PlayerScore := PlayerGameStart(hand1) //先进行玩家回合，玩家21点或爆牌则直接游戏结束
	if PlayerScore > 21 {
		fmt.Println("庄家最终的手牌是:", hand2)
		fmt.Println("庄家的点数是:", Score(hand2))
		fmt.Println("————游戏结束，玩家爆牌了，庄家胜利————")
	} else if PlayerScore == 21 {
		fmt.Println("庄家最终的手牌是:", hand2)
		fmt.Println("庄家的点数是:", Score(hand2))
		fmt.Println("————游戏结束，玩家以王者之姿胜利————")
	} else {
		fmt.Println("____________________庄家回合_______________________")
		time.Sleep(1 * time.Second)
		DealerScore := ComputerGameStart(hand2) //再进行庄家回合
		if DealerScore == 21 {
			fmt.Println("————游戏结束，庄家以王者之姿胜利————")
		} else if DealerScore > 21 {
			fmt.Println("————游戏结束，庄家爆牌了，玩家胜利————")
		} else {
			if DealerScore > PlayerScore {
				fmt.Println("————游戏结束，庄家点数大于玩家，庄家胜利————")
			} else {
				fmt.Println("————游戏结束，玩家点数大于庄家，玩家胜利————")
			}
		}
	}
}
func main() {
	fmt.Println("————————————————————————欢迎来到纸牌游戏21点———————————————————————————")
	fmt.Println("（如果想更好观察游戏过程可以取消96、142、196、198、201、205、207行有关于牌堆和庄家的注释）")
	// 洗牌
	Xipai(Deck)

	// 发牌给两名玩家，每人发2张牌
	hands := Fapai(Deck, 2, 2)

	// 对每个玩家的手牌进行排序,并分别分至不同变量之中保存
	var hand1, hand2 []string
	for i, hand := range hands {
		Paixu(hand)

		switch i { //将玩家手牌传至不同变量之中
		case 0:
			hand1 = hand
		case 1:
			hand2 = hand
		}
	}
	var score1 int
	//var score2 int
	score1 = Score(hand1)
	//score2 = Score(hand2)

	fmt.Printf("玩家的手牌: %v\n", hand1)
	//fmt.Printf("庄家的手牌: %v\n", hand2)

	count := len(Deck)
	fmt.Println("剩余牌的数量:", count)
	//fmt.Println("牌堆里还有哪些牌:", Deck)
	fmt.Println("你的初始点数为:", score1)
	//fmt.Println("庄家的初始点数为:", score2)

	//分别设置玩家回合和电脑回合，最终在EndGame函数中进行最终结果的判断
	EndGame(hand1, hand2)
}
