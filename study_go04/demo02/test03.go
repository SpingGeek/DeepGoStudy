package main

import "fmt"

type Job struct {
	position string
	year     int
}

type Info struct {
	name   string
	age    int
	record Job
}

func NewJob(position string, year int) *Job {
	return &Job{
		position: position,
		year:     year,
	}
}

func NewInfo(name string, age int, position string, year int) *Info {
	record := *NewJob(position, year)
	return &Info{
		name:   name,
		age:    age,
		record: record,
	}
}

func main() {

	p := NewInfo("张三", 32, "技术总监", 6)
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.record.position)
	fmt.Println(p.record.year)
}
