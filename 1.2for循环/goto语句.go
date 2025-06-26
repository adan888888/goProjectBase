package main

import "fmt"

// goto 语句可以让我们的代码跳到指定的代码块中运行
func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 2 {
				//break //只能跳出当前for循环
				goto over //跳到指定的代码块中运行(两层for循环都退出了)
				//return //使用return也能跳出两层for循环
			}
			fmt.Println(i, j)
		}
	}

over:
	fmt.Println("over")
}
