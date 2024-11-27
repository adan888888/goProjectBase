package main

import "fmt"

// 全局变量，定义在函数外的变量
var n7 = 100
var n8 = 200

// 设计者认为上面的全局变量的写法太麻烦了，可以一次性声明
var (
	n9  = 500
	n10 = "hello"
)

func main() {

	//1.变量的声明
	var age int
	//2.变量的赋值
	age = 18
	//3.变量的使用
	fmt.Println("age=", age)

	//声明和赋值可以合成一句：
	var age2 = 19
	fmt.Println("age2=", age2)

	//省略var, 就需要写 :=
	//var sex = "男"
	sex := "男"
	fmt.Println(sex)

	//声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var n4, n5, n6 = 10, "hello", 3
	fmt.Println(n4, n5, n6)

	n7, n8, n9 := 1, 2, 3 //省略var关键字
	fmt.Println(n7, n8, n9)

	fmt.Println(n9, n10)
}
