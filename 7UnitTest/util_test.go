package _test

import (
	"fmt"
	"goProjectBase/7UnitTest/hahahaha"
	"testing"
)

// 1.文件名必须要 _test结尾
// 2.方法接收的参数必须是(t *testing.T)
// 3.go test [-v] 或者  go test -run Test2（只运行某一个方法）
func Test1(t *testing.T) {
	if hahahaha.HahaFun() != 1 {
		t.Error("min()=", hahahaha.HahaFun())
	}
}
func Test2(t *testing.T) {
	fmt.Println("我是test2")
}
func Test3(t *testing.T) {
	fmt.Println("我是test3")
}
func Test4(t *testing.T) {

}
func Test5(t *testing.T) {

}
