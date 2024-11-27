package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁的本质是将并行的代码串行化了，使用lock肯定会影响性能
// 即使是设计锁，那么也应该尽量的保证并行
// 假设有两组协程，其中一组负责写数据，另一组负责读数据，web系统中绝大部分场景都是读多写少，
// 虽然有多个goroutine,但是仔细分析会发现，读协程之间应该并发，读和写之间应该串行，读和读之间也不应该并行
// 读写锁
func main() {
	wg := sync.WaitGroup{}
	var num int
	var rwlock sync.RWMutex
	wg.Add(6)
	//写gorotine
	go func() {
		time.Sleep(3 * time.Second) //先让程序拿到读锁  3秒后 写锁工作（写锁也会防止，读锁加锁），所以停止5秒卡在读的锁那里，再继续执行后面程序
		defer wg.Done()
		rwlock.Lock() //加写锁，写锁会防止别的写锁获取和读锁获取
		defer rwlock.Unlock()
		num = 12
		time.Sleep(5 * time.Second)
		fmt.Println("get write lock:", num)
	}()
	//time.Sleep(time.Second)

	//读的gorotine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() //加读锁，读锁不会阻止别的锁
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock:", num)
				rwlock.RUnlock()
			}
		}()
	}

	wg.Wait()
}
