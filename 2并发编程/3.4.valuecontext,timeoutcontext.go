package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用context
func cupIInfo1(context context.Context) {
	fmt.Printf("tracid :%s \r\n", context.Value("traceid"))
	//记录一些日志，这次请求是哪个tracid打印的。排查的时候变的非常容易
	defer wg3.Done()
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

var wg3 sync.WaitGroup

func main() {
	wg3.Add(1)
	//2.主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//3.context.WithDeadline()//在时间点

	//4.context.withValue //在某个值

	ctx = context.WithValue(ctx, "traceid", "123")
	go cupIInfo1(ctx)
	//time.Sleep(6 * time.Second) 现在也不需要这个了
	wg3.Wait()
	fmt.Println("监控完成")
}
