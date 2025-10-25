package main

import (
	"4gorm/global"
	"4gorm/models"
	"fmt"
)

func main() {
	global.InitDB()
	global.MigrateDB()
	fmt.Println("=== GORM serializer:json 标签演示 ===")

	// 创建用户业务逻辑
	//createUser1()

	// 查询用户业务逻辑
	find()

	fmt.Println("\n=== 演示完成 ===")
}

func find() {
	fmt.Println("\n=== 查询用户业务逻辑 ===")
	var user models.UserZdy
	global.DB.First(&user)
	fmt.Println(user, user.Settings.Language)
}

// createUser1 创建用户业务逻辑
func createUser1() {
	fmt.Println("\n=== 创建用户业务逻辑 ===")

	user := models.UserZdy{
		ID:       1,
		Username: "alice",
		Tags:     []string{"frontend", "react", "typescript"},
		Settings: models.Settings{
			Theme:    "light",
			Language: "zh",
		},
	}

	// 保存到数据库
	result := global.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("创建用户失败: %v\n", result.Error)
		return
	}

	fmt.Printf("创建用户成功: %s\n", user.Username)
	fmt.Printf("用户标签: %v\n", user.Tags)
	fmt.Printf("用户设置: %+v\n", user.Settings)
}
