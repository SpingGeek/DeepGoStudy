// @Author Gopher
// @Date 2025/1/29 14:23:00
// @Desc
package main

import "fmt"

func checkPalindrome(str string) {
	runeStr := []rune(str)
	length := len(runeStr)
	for i := 0; i < length/2; i++ {
		if runeStr[i] != runeStr[length-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

func main() {
	var str string
	fmt.Println("请输入一个字符串")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	checkPalindrome(str)
}
