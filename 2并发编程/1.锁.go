package main

import (
	"fmt"
	"sync"
)

/*
锁(Mutex):
资源的竞争
*/
var total int32
var wg1 sync.WaitGroup

var lock sync.Mutex

// 锁能复制吗？复制后就会失去锁的效果
// atomic 原子锁 像加减的比较适合
// Mutex 可以锁逻辑（一段代码）
func add() {
	defer wg1.Done()
	for i := 0; i < 10000000; i++ {
		//atomic.AddInt32(&total, 1)
		lock.Lock()
		total += 1
		lock.Unlock()
	}

}
func sub() {
	defer wg1.Done()
	for i := 0; i < 10000000; i++ {
		//atomic.AddInt32(&total, -1)
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}
func main() {
	wg1.Add(2)
	go add()
	go sub()
	wg1.Wait()
	fmt.Println(total)

}
