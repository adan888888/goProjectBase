package main

import (
	"fmt"
	p1 "goProjectBase/package1" //可以要别名，也可以不要别名。不要别名，就是包名package1.
	xx "goProjectBase/package3"
)

func main() {
	p1.F1()
	fmt.Println(xx.N7)
}
