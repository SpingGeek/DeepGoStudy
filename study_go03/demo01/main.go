package main

import "fmt"

func myFunc() (string, string, int) {
	name := "zhangSan"
	sex := "nan"
	age := 20
	return name, sex, age
}

func main() {
	name, _, _ := myFunc()
	fmt.Println(name)
}
