package main

import (
	"fmt"
	"strconv"
)

type profile struct {
	name string
	age  int
}

func (p profile) get_info(subject string, score int) string {
	info := " name : " + p.name + "\n age: " + strconv.Itoa(p.age)
	info += "\n 考试科目: " + subject + "\n 考试分数: " + strconv.Itoa(score)
	return info
}

func main() {
	p := profile{name: "张三", age: 20}
	result := p.get_info("综合知识", 76)
	fmt.Println(result)
}
