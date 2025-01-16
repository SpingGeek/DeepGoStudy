package main

import "fmt"

func main() {

	var subject_score map[string]int
	subject_score = map[string]int{"Chinese": 95, "math": 97}
	fmt.Println(subject_score["Chinese"])
	fmt.Println(subject_score["math"])
	fmt.Println(subject_score["maths"])
	subject_score["English"] = 92
	fmt.Println(subject_score["English"])
	fmt.Println(subject_score["English01"])
	// if not to match the value , it can print the zero
}
