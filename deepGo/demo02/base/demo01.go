// Package base @Author Spring
// @Date 2025/1/19 16:09:00
// @Desc
package main

import "fmt"

func main() {

	m1 := make(map[string]int, 10)
	m1["lucky"] = 18
	m1["jason"] = 24
	fmt.Println(m1)

	fmt.Println(m1["jason"])
	fmt.Println(m1["jason1"]) //map查询不存在的key也不会报错 返回了空值

	value, ok := m1["jason1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	for k, v := range m1 {
		fmt.Println(k, v) // k : lucky , v : 18
		// jason 24
		// lucky 18
	}

	delete(m1, "jason")
	fmt.Println(m1)

	delete(m1, "helloWorld")
	fmt.Println(m1)

}
