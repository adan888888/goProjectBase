package main

import (
	"fmt"
	"sync"
	"time"
)

/*
这是一个等待多个 goroutine 完成的案例，但实现方式不推荐。
*/
var done bool
var lock1 sync.Mutex //互斥锁,保护共享变量的并发访问

func main() {
	fmt.Println(done)
	//现在两个goroutine都在执行，但是我在主goroutine中当一个完成之后，这时候我会立马通知
	go g1()
	go g2()
	//主协程里面的一个死循环,卡死主协程 不让子协程退出
	for {
		if done {
			fmt.Println("done1", done)
			time.Sleep(10 * time.Millisecond) //防止消耗太多的内存
			return                            //return会退出main这个方法  ，break的话，只会退出这个for循环
		}
	}
	fmt.Println("程序结束！")
}

func g1() {
	time.Sleep(10 * time.Second)
	fmt.Println("g1")
	lock1.Lock()
	defer lock1.Unlock()
	done = true

}
func g2() {
	time.Sleep(2 * time.Second)
	fmt.Println("g2")
	lock1.Lock()
	defer lock1.Unlock()
	done = true
}
