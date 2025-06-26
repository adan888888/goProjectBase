package main

import (
	"fmt"
	"time"
)

func main() {
	//for init;condition;post{}
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	//这就相当于while ，go中没有while
	var i int
	for i < 3 {
		time.Sleep(time.Second)
		fmt.Print(i)
		i++
	}
	fmt.Println("------------------------------------------")
	//for循环还有一种用法， for range， 主要是对 字符串，数组，切片 ，map， channel
	name := "imooc go"
	//for index, value := range name {
	//	//fmt.Println(index, value) //打印的是ASCII值
	//	fmt.Printf("%d, %c\r\n", index, value) //%c是ASCII 码对应的字符
	//}
	/*
			字符串  如果不写key，那么返回的是索引
			数组 如果不写key，那么返回的是索引
		    切片 如果不写key，那么返回的是索引
		    map 如果不写key，那么返回的是map的值
	*/
	//for rang key，value

	name = "imooc go你好"
	for _, value := range name { //如果有中文的话，必须这样写（key，value都要写，把key写成_）帮我们处理好这些细节了
		fmt.Printf("%c\n", value) //打印的是ASCII值
	}
}
