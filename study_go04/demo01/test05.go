package main

import (
	"encoding/json"
	"fmt"
)

type spots struct {
	Spot []struct {
		Name   string
		Level  string
		Ticket int
	}
}

type city struct {
	name    string
	code    string
	tourism spots
}

func main() {

	var s spots
	c := city{name: "beijing", code: "0431", tourism: s}
	j := `{"spot":[
	     {"name":"go","level":"P8","ticket":"100"}]}`
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		return
	}
	fmt.Println(c.name)
	fmt.Println(c.code)
}
