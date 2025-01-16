package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 5
	var f float64 = 7.11
	var str1 string = "false"
	var str2 string = "100"
	// divide to make the int and float to the string
	fmt.Println("转化后的字符串的值: " + strconv.Itoa(i) + " " + strconv.FormatFloat(f, 'g', 5, 32))
	num, err := strconv.Atoi(str2)
	fmt.Println(num + 50)
	fmt.Println(err)
	fmt.Println(strconv.ParseBool(str1))
}
