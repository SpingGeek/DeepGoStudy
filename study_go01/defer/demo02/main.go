package main

import "fmt"

func Demo(i int) {

	var arr [10]int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 10
}

func main() {
	Demo(10)
	fmt.Println("the program continue to execute")
}
