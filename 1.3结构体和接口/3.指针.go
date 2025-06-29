package main

import "fmt"

type PersonB struct {
	name string
}

func changeName(p *PersonB) { //传指针就是值传递
	p.name = "aaa"
}
func main() {
	// 指针，
	p := PersonB{
		"zs",
	}
	changeName(&p)      //&这个符号 是取地址的符号就是指针
	fmt.Println(p.name) //会改变p里面的值

	po := &PersonB{
		"zs",
	}
	(*po).name = "boddy3" //*po 取指针的值
	po.name = "boddy4"
	po = &p
	fmt.Println(po)

	//给指针赋值
	//var pi *PersonB //指针必须初始化
	//ps := &PersonB{} //第一种初始化方式
	//fmt.Println(b)
}
