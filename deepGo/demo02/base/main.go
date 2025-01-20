// @Author Spring
// @Date 2025/1/19 16:31:00
// @Desc
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// init the Random the number of seed
	rand.Seed(time.Now().UnixNano())

	// declare and init the map
	var scoreMap = make(map[string]int, 100)

	// fill the data
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value            //这里的赋值用= 而不是:=
	}
	//总结： =是赋值 :=是声明并赋值

	// take the map of Key to the keys slice
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的keys遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
