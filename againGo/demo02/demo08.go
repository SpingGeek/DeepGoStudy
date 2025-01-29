// @Author Gopher
// @Date 2025/1/29 11:03:00
// @Desc
package main

import "fmt"

func f3() {
	m := make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

func main() {
	f3()
}
