// @Author Gopher
// @Date 2025/1/19 16:51:00
// @Desc
package main

import "fmt"

func main() {
	var s = make([]map[string]int, 1, 10)
	fmt.Println(s)

	s[0] = make(map[string]int, 1)
	s[0]["lucky"] = 18
	fmt.Printf("s类型:%T s的值：%v\n", s, s)

	// 值为切片类型的map
	var m = make(map[string][]int, 1)
	m["北京"] = []int{1, 2, 3} //声明并初始化了
	fmt.Printf("m类型：%T m的值：%v", m, m)

	// 值为map类型的map
	var m1 = make(map[string]map[string]int, 1)
	m1["北京"] = make(map[string]int, 1)
	m1["北京"]["lucky"] = 18
	fmt.Printf("m1类型:%T m1的值:%v\n", m1, m1)
}
