package main

import (
	"fmt"
	"src/easyGo/study_go03/function"
)

func main() {

	function.Hello()
	fmt.Println(function.Product(12, 24))
	function.GetMax()
	function.UserInfo("LeoChen", 20)
	function.Func_bmi("leochen", 1.81, 80)
	fmt.Println()
	fmt.Println("===================")
	num := 5
	function.MyFunc01(num)
	fmt.Println("out the func : num = ", num)
	fmt.Println("===================")
	function.MyFunc02(&num)
	fmt.Println("out the func : num = ", num)
	fmt.Println("===================")
	function.Sum01(1, 2, 3, 4, 5)
	function.Sum02(1, 2, 3, 4, 5)
	function.Sum02("chen", "leo", "dot")
	s_names := [3]string{"chen", "leo", "spring"}
	fmt.Println(function.ShowName01(s_names))
	fmt.Println(s_names)
	s_names01 := &[3]string{"chen", "leo", "spring"}
	fmt.Println(function.ShowName02(s_names01))
	fmt.Println(s_names01)
	s_names02 := []string{"chen", "leo", "spring"}
	fmt.Println(function.ShowName03(s_names02))
	fmt.Println(s_names02)
	minValue, maxValue := function.MinMax([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(minValue, maxValue)
	var nums = []int{5, 9, 7, 3, 2, 6}
	var rule string = "desc"
	if rule == "desc" {
		function.SortSlice(nums, function.DescSlice)
	} else {
		function.SortSlice(nums, function.AscSlice)
	}
}
