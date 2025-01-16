package main

import "fmt"

type exam struct {
	subject string
	score   int
}

type info struct {
	name   string
	age    int
	record exam
}

func main() {

	e := exam{subject: "picture", score: 80}
	i := info{name: "spring", age: 23, record: e}
	fmt.Println(i.name)
	fmt.Println(i.age)
	fmt.Println(i.record.subject)
	fmt.Println(i.record.score)

}
