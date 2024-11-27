package main

import (
	"fmt"
	"time"
)

var done1 = make(chan struct{}) //结构体的channel 是多线程安全的   //一定要make(chan struct{})  make初始化要不然报死锁错误
func main() {
	//select 的作用主要用于多个channel

	//现在两个goroutine都在执行，但是我在主goroutine中当一个完成之后，这时候我会立马通知
	var ch1 = make(chan struct{})
	var ch2 = make(chan struct{})
	go g11(ch1)
	go g22(ch2)
	/*data := <-ch1
	data1 := <-ch2
	fmt.Println("第一个channel", data)
	fmt.Println("第二个channel", data1)*/ //再这样写，就会卡一个速度快channel

	//	我们要监控多个channel，任何一个channel返回都知道
	tc := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ch2:
			fmt.Println("g1 done")
		case <-ch1:
			fmt.Println("g2 done")
		case <-tc.C: //通常我们不使用default而作用Ticker, 三秒钟还没有来，我们就结束
			time.Sleep(time.Second)
			fmt.Println("time out")
			return
			//  default:
			//	 fmt.Println("default")

		}
	}
	fmt.Println("程序结束！")
}

//两个方法向不同的channel写入数据

func g11(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{} //struct{}空结构体的定义 {}再实例化

}
func g22(ch chan struct{}) {
	time.Sleep(4 * time.Second)
	ch <- struct{}{}
}
