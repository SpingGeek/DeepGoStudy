package main

import "fmt"

func main() {
	var test1 map[string]string
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1)

	fmt.Println("------------------")

	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2)

	fmt.Println("------------------")

	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(test3)
}
