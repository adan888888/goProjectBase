package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
锁(Mutex):
资源的竞争
*/
var total int32
var wg1 sync.WaitGroup

var lock sync.Mutex //互斥锁

// 锁能复制吗？复制后就会失去锁的效果
// atomic 原子锁 像加减的比较适合
// Mutex 可以锁逻辑（一段代码）
func add() {
	defer wg1.Done() // 协程计数器减1
	for i := 0; i < 10000000; i++ {
		//atomic.AddInt32(&total, 1)
		lock.Lock()
		total += 1
		lock.Unlock() //其它的协程要等待unlock以后才能运行
	}

}
func sub() {
	defer wg1.Done() // 协程计数器减1
	for i := 0; i < 10000000; i++ {
		//atomic.AddInt32(&total, -1)
		lock.Lock()
		total -= 1
		lock.Unlock()
	}
}
func main() {
	wg1.Add(2) //协程计数器 加2
	go add()
	go sub()
	wg1.Wait()         //等待协程计数器执行完毕
	fmt.Println(total) //0 不加锁最后输出的可能不为0

	numCPU := runtime.NumCPU()
	fmt.Println("CPU个数", numCPU)
}
