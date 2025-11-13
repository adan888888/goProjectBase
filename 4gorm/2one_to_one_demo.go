package main

import (
	"fmt"

	"4gorm/global"
	"4gorm/models"
)

// 创建用户
func createUser(name string, age int) {
	user := models.UserModel{
		Name: name,
		Age:  age,
	}
	result := global.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("创建用户失败: %v\n", result.Error)
		return
	}
	fmt.Printf("创建用户成功! ID: %d, 姓名: %s, 年龄: %d\n", user.ID, user.Name, user.Age)
}

// 查询所有用户
func getAllUsers() {
	var users []models.UserModel
	result := global.DB.Find(&users)
	if result.Error != nil {
		fmt.Printf("查询用户失败: %v\n", result.Error)
		return
	}
	fmt.Printf("查询到 %d 个用户:\n", len(users))
	for _, user := range users {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 创建时间: %s\n",
			user.ID, user.Name, user.Age, user.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

// 根据ID查询用户
func getUserByID(id uint) {
	var user models.UserModel
	result := global.DB.First(&user, id)
	if result.Error != nil {
		fmt.Printf("查询用户失败: %v\n", result.Error)
		return
	}
	fmt.Printf("查询到用户: ID: %d, 姓名: %s, 年龄: %d, 创建时间: %s\n",
		user.ID, user.Name, user.Age, user.CreatedAt.Format("2006-01-02 15:04:05"))
}

// 根据姓名查询用户
func getUserByName(name string) {
	var users []models.UserModel
	result := global.DB.Where("name = ?", name).Find(&users)
	if result.Error != nil {
		fmt.Printf("查询用户失败: %v\n", result.Error)
		return
	}
	fmt.Printf("查询到 %d 个名为 '%s' 的用户:\n", len(users), name)
	for _, user := range users {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", user.ID, user.Name, user.Age)
	}
}

// 更新用户信息
func updateUser(id uint, name string, age int) {
	var user models.UserModel
	result := global.DB.First(&user, id)
	if result.Error != nil {
		fmt.Printf("查询用户失败: %v\n", result.Error)
		return
	}

	user.Name = name
	user.Age = age
	result = global.DB.Save(&user)
	if result.Error != nil {
		fmt.Printf("更新用户失败: %v\n", result.Error)
		return
	}
	fmt.Printf("更新用户成功! ID: %d, 姓名: %s, 年龄: %d\n", user.ID, user.Name, user.Age)
}

// 删除用户（软删除）
func deleteUser(id uint) {
	var user models.UserModel
	result := global.DB.Delete(&user, id)
	if result.Error != nil {
		fmt.Printf("删除用户失败: %v\n", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Printf("未找到ID为 %d 的用户\n", id)
		return
	}
	fmt.Printf("删除用户成功! 删除了 %d 条记录\n", result.RowsAffected)
}

// 永久删除用户
func permanentlyDeleteUser(id uint) {
	var user models.UserModel
	result := global.DB.Unscoped().Delete(&user, id)
	if result.Error != nil {
		fmt.Printf("永久删除用户失败: %v\n", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Printf("未找到ID为 %d 的用户\n", id)
		return
	}
	fmt.Printf("永久删除用户成功! 删除了 %d 条记录\n", result.RowsAffected)
}

// 创建用户详情
func createUserDetail(userID int, email string) {
	userDetail := models.UserDetailModel{
		UserID: userID,
		Email:  email,
	}
	result := global.DB.Create(&userDetail)
	if result.Error != nil {
		fmt.Printf("创建用户详情失败: %v\n", result.Error)
		return
	}
	fmt.Printf("创建用户详情成功! ID: %d, 用户ID: %d, 邮箱: %s\n", userDetail.ID, userDetail.UserID, userDetail.Email)
}

// 查询用户及其详情（预加载）
func getUserWithDetail(id uint) {
	var user models.UserModel
	result := global.DB.Preload("UserDetailModel").First(&user, id)
	if result.Error != nil {
		fmt.Printf("查询用户详情失败: %v\n", result.Error)
		return
	}
	fmt.Printf("查询到用户: ID: %d, 姓名: %s, 年龄: %d\n", user.ID, user.Name, user.Age)
	if user.UserDetailModel != nil {
		fmt.Printf("用户详情: 邮箱: %s\n", user.UserDetailModel.Email)
	} else {
		fmt.Println("该用户暂无详情信息")
	}
}

func main() {
	// 初始化数据库连接
	global.InitDB()

	// 自动迁移数据库表
	global.MigrateDB()

	fmt.Println("\n=== GORM 增删改查演示 ===")

	// 1. 创建用户
	demonstrateCreateUsers()

	// 2. 查询所有用户
	demonstrateGetAllUsers()

	// 3. 根据ID查询用户
	demonstrateGetUserByID()

	// 4. 根据姓名查询用户
	demonstrateGetUserByName()

	// 5. 更新用户信息
	demonstrateUpdateUser()

	// 6package_manager. 再次查询所有用户
	demonstrateGetAllUsersAfterUpdate()

	// 7. 创建用户详情
	demonstrateCreateUserDetails()

	// 8. 查询用户及其详情
	demonstrateGetUsersWithDetails()

	// 9. 软删除用户
	demonstrateSoftDeleteUser()

	// 10. 查询所有用户（包括软删除的）
	demonstrateGetAllUsersIncludingDeleted()

	// 11. 查询未删除的用户
	demonstrateGetActiveUsers()

	// 12. 永久删除用户
	demonstratePermanentlyDeleteUser()

	// 13. 最终查询所有用户
	demonstrateFinalGetAllUsers()

	fmt.Println("\n=== 演示完成 ===")
}

// 演示创建用户
func demonstrateCreateUsers() {
	fmt.Println("\n1. 创建用户:")
	createUser("张三", 25)
	createUser("李四", 30)
	createUser("王五", 28)
}

// 演示查询所有用户
func demonstrateGetAllUsers() {
	fmt.Println("\n2. 查询所有用户:")
	getAllUsers()
}

// 演示根据ID查询用户
func demonstrateGetUserByID() {
	fmt.Println("\n3. 根据ID查询用户:")
	getUserByID(1)
}

// 演示根据姓名查询用户
func demonstrateGetUserByName() {
	fmt.Println("\n4. 根据姓名查询用户:")
	getUserByName("张三")
}

// 演示更新用户信息
func demonstrateUpdateUser() {
	fmt.Println("\n5. 更新用户信息:")
	updateUser(1, "张三丰", 35)
}

// 演示更新后查询所有用户
func demonstrateGetAllUsersAfterUpdate() {
	fmt.Println("\n6. 更新后查询所有用户:")
	getAllUsers()
}

// 演示创建用户详情
func demonstrateCreateUserDetails() {
	fmt.Println("\n7. 创建用户详情:")
	createUserDetail(1, "zhangsan@example.com")
	createUserDetail(2, "lisi@example.com")
}

// 演示查询用户及其详情
func demonstrateGetUsersWithDetails() {
	fmt.Println("\n8. 查询用户及其详情:")
	getUserWithDetail(1)
	getUserWithDetail(2)
}

// 演示软删除用户
func demonstrateSoftDeleteUser() {
	fmt.Println("\n9. 软删除用户:")
	deleteUser(2)
}

// 演示查询所有用户（包括软删除的）
func demonstrateGetAllUsersIncludingDeleted() {
	fmt.Println("\n10. 查询所有用户（包括软删除的）:")
	var users []models.UserModel
	global.DB.Unscoped().Find(&users)
	fmt.Printf("查询到 %d 个用户（包括软删除的）:\n", len(users))
	for _, user := range users {
		deletedInfo := ""
		if user.DeletedAt.Valid {
			deletedInfo = " (已删除)"
		}
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d%s\n", user.ID, user.Name, user.Age, deletedInfo)
	}
}

// 演示查询未删除的用户
func demonstrateGetActiveUsers() {
	fmt.Println("\n11. 查询未删除的用户:")
	var activeUsers []models.UserModel
	global.DB.Find(&activeUsers)
	fmt.Printf("查询到 %d 个活跃用户:\n", len(activeUsers))
	for _, user := range activeUsers {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d\n", user.ID, user.Name, user.Age)
	}
}

// 演示永久删除用户
func demonstratePermanentlyDeleteUser() {
	fmt.Println("\n12. 永久删除用户:")
	permanentlyDeleteUser(3)
}

// 演示最终查询所有用户
func demonstrateFinalGetAllUsers() {
	fmt.Println("\n13. 最终查询所有用户:")
	getAllUsers()
}
