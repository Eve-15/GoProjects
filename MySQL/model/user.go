package model

import (
	"gorm.io/gorm"
	"math/rand"
	"sort"
	"time"
)

var Deck = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

var cardValue = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 10, "Q": 10, "K": 10,
}

func Fapai(numPlayers, numCardsPerPlayer int) [][]string {
	hands := make([][]string, numPlayers)
	for i := 0; i < numCardsPerPlayer; i++ {
		for j := 0; j < numPlayers; j++ {
			hands[j] = append(hands[j], Deck[i*numPlayers+j])
			Deck = append(Deck[:i*numPlayers+j], Deck[i*numPlayers+j+1:]...)
		}
	}
	return hands
}

type Player interface {
	Xipai()
	Paixu()
	Score() int
	GetCard() []string
}

type User struct {
	ID     string `gorm:"primaryKey"`
	Hand   []string
	Uscore int
}

func (u *User) Xipai() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Deck), func(i, j int) {
		Deck[i], Deck[j] = Deck[j], Deck[i]
	})
}

func (u *User) Paixu() {
	sort.Slice(u.Hand, func(i, j int) bool {
		score1 := cardValue[u.Hand[i]]
		score2 := cardValue[u.Hand[j]]
		return score1 < score2
	})
}

func (u *User) Score() int {
	var score int
	for i := 0; i < len(u.Hand); i++ {
		score += cardValue[u.Hand[i]]
	}
	return score
}

func (u *User) GetCard() []string {
	if len(Deck) == 0 {
		return u.Hand
	}
	newCard := Deck[len(Deck)-1]
	Deck = Deck[:len(Deck)-1]
	return append(u.Hand, newCard)
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) GetByID(id string) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *MySQLUserRepository) Delete(id string) error {
	return r.db.Delete(&User{}, id).Error
}
