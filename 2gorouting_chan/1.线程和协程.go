package main

import "time"

// ✅线程的管理调度切换开销大（一个线程要占用好几M的内存空间）
// ✅协程就是将一段程序的运行状态打包，可以在线程之间调度
// ✅协程并不取代线程，协程也要在线程上运行

// ✅协程的优势：
// ✅资源利用，快速调度，超高并发（协程可以做到用几个线程，做到上千个并发）
func do1() {
	do2()
}
func do2() {
	do3()
}
func do3() {
	print("do do do...")
}
func main() {

	go do1()                    //加一个go关键字就是开启了一个协程
	time.Sleep(1 * time.Second) //防止主协程退出
}
