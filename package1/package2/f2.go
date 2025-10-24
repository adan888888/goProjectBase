package package2

import (
	"fmt"
	"goProjectBase/package1"
)

func F2() {
	package1.F1()
	fmt.Println("F2")
}
