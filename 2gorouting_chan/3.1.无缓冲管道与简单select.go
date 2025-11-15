package main

import (
	"fmt"
	"time"
)

/*
*
1.无缓冲 channel //✅✅✅发送和接收必须配对，否则会阻塞（一个发就是一个接收） ❌否则会出现死锁。
2.select //select 更适合多 channel 监听
*/
func main() {
	var chanMsg chan string
	//1.无缓冲channel适用于 通知，B要第一时间知道A是否已经完成
	//2.有缓冲channel适用于消费者和生产者之间的通讯  ，比如爬虫：生产者不断的往里面放url，消费者不断的从里面读取到url进行数据抓取
	chanMsg = make(chan string, 0) //无缓冲的的channel（就是容量为0的管道）

	//读取1------------------------------------
	go func(msg chan string) { //go有一种happen-before的机制，可以保障
		fmt.Println("协程1开始1")
		select {
		case data := <-msg: //<-msg从管道里面读取--监听通道
			//卡在这里等待 从管道里读取值
			fmt.Println("select读取=", data)
		}
	}(chanMsg)
	time.Sleep(2 * time.Second)
	fmt.Println("...主协程")
	chanMsg <- "hello1" // 第1次：同步传递 //(chanMsg <-)向管道里面写入数据，一写入数据子协程马上会读取(必须要发一次，马上接一次)

	//读取2------------------------------------
	go func() {
		fmt.Println("协程2开始2")
		v := <-chanMsg //阻塞
		fmt.Println("直接读取=", v)
	}()
	chanMsg <- "hello2" // 第2次：同步传递

	time.Sleep(time.Second * 3)

}
