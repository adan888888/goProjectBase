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

	//var singeTable = models.SingleTable{
	//	Name:  "李四1",
	//	Age:   24,
	//	Adres: "北京",
	//}
	//global.DB.Create(&singeTable)

	//更新
	//更新有很多方法， Save，Update，UpdateColumn，Updates

	//1.有组件记录就是更新，并且可以更新❗️零值。否则就是创建
	//SaveTest()

	//2
	updateTest()
}

func updateTest() {
	var t1 = models.SingleTable{ID: 1}
	//global.DB.Model(t1).Update("age", 0) //零值也能更新
	global.DB.Model(t1).Update("age", 24) //只更新某一个字段
	//global.DB.Model(t1).UpdateColumn("age", 23) //不会走钩子函数
	fmt.Println(t1) //不回填数据
}

func SaveTest() {
	var t1 models.SingleTable
	t1.ID = 2
	t1.Age = 100
	t1.Name = "枫2"
	t1.CreateTime = time.Now()
	global.DB.Save(&t1)
	fmt.Println(t1)
}
