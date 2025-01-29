// @Author Gopher
// @Date 2025/1/29 11:10:00
// @Desc
package main

import "fmt"

func f4() {
MY_LABEL:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break MY_LABEL
		}
		fmt.Printf("%v\n", i)
	}
	fmt.Println("end...")
}
func main() {
	f4()
}
