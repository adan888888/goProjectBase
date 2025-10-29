package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	ch1 := make(chan string, 2)

	//一。读取协程
	go func(ch1 chan string) { //✅用协程配合管道，才不会死锁。 直接读里面没有数据的时候，就报...
		//管道是安全的不会没有数据，就读取报错（可以一边写入一边读取，还可以等待写入）
		for data := range ch1 { //会阻塞在这里，直到close
			fmt.Println(data)
		}
		fmt.Println("all done1 ")
		wg.Done()
	}(ch1)

	//一。写入协程
	go w(ch1)

	//ch1 <- "hello3" //已经关闭的就不能再放值了，但是关闭的可以再取值
	//time.Sleep(time.Second * 3) //让协程走完，代替sync.WaitGroup
}
func w(ch1 chan string) {
	ch1 <- "hello1"
	ch1 <- "hello2"
	time.Sleep(time.Second * 3)
	ch1 <- "hello3"
	time.Sleep(time.Second * 3)
	close(ch1) //关闭后 上面for range就会退出 打印 all done1
}
