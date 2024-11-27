package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string, 2)
	go func(msg chan string) {
		for data := range msg { //如果是for each 是需要知道定长的
			fmt.Println(data)
		}
		fmt.Println("all done1 ")
	}(msg)

	msg <- "hello"
	msg <- "hello1"
	close(msg) //关闭后 上面for range就会退出 打印 all done1
	//msg <- "hello3" //已经关闭的就不能再放值了，但是关闭的可以再取值
	time.Sleep(time.Second * 3)
}
