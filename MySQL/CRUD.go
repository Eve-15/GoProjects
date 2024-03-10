package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"sort"
	"time"
)

type User struct {
	ID string `gorm:"primaryKey"` //设为主键
	//Conn  net.Conn
	hand  []string
	score int
}

var hands [2][]string

// Deck 纸牌
var Deck = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

// cardValue 点数
var cardValue = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10,
}

//游戏逻辑所需要的函数

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

// GetCard 再次摸牌,在此处需要注意切片索引和追加的语法规范,将新摸得的牌加入到手牌中
func GetCard(cards []string) []string {
	if len(Deck) == 0 {
		return cards
	}
	newCard := Deck[len(Deck)-1]
	Deck = Deck[:len(Deck)-1]
	return append(cards, newCard)
}

// UserRepository 定义接口用户仓库
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

// MySQLUserRepository 此结构体实现了UserRepository接口，注意熟悉接口定义及其格式
// 使用使用 gorm.DB 类型的 db 字段作为数据库连接对象，只是结构体类型，使用需要实例化
type MySQLUserRepository struct {
	db *gorm.DB
}

// NewMySQLUserRepository 用于创建MySQLUserRepository的结构体实例
func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

// Create 向表中插入数据
// r.db：代表一个数据库连接对象
func (r *MySQLUserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

// GetByID 根据主键id来从表中提取数据存储在user中，并返回一个查询到的行记录
func (r *MySQLUserRepository) GetByID(id string) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新操作，传入一个已经查到的行记录将其修改
func (r *MySQLUserRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete 根据查询到的行记录使用delete语句删除指定的数据
func (r *MySQLUserRepository) Delete(id string) error {
	return r.db.Delete(&User{}, id).Error
}

func testCRUD(repo UserRepository, hands [][]string) {
	// 发牌及抽牌

	var newhands []string
	newhands = GetCard(hands[0])
	var score int
	score = Score(newhands)
	// 创建用户
	user := &User{ID: "1", hand: hands[0], score: Score(hands[0])}
	if err := repo.Create(user); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("创建用户: %+v\n", user)

	// 查询用户
	retrievedUser, err := repo.GetByID(user.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v", err)
	}
	fmt.Printf("查询用户: %+v\n", *retrievedUser)

	// 更新用户
	retrievedUser.hand = newhands
	retrievedUser.score = score
	if err := repo.Update(retrievedUser); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	fmt.Printf("用户更新: %+v\n", *retrievedUser)

	// 删除用户
	if err := repo.Delete(retrievedUser.ID); err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("删除成功")
}

func main() {
	// 连接数据库

	var err error
	dsn := "root:20050315yisheng@tcp(127.0.0.1:3306)/YiSheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 自动迁移模型
	db.AutoMigrate(&User{})

	// 初始化用户仓库,一旦实例化了 MySQLUserRepository 结构体，就可以直接通过该实例来访问 db 字段，而不需要再次使用 r.db。
	userRepo := NewMySQLUserRepository(db)

	//游戏逻辑实现

	Xipai(Deck)
	hands := Fapai(Deck, 2, 2) // 两名玩家，每人发2张牌
	Paixu(hands[0])            // 玩家1的手牌排序
	Paixu(hands[1])            // 玩家2的手牌排序

	// 测试增删改查
	testCRUD(userRepo, hands)
}
