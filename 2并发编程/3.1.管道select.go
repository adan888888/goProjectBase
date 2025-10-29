package main

import (
	"fmt"
	"time"
)

/* ✅✅✅正确理解： 无缓冲 channel 像一个"传话筒"，容量为 0，但可以反复使用，每次都是一对一同步传递。*/
func main() {
	var msg chan string
	//1.无缓冲channel适用于 通知，B要第一时间知道A是否已经完成
	//2.有缓冲channel适用于消费者和生产者之间的通讯  ，比如爬虫：生产者不断的往里面放url，消费者不断的从里面读取到url进行数据抓取
	//创建一个管道存string类型的数据
	msg = make(chan string, 0) //无缓冲的的channel（就是容量为0的管道）

	//------------------------------------
	go func(msg chan string) { //go有一种happen-before的机制，可以保障
		fmt.Println(".....1")
		select {
		case data := <-msg: //<-msg从管道里面读取--监听通道
			//卡在这里等待 从管道里读取值
			fmt.Println(data)
		}
	}(msg)
	time.Sleep(2 * time.Second)
	fmt.Println(".....2")
	msg <- "hello1" // 第1次：同步传递 //(msg <-)向管道里面写入数据，一写入数据子协程马上会读取(必须要发一次，马上接一次)

	//------------------------------------
	go func() {
		v := <-msg
		fmt.Println(v)
	}()
	// 可以反复使用
	msg <- "hello2" // 第2次：同步传递

	//waitgroup 如果少了done调用，容易出现deadlock,无缓冲的channel也容易出现
	time.Sleep(time.Second * 5)

}
