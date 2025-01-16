package main

import "fmt"

func main() {

	var actor struct {
		name        string
		work        []string
		achievement string
	}

	actor.name = "spring"
	actor.work = []string{"code dev", "code top"}
	actor.achievement = "good coder and developer"

	fmt.Println(actor.name)
	fmt.Println(actor.work)
	fmt.Println(actor.achievement)

	actor01 := struct {
		name        string
		work        []string
		achievement string
	}{
		name:        "leochen",
		work:        []string{"coder"},
		achievement: "best",
	}

	fmt.Println(actor01.name)
	fmt.Println(actor01.work)
	fmt.Println(actor01.achievement)
}
