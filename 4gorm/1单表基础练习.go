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
	SaveTest()

	//2
	//updateTest()

	//3
	//UpdatesTest()
}

func UpdatesTest() {
	var table = models.SingleTable{ID: 5}
	//UPDATE `single_tables` SET `name`='高高',`age`=18 WHERE `id` = 5
	//global.DB.Model(&table).Updates(&models.SingleTable{Name: "高高", Age: 0}) //不会更新零值(就是原来表里面是什么还是什么) 也会走钩子函数
	global.DB.Model(&table).Updates(map[string]any{"Name": "高高", "Age": 0}) //传map会更新零值
	fmt.Println(table)                                                      //不会回值数据（仅填，我在代码里传过去的值。 数据库里面的不会回过来）

	//‼️重要。如果想使用Updates 又要传对象进去 又想更零值 可以使用
	//1..Select("name", "age", "adres").Updates(user)
	//2...Select("*")
	//3...Omit("id", "create_time") Omit 排除不需要更新的字段
}

func updateTest() {
	var t1 = models.SingleTable{ID: 4} //通过这个id来更新
	//global.DB.Model(t1).Update("age", 0) //零值也能更新
	global.DB.Model(t1).Update("age", 24) //只更新某一个字段（正常会走钩子函数，但是这里好像有冲突。加了钩子函数就报错）
	//global.DB.Model(t1).UpdateColumn("age", 230) //不会走钩子函数和Update的唯一的区别
	fmt.Println(t1) //不回填数据
}

func SaveTest() {
	var t1 models.SingleTable
	t1.ID = 2
	t1.Age = 22
	t1.Name = "" //save 会更新零值
	t1.CreateTime = time.Now()
	global.DB.Save(&t1)
	fmt.Println(t1)
}
