package main

import "fmt"

func main() {
	//userInfo := map[string]string{} //空的map
	userInfo := map[string]string{"name": "小武", "age": "18"}

	userInfo["name"] = "小张"
	fmt.Println(userInfo)
	fmt.Println(userInfo)

	//使用make初始化一个map
	m := make(map[int]int /*, 10*/) //make([键类型] 值类型)    后面加个10，给十个位置 ，不写也是可以
	m[1] = 100
	m[2] = 200
	fmt.Println(m)

	//声明nil    能过new来创建map    new出来的返回来是个指针(是map的一个指针)
	value := new(map[int]int)
	//value["k1"] = 100  报错
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

	//只拿key
	for k := range dv {
		fmt.Println(k)
	}
	//只拿value
	for _, v := range dv {
		fmt.Println(v)
	}
}
