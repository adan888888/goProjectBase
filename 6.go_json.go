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
	j, err := json.Marshal(m) //map对象转json
	if err != nil {
		fmt.Println("转换为json错误")
	}
	fmt.Println(string(j))
	json.Unmarshal(j, &student) //json转结构体
	fmt.Println("Name:", student.Name, "Age:", student.Age)
}
