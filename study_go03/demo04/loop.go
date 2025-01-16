package main

import (
	"fmt"
	"strings"
)

func strEntry(r rune) rune {
	return r + 1
}

func main() {

	str := "hello"
	fmt.Println(str)
	newStr := strings.Map(strEntry, str)
	fmt.Println(newStr)

	str0 := "张三hello"
	str1 := str0[0:2]
	srn := []rune(str1)
	str2 := srn[0:2]
	fmt.Println(string(str1))
	fmt.Println(string(str2))

	string01 := "Hello World!"
	strArr := strings.Fields(string01)
	fmt.Println(strArr)
}
