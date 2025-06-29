package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	//p      Person //第一种嵌套方式
	Person //第二种嵌套方式 -直接写个类型，变量都不要了
	score  float32
}

func main() {

	//s := Student{
	//	Person{
	//		name: "zs",
	//		age:  0,
	//	}, 100,
	//}

	s := Student{}
	/*s.p.name = "bobby"
	fmt.Println(s.p.name)*/

	//第二种方式,
	s.name = "boddy"    //可以直接赋值（在第一层）
	fmt.Println(s.name) //取值也是在第一层

	//但是在初始化的时候，不能写在一层
	//x:=Student{
	//	name:"zs"
	//	score: 98,
	//}

	//如何外层和里屋都有相同的字段 会覆盖里面一层的

}
