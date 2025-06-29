package main

import (
	"fmt"
)

func main() {
	//map就只有两种正确的初始化方法。 map、slice、channel 必须通过 make 或字面量初始化后才能使用！

	//var myMap map[string]string //会报panic: assignment to entry in nil map //map类型想要存值必须先初始化
	//myMap["name"] = "zs"
	//fmt.Println(myMap)

	/*第一种初始化方法*/
	//userInfo := map[string]string{} //没有数据的map，来初始化
	userInfo := map[string]string{"name": "小武", "age": "18"} //初始化的时候直接赋值

	userInfo["name"] = "小张"
	fmt.Println(userInfo)
	fmt.Println(userInfo)

	/*第二种初始化方法:使用make初始化一个map*/
	//make 是内置函数，主要用于初始化 slice ，map ，channel
	m := make(map[int]int /*, 10*/) //make([键类型] 值类型)    后面加个10，给十个位置 ，不写也是可以
	m[1] = 100
	m[2] = 200
	fmt.Println(m)

	/*new 函数的作用是分配内存并返回指针，但不会初始化对象。*/
	//错误的初始化方法：new来创建map    new出来的返回来是个指针(是map的一个指针)
	value := new(map[int]int)
	//value["k1"] = 100  报错（panic: assignment to entry in nil map）， 此时 value 为 &map(nil)，无法直接使用
	fmt.Println(value)

	data := make(map[int]int)
	data[1] = 999
	data[2] = 888
	fmt.Println(data)

	value = &data //data的指针指向value
	fmt.Println(value)
	fmt.Println(*value)      //指针*
	fmt.Println(len(*value)) //map的长度
	//fmt.Println(cap(data))   //map的容量   报错

	//删除
	//delete(data, 1)
	//fmt.Println(data)

	//循环
	dv := map[string]string{"n1": "张三", "n2": "李四"}
	for k, v := range dv {
		fmt.Println(k, v)
	}

	//一个参数的时候，只拿key
	for k := range dv {
		fmt.Println(k)
	}
	fmt.Println("-------------------")

	//map的集合也是无序的，而且不保证每次每次打印都是相同的顺序
	//只拿value
	for _, v := range dv {
		fmt.Println(v)
	}
	//如何判断到底是""还是没有这个key  （是否存在）
	s, ok := dv["java"] //有返回两个参数的
	fmt.Println(s, ok)  //ok返回是false

	//还可以写成一行
	if _, ok := dv["php"]; !ok {
		fmt.Println("php不存在")
	} else {
		fmt.Println("php存在")
	}
	//删除
	delete(dv, "n1")
	fmt.Println(dv)

	//很重要的提示， map不是线程安全的
	//var sx sync.Map //需要使用sync.Map
}
