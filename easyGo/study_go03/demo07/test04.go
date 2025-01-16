package main

import "fmt"

func main() {

	var subject_score map[string]int
	subject_score = map[string]int{"Chinese": 97, "math": 95}

	subject_score["English"] = 92
	delete(subject_score, "math")
	for k, v := range subject_score {
		fmt.Println(k, v)
	}
}
