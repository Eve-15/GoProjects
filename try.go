package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"primaryKey"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
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

func main() {
	// 连接数据库
	var err error
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 自动迁移模型
	db.AutoMigrate(&User{})

	// 初始化用户仓库
	userRepo := NewMySQLUserRepository(db)

	// 测试增删改查
	testCRUD(userRepo)
}

func testCRUD(repo UserRepository) {
	// 创建用户
	user := &User{ID: "1", Name: "John", Age: 30}
	if err := repo.Create(user); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("User created: %+v\n", user)

	// 查询用户
	retrievedUser, err := repo.GetByID(user.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v", err)
	}
	fmt.Printf("User retrieved: %+v\n", *retrievedUser)

	// 更新用户
	retrievedUser.Age = 40
	if err := repo.Update(retrievedUser); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	fmt.Printf("User updated: %+v\n", *retrievedUser)

	// 删除用户
	if err := repo.Delete(retrievedUser.ID); err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("User deleted successfully")
}
