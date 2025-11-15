package main

import (
	"fmt"
	"time"
)

// ✅select 的作用主要用于多个channel
func main() {
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
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	defer fmt.Println("程序结束1！") //这个在2的后面。 越是最前，越是最后。
	defer fmt.Println("程序结束2！")
	for {
		select {
		case <-ch2:
			fmt.Println("g1 done")
		case <-ch1:
			fmt.Println("g2 done")
		case <-ticker.C: //通常我们不使用default而作用Ticker, 三秒钟一来，我们就结束
			time.Sleep(time.Second)
			fmt.Println("time out")
			return //会退出当前main()
			//  default:
			//	 fmt.Println("default")

		}
	}
}

// 两个方法向不同的channel写入数据
func g11(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{} //struct{}空结构体的定义 {}再实例化

}
func g22(ch chan struct{}) {
	time.Sleep(4 * time.Second)
	ch <- struct{}{}
}
