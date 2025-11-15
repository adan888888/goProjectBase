package main

import (
	"fmt"
	"sync"
	"time"
)

/*
写入一个，读取一个
*/

var wg1 sync.WaitGroup

func main() {
	wg1.Add(1)
	defer wg1.Wait()
	ch1 := make(chan string, 2)

	//🩸一。读取协程
	go func(ch1 chan string) { //✅用协程配合管道，才不会死锁。 直接读里面没有数据的时候，就报...
		//管道是安全的，不会没有数据，就读取报错（可以一边写入一边读取，还可以等待写入）
		for data := range ch1 { //会阻塞在这里，直到close
			fmt.Println(data)
		}
		fmt.Println("all done1 ")
		wg1.Done()
	}(ch1)

	//🩸二。写入协程
	go write(ch1)

	//ch1 <- "hello3" //已经关闭的就不能再放值了，但是关闭的可以再取值
}
func write(ch1 chan string) {
	ch1 <- "hello1"
	ch1 <- "hello2"
	time.Sleep(time.Second * 3)
	ch1 <- "hello3"
	time.Sleep(time.Second * 10)
	close(ch1) //关闭后 上面for range就会退出 打印 all done1
	/*
		总结：for range 需要 channel 被关闭才会退出；
		如果不关闭，读取协程会一直阻塞，导致主协程也一直等待，从而死锁。 4.❌不关闭 也会造成死锁
	*/
}
