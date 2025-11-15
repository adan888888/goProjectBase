package main

import "fmt"

func caseNew() {
	//用于分配值类型的零值内存，并返回指向该内存的指针
	p := new(int)
	str := new(string)
	s := new(struct{})
	fmt.Println(p, str, s)
	fmt.Println(*p, *str, *s)
	fmt.Println("---------------------------------------------------------------")
}

func caseMake() {
	s1 := make([]int, 5, 10)   //创建一个长度为5，容量为10的切片，⚠️切片这里长度为5，会先存5个0进去。
	m2 := make(map[string]int) //创建一个空的map
	ch := make(chan int, 5)    //创建一个带缓冲区的通道，缓冲区大小为5,这个比较特殊指向该值类型的地址
	fmt.Println(s1, m2, ch)
}
func main() {
	caseNew()
	caseMake()
}
