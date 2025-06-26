package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//%d 十进输出
	//%s 字符串原样输出
	userName := "bobby"
	age := 18
	address := "北京"
	mobile := "18978789999"
	fmt.Println("用户名："+userName, "，年龄："+strconv.Itoa(age), "，地址："+address, "，电话："+mobile) //极难维护
	fmt.Printf("用户名：%s，年龄：%d，地址：%s，电话：%s\n", userName, age, address, mobile)              //这个非常的常用，但是性能
	usrMsg := fmt.Sprintf("用户名：%T，年龄：%T，地址：%s，电话：%s\n", userName, age, address, mobile)   //这个性能最差
	fmt.Println(usrMsg)

	//通过string的builder进行字符串拼接，高性能
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(userName)
	builder.WriteString("，年龄：")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString("，地址：")
	builder.WriteString(address)
	fmt.Println(builder.String())

	/////////////////////////// 2.字符串的比较//////////////////////////////////////////////////////////////////////////
	a := "hello"
	b := "bello"
	fmt.Println(a == b) //字符串的比较

	//字符串的大小比较
	fmt.Println(a > b) //比较的是首字母的asci码的大小

	/////////////////////////// 3.strings//////////////////////////////////////////////////////////////////////////
	name := "你好-go工程师课程"
	fmt.Println(strings.Contains(name, "go"))       //是否包含
	fmt.Println(strings.Count(name, "go"))          //出现次数
	fmt.Println(strings.Split(name, "-")[0])        //分割
	fmt.Println(strings.Trim("#$hell #go#$", "#$")) //去掉前后的

}
