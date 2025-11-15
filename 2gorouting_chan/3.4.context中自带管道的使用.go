package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
监控cpu
使用context实现优雅的退出程序
*/

/* 普通的
func cupIInfo(stop chan struct{}) {
	defer wg2.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出cup监控")
			return
		default:
			fmt.Println("cup信息")
			time.Sleep(2 * time.Second)
		}
	}
}
*/
//使用context
func cupIInfo(context context.Context) {
	defer wg2.Done()
	for {
		select {
		case <-context.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			fmt.Println("cup信息")
			time.Sleep(2 * time.Second)
		}
	}
}

var wg2 sync.WaitGroup

/*
监控cpu
使用context实现优雅的退出程序
*/
func main() {
	//在主线程中的是主groutine
	wg2.Add(1)

	//context包提供了三种函数，withCancel ,withTimeout,WithValue
	context1, cancel := context.WithCancel(context.Background())
	context2, _ := context.WithCancel(context1) //会有传递性
	//var stop = make(chan struct{})
	go cupIInfo(context2) //使用context2也可以，因为有传递
	time.Sleep(6 * time.Second)
	//stop <- struct{}{} //1秒后给子协程发一个停止消息。

	cancel() //一调用取消方法，就会发送context.Done()的协程消息，
	wg2.Wait()
	fmt.Println("监控完成")
}
