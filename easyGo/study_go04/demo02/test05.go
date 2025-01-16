package main

import (
	"fmt"
	"strconv"
)

type profile1 struct {
	name     string
	age      int
	interest []string
	evaluate string
}

func (p *profile1) get_info(subject string, score int) string {
	info := " name : " + p.name + "\n age: " + strconv.Itoa(p.age)
	info += "\n 考试科目: " + subject + "\n 考试分数: " + strconv.Itoa(score)
	return info
}

func main() {
	p := &profile1{name: "张三", age: 20}
	result := p.get_info("综合知识", 76)
	fmt.Println(result)
}
