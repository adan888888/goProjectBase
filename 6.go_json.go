package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

var student *Student

func main() {
	m := map[string]any{"name": "张三", "age": 18}
	if j, err := json.Marshal(m); /*map对象转json*/ err != nil {
		fmt.Println("转换为json错误", err.Error())
	} else {
		fmt.Println(string(j))      //转过来是byte字节数组，所以还需要string(j)转为字符串
		json.Unmarshal(j, &student) //json转结构体
	}
	fmt.Println("Name:", student.Name, "Age:", student.Age)
}
