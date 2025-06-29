package main

import "fmt"

// 1.结构体和接口里面 只能定义字段不是变量，也不能放函数
// 2.允许直接嵌入类型（不指定字段名）-->匿名字段
type person struct {
	name    string
	age     int
	address string
}
type student struct {
	person //匿名字段 用于实现类似继承的行为
	score  int
}

func main() {
	//匿名结构体,一次性使用，也不用给名字
	person := struct {
		name    string
		age     int
		address string
	}{"lx", 30, "北京市..."}
	fmt.Println(person.age)
}
