package main

import (
	"fmt"
	"time"
)

func main() {

	round := 0
	for {
		time.Sleep(1 * time.Second)
		round++
		if round == 5 {
			continue //跳过本次循环，直接进行下一次循环
		}
		if round > 10 {
			break //退出for循环
		}
		fmt.Println(round)
	}
}
