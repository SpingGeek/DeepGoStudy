package main

import "fmt"

func main() {

	var subject_score map[string]int
	subject_score = make(map[string]int, 2)
	subject_score["Chinese"] = 98
	subject_score["math"] = 100
	fmt.Println(subject_score["Chinese"])
	fmt.Println(subject_score["math"])
	subject_score["English"] = 97
	fmt.Println(subject_score["English"])
}
