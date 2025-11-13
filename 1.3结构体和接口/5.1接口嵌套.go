package main

import "fmt"

type MyWriter interface {
	Write(string)
}
type MyReader interface {
	Reader() string
}
type MyReaderWriter interface {
	MyReader // 变量类型  -->结构体和接口里面，都可以直接放类型，把变量名省略掉
	MyWriter
	ReadWriter()
}
type SreadWriter struct {
}

func (s *SreadWriter) Reader() string {
	fmt.Println("Reader...")
	return "Reader..."
}

func (s *SreadWriter) Write(s2 string) {
	fmt.Println("Write...")
}

func (s *SreadWriter) ReadWriter() {
	fmt.Println("ReadWriter...")
}

func main() {
	var w MyReaderWriter = /*接口类型*/ &SreadWriter{} //指针对象的接收者必须使用 & ，值对象的接收者两个都可以使用
	w.Reader()
}
