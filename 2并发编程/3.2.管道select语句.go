package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var lock1 sync.Mutex

func main() {
	//现在两个goroutine都在执行，但是我在主goroutine中当一个完成之后，这时候我会立马通知
	go g1()
	go g2()
	for {
		if done {
			fmt.Println("done1")
			time.Sleep(10 * time.Millisecond) //防止消耗太多的内存
			return                            //return会退出main这个方法  ，break的话，只会退出这个for循环
		}
	}
	fmt.Println("程序结束！")
}

func g1() {
	time.Sleep(time.Second)
	lock1.Lock()
	defer lock1.Unlock()
	done = true

}
func g2() {
	time.Sleep(2 * time.Second)
	lock1.Lock()
	defer lock1.Unlock()
	done = true
}
