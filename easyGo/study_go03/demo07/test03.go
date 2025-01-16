package main

import "fmt"

func main() {

	var subject_score map[string]int
	subject_score = map[string]int{"Chinese": 97, "math": 95}
	var value, isOK = subject_score["math"]
	fmt.Println(value, isOK)
}
