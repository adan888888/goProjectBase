package main

import (
	"container/list"
	"fmt"
)

func main() {
	//切片是本身是一个动态数据的概念
	//list是一个链表的数据结构 （查询比较慢，插入比较快）

	var myList list.List
	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("mysql")

	//从头到尾
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
