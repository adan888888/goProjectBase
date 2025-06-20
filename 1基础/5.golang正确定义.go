package main

import (
	"fmt"
	_ "fmt" //下划线在go中是引用而不使用
	"strconv"
	"unicode/utf8"
)

func main() {

	////////////////////////1.变量定义////////////////////////////////////////////////////////////////////////////
	var name string = "这里是名称"
	var arr [10]int
	var slice []string
	var mp map[string]int

	fmt.Println(name, arr)
	if slice == nil {
		fmt.Println("slice is nil")
	}
	if mp == nil {
		fmt.Println("map is nil")
	}

	/////////////////////////////2.下划线的使用////////////////////////////////////////////////////////////////////
	_, err := get() //不使用可以用下划线
	fmt.Println(err)

	///////////////////////////// 3.代码块//////////////////////////////////////////////////////////////////////
	x := 1
	fmt.Println(x)
	{ //大括号圈定了代码块中声明的变量的作用域
		fmt.Println(x) //1
		x := 2
		fmt.Println(x) //2
	}
	fmt.Print(x) //1

	///////////////////// 4.不能使用nil的///////////////////////////////////////////////////////////////////////////
	//xx := nil //error
	var xx interface{} = nil
	fmt.Println(xx)

	/////////////////////5.类型转化///////////////////////////////////////////////////////////////////////////
	var num1 int = 100
	fmt.Println(num1) //100

	//数值类型之间的可以强转
	var num2 int64 = 100
	fmt.Println(float64(num2))

	//数字(int)转字符串
	var nmu3 int = 100
	fmt.Println(strconv.Itoa(nmu3) + "abc") //100abc

	//字符串转数字(int)
	var str1 string = "100abc"
	fmt.Println(strconv.Atoi(str1)) //100 <nil>  /  0 strconv.Atoi: parsing "100abc": invalid syntax

	/**
	  2.字符串转基本类型
	*/
	//字符串转换为float32，转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	//字符串转换为int
	fmt.Println(strconv.ParseInt("10", 10, 64))

	/**
	  3.基本类型转字符串
	*/
	//布尔值转string
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	//int64转字符串
	var num4 int64 = 1010
	fmt.Println(strconv.FormatInt(num4, 10)) //1010

	//int64转字符串
	var ff float64 = 3.1415926
	fmt.Println(strconv.FormatFloat(ff, 'f', -1, 64))

	//字符串与[]byte转化 （字符串与切片）
	var str3 string = "今天天气很好"
	fmt.Println([]byte(str3)) //[228 187 138 229 164 169 229 164 169 230 176 148 229 190 136 229 165 189]

	var bytes1 = []byte{228, 187, 138, 229, 164, 169, 229, 164, 169, 230, 176, 148, 229, 190, 136, 229, 165, 189}
	fmt.Println(string(bytes1))

	//字符串与字符
	//https://www.bilibili.com/video/BV17hHseLEiU/?p=4&spm_id_from=pageDriver&vd_source=55f7073cc1049edc8b91cea83217e7b6

	/////////////////////////// 6.接口转具体类型//////////////////////////////////////////////////////////////////////////
	var inf interface{} = "100abc"
	t, ok := inf.(int) //转成int 通过断言
	fmt.Println(t, ok) //100 false

	/////////////////////////// 7.引用包//////////////////////////////////////////////////////////////////////////
	//1.包名是文件夹名字
	//2.文件的名字，不一定是包名

	/////////////////////////// 8.//////////////////////////////////////////////////////////////////////////
	str1 = "ABC"
	fmt.Println(utf8.ValidString(str1)) //难证是否是utf8
	str1 = "A\\n fcC"
	fmt.Println(utf8.ValidString(str1))
}
func get() (res int, err error) {
	fmt.Println("call get")
	return 1, nil
}
