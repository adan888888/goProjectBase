package main

import "fmt"

func main() {

	//中文的是星期几，输出对应的英文
	day := "星期三"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	case "星期三":
		fmt.Println("Wednesday")
	default:
		fmt.Println("其它")
	}

	//比其它的语言更加灵活的
	score := 60
	switch {
	case score < 60:
		fmt.Println("E")
	case score >= 60 && score < 70:
		fmt.Println("D")
	case score >= 80 && score < 90:
		fmt.Println("B")
	case score >= 90 && score < 100:
		fmt.Println("A")

	default:
		fmt.Println("差")
	}

	//比其它的语言更加灵活的
	switch score {
	case 60, 70, 80:
		fmt.Println("A")
	}
}
