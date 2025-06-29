package main

import "fmt"

type PersonA struct {
	name string
	age  int
}

// 接收器有两种形态(p *PersonA)指针传递 和 (p PersonA)
func (p *PersonA) printX() {
	//1.有可能该函数中想改变p的值， 2.person对象很大
	p.age = 20
	fmt.Printf("name: %s ,age :%d \n", p.name, p.age)
}

func main() {
	p := PersonA{"zs", 18}
	p.printX()
	fmt.Println(p.age) //当(p *PersonA)这里是指针时 p.age=20 ,相当于引用传递
}
