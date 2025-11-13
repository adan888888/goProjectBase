package main

import "fmt"

func mPrint(datas ...interface{}) { //... 出现在函数参数列表的最后一个参数类型前时，表示该函数可以接受任意数量的该类型参数
	for _, value := range datas {
		fmt.Println(value)
	}
}
func mPrint2(datas interface{}) {
	fmt.Println(datas)
}

type myinfo struct{}

// 这里说明只要方法名字和接口里面方法一样就是实现他的接口，绑定哪个结构体都一样
func (mi *myinfo) Error() string {
	return "我不是error"
}
func main() {
	//var data = []interface{}{"zs", 2, 3}
	//mPrint(data) //这样传进去是没有问题的

	var data = []string{"zs", "s", "ls"} //切片类型
	//mPrint2(data) //这样放进去运行就会报错

	mPrint2(2)

	//var err error
	//err = &myinfo{}

	var datai []interface{}
	for _, value := range data {
		datai = append(datai, value)
	}
	mPrint(datai...) //等价于 mPrint("zs", "s", "ls")
}
