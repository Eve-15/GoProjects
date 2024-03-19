package main

import (
	"fmt"
	"github.com/Eve-15/GoProjects/MySQL/controller"
	"github.com/Eve-15/GoProjects/MySQL/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	var hands [][]string
	var score int
	var err error
	hands = model.Fapai(2, 2)
	//fmt.Println(hands[0])
	dsn := "root:20050315yisheng@tcp(127.0.0.1:3306)/YiSheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 自动迁移模型
	db.AutoMigrate(&model.User{})

	user := &model.User{
		ID: "1", Hand: hands[0], Uscore: score,
	}

	// 创建用户信息控制器
	userController := &controller.UserController{
		UserRepo: model.NewMySQLUserRepository(db),
	}

	// 创建游戏控制器时，使用实现了 Player 接口的结构体实例
	gameController := controller.NewGameController(&model.User{})

	//进行游戏操作
	gameController.Xi()
	// 调用 GameController 的方法

	if len(hands) == 0 {
		log.Fatalf("Failed to get hands from GameController.Fa()")
	}

	score = gameController.UserScore()
	gameController.Pai()
	//进行CRUD操作
	//创建用户
	if err := userController.CreatUser(user); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("创建用户: %+v\n", user)
	// 查询用户（此处有问题还需思考）
	retrievedUser, err := userController.GetUserByID(user.ID) //此处的ID是基于已经创建的表进行还是单只第一个ID，如果表中有两个ID的时候会怎么办
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v", err)
	}
	fmt.Printf("查询用户: %+v\n", *user) //此处若为user则正常显示，若为retrievedUser则显示为空切片
	// 更新用户
	retrievedUser.Hand = gameController.Get()
	retrievedUser.Uscore = gameController.UserScore()
	if err := userController.UpdateUser(retrievedUser); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	fmt.Printf("用户更新: %+v\n", *retrievedUser)

	// 删除用户
	if err := userController.DeleteUser(retrievedUser.ID); err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("删除成功")

}
