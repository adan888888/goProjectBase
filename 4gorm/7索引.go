package main

import (
	"4gorm/global"
	"4gorm/models"
	"fmt"
	"time"
)

func main() {
	// 初始化数据库连接
	global.InitDB()

	// 自动迁移数据库表
	//global.MigrateDB()

	//-- INSERT INTO user_test_index (`username`,`email`,`age`,city,phone) SELECT username,email,age,city,phone FROM user_test_index  #把所有的数据复制一遍在插入
	//-- SELECT COUNT(*) FROM user_test_index  #统计一共有多少条数据
	//SELECT * FROM user_test_index WHERE username="张三"
	//SHOW INDEX FROM user_test_index  #查看索引
	//CREATE INDEX index_username ON user_test_index(username) #给username这个字段设置索引
	//DROP INDEX index_username ON user_test_index #删除索引

	//数据多的时候才看的出来索引的作用
	performanceTest()
}

// 性能测试函数
func performanceTest() {
	// 测试索引查询性能
	testIndexPerformance()
}

// testIndexPerformance 测试索引查询性能
func testIndexPerformance() {
	fmt.Println("测试索引查询性能...")

	// 测试用户名索引
	start := time.Now()
	var users []models.UserTestIndex
	//SELECT * FROM `user_test_index` WHERE username = '张三'
	global.DB.Unscoped().Where("username = ?", "张三").Find(&users)
	usernameTime := time.Since(start)
	fmt.Printf("用户名索引查询耗时: %v，查询到 %d 条记录\n", usernameTime, len(users))

}
