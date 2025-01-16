package test

var Name = "zhangSan"
var Age = 20

func init() {
	println("this is an initial demo")
}

// MyFunc 首字母大写：导出，可以被其他包访问。
// 首字母小写：未导出，仅能在当前包内访问。
func MyFunc() string {
	address := "China Beijing"
	return address
}

func Add() int {
	x := 10
	y := 20
	return x + y
}
