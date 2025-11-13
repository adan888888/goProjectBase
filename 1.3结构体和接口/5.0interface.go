package main

import "fmt"

/*
  ✅在go语言中，struct实现interface都是隐式实现的，不像Java中要使用implement来实现接口
    可以在不改业务代码的情况下，提取接口（方法）。不用先定义接口再去实现。可以反着来：先写方法，再写接口。

  ✅go中是通过组合匿名字段来达到类似继承的效果

  ✅对过以上手段去掉了面向对象中复杂而冗余的部分

  ✅保留了基本的面向对象的特性
*/

//接口里面只能放函数(不能放变量,字段) 也不需要关键字fun

// 关键理解：  假如这个接口里面什么方法都没有，他将是任何结构体的父类，也就是any，任何类型都是interface
type Duck interface {
	//Name string 只有结构体才能放变量
	Gaga()
	Swiming()
}

// 结构1
type psDuck struct {
	//legs int
}

func (p *psDuck) Gaga() {
	fmt.Println("嘎嘎叫。。。")
}
func (p *psDuck) Swiming() {
	fmt.Println("游泳。。。")
}

// 结构2 都是实现了Duck这个接口
type pkDuck struct {
	//legs int
}

func (p *pkDuck) Gaga() {
	fmt.Println("p嘎嘎叫。。。")
}
func (p *pkDuck) Swiming() {
	fmt.Println("p游泳。。。")
}

// 还可以把接口放在结构里面，来实现多态（使用的时候传他具体的实现结构体）
type Factory struct {
	//duck Duck
	Duck //可以省略掉变量名称
}

func main() {
	//go语言中处处都是interface ，处处都是鸭子类型

	//var d Duck = &psDuck{}
	var d Duck = &pkDuck{} //也可以接收一个另外的一个结构体的实现，这就是多态
	d.Swiming()
}
