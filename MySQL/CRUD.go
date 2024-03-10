package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID string `gorm:"primaryKey"` //设为主键
	//Conn  net.Conn
	hand  []string
	score int
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

	// 初始化用户仓库,一旦实例化了 MySQLUserRepository 结构体，就可以直接通过该实例来访问 db 字段，而不需要再次使用 r.db。
	userRepo := NewMySQLUserRepository(db)

	// 测试增删改查

}
