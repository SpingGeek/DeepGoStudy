package main

import "fmt"

func main() {

	var str01 = "from the young\n to study coding"
	fmt.Println(str01)

	var str02 = `新年快乐
老爸老妈`
	fmt.Println(str02)

	str1 := "Hello, "
	str2 := "world"
	str := str1 + str2
	fmt.Println(str)

	str3 := "Hello LeoChen"
	ls1 := len(str3)
	lrs1 := len([]rune(str3))
	fmt.Println(ls1, lrs1)
	fmt.Println("--------------")
	str4 := "Hello张三" // 5 + 6 one Chinese word account 3 byte
	ls2 := len(str4)
	lrs2 := len([]rune(str4))
	fmt.Println(ls2, lrs2)

}
