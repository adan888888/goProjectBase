package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("开始程序")
	quit := make(chan os.Signal, 1)
	//// 监听 os.Interrupt (Ctrl+C) 和 syscall.SIGTERM (终止信号)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) //程序运行时监听指定的系统信号，特别是 os.Interrupt 信号（通常由用户按下 Ctrl+C 产生），从而实现程序的优雅退出或其他信号处理操作。

	fmt.Println("程序正在运行。按 Ctrl+C 退出...")

	// 模拟程序正在运行的情况
	go func() {
		for {
			fmt.Println("协程工作中...")
			time.Sleep(time.Millisecond * 1200) //1500毫秒 1.5秒
		}
	}()

	t := <-quit //程序会阻塞在这里, // 等待接收到退出信号

	if t != nil {
		defer os.Exit(0)
	}
	print("收到信号，程序优雅退出。")
}
