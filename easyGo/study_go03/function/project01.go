package function

import (
	"fmt"
	"strconv"
)

func Hello() {
	fmt.Println("hello world")
}

func Product(m int, n int) int {
	var r int
	r = m * n
	return r
}

func GetMax() {
	num := []int{3, 5, 7, 4, 8, 9}
	maxValue := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > maxValue {
			maxValue = num[i]
		}
	}
	fmt.Println(maxValue)
}

func UserInfo(name string, age int) {
	fmt.Printf("name: %s age: %d", name, age)
}

func Func_bmi(person string, height float64, weight float64) {
	fmt.Printf("%s 的身高: %v 米\t 体重: %v 千克\n", person, height, weight)
	bmi := weight / (height * height)
	bmi_s := fmt.Sprintf("%.1f", bmi)
	bmi_f, _ := strconv.ParseFloat(bmi_s, 64)
	fmt.Printf("%s 的 BMI 指数为: %v\n", person, bmi_f)
	if bmi_f < 18.5 {
		fmt.Println("too thin")
	} else if bmi_f >= 18.5 && bmi_f < 24.9 {
		fmt.Printf("good")
	} else {
		fmt.Printf("notice too fat")
	}
}

func MyFunc01(num int) {
	num *= 2
	fmt.Println("in the func : num = ", num)
}

func MyFunc02(num *int) {
	*num *= 2
	fmt.Println("in the func : num = ", *num)
}

func Sum01(num ...int) {
	result := 0
	for _, value := range num {
		result += value
	}
	fmt.Println(result)
}

func Sum02(num ...interface{}) {
	for _, value := range num {
		fmt.Println(value)
	}
}

func ShowName01(names [3]string) [3]string {
	names[len(names)-1] = "Jerry"
	return names
}

func ShowName02(names *[3]string) *[3]string {
	names[len(names)-1] = "Jerry"
	return names
}

func ShowName03(names []string) []string {
	names[len(names)-1] = "Jerry"
	return names
}

func MinMax(num []int) (int, int) {
	minValue := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < minValue {
			minValue = num[i]
		}
	}
	maxValue := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > maxValue {
			maxValue = num[i]
		}
	}
	return minValue, maxValue

}

func AscSlice(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[j] < num[i] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
}

func DescSlice(num []int) {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[j] > num[i] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
}

func SortSlice(num []int, sort func(num []int)) {
	sort(num)
	fmt.Println(num)
}
