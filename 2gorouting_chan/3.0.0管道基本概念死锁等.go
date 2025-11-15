package main

import "fmt"

/*管道就像水管，遵循先进先出*/
func main() {
	//1.创建管道
	ch := make(chan int, 3)

	//2.给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3.获取管道里面的内容
	a := <-ch
	fmt.Println(a) //10

	<-ch //21 从管道里面取值，可以不赋值给任何变量

	c := <-ch
	fmt.Println(c) //32

	ch <- 56 //管道又空了，又可以存储数据了

	//4.打印管道的容量和长度
	fmt.Printf("值%v 容量%v 长度%v\n", ch, cap(ch), len(ch)) //值0x14000090080 容量3 长度1。(只有个值)
	fmt.Println("---------------------------------------------------")

	//5.管道的类型（✅是引用类型，❌不是值类型）
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) //69 说明是引用类型
	fmt.Println("---------------------------------------------------")

	//6.管道死锁
	/*ch6 := make(chan int, 1)
	ch6 <- 34
	ch6 <- 35 //all goroutines are asleep - deadlock!  1.❌超过容量还在存，死锁 造成阻塞*/

	/*ch7 := make(chan string, 2) //..eep - deadlock！ 2.❌取完了，还在取 也会造成死锁
	ch7 <- "数据1"
	ch7 <- "数据2"
	m1 := <-ch7
	m2 := <-ch7
	m3 := <-ch7
	fmt.Println(m1, m2, m3) */

	//7.管道的遍历
	var ch8 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch8 <- i
	}

	//1.通过for rang循环遍历管道的时候， ✅一定要先关闭管道 3.❌不先关闭 也会造成死锁
	//close(ch8) // ✅数据写完一定要关闭
	//for v := range ch8 {
	//	fmt.Println(v)
	//}
	fmt.Println(" for循环---------------------------------------------------")

	//2.通过for就不需要   4.❌for循环取完了再取 也会造成死锁
	for i := 0; i < 10; /*11*/ i++ {
		fmt.Println(<-ch8)
	}
	fmt.Println("---------------------------------------------------")
}
