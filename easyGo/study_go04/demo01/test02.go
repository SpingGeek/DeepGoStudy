package main

import "fmt"

// 在匿名函数中的数据类型只能是基本数据类型
type intro struct {
	string
	int
	float64
	bool
}

func main() {
	info := intro{"Jerry", 23, 181, false}
	fmt.Println(info.string)
	fmt.Println(info.int)
	fmt.Println(info.float64)
	if info.bool {
		fmt.Print("have the experience of work")
	} else {
		fmt.Println("do not have the experience of work")
	}
	info01 := intro{"spring", 20, 181, true}
	fmt.Println(info01.string)
	fmt.Println(info01.int)
	fmt.Println(info01.float64)
	if info01.bool {
		fmt.Print("have the experience of work")
	} else {
		fmt.Println("do not have the experience of work")
	}
}
