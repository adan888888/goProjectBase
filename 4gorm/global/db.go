package global

import (
	"4gorm/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	dsn := "root:mima123@tcp(localhost:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,                               // 禁用迁移时的外键约束，手动创建
		Logger:                                   logger.Default.LogMode(logger.Info), // 启用SQL日志打印
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	fmt.Println("数据库连接成功!")
}

// 自动迁移数据库表
func MigrateDB() {
	// 先删除现有的表（如果存在）
	//dropExistingTables()

	err := DB.AutoMigrate(
		//&models.UserModel{},
		//&models.UserDetailModel{},
		//&models.DepartmentModel{},
		//&models.EmployeeModel{},
		&models.StudentModel{},
		&models.CourseModel{},
	)
	if err != nil {
		log.Fatal("创建用户表失败:", err)
	}

	fmt.Println("数据库表创建/更新成功!")
}

// 删除现有的表
func dropExistingTables() {
	// 删除表（按依赖顺序）
	DB.Exec("DROP TABLE IF EXISTS employee_models")
	DB.Exec("DROP TABLE IF EXISTS department_models")

	fmt.Println("已删除现有表")
}
