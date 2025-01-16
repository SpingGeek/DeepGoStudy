package main

import "fmt"

var totalTime int

/*
计算上班的通勤时间
time_w: 上班的步行时间
tine_s: 上班的乘坐地铁的时间
return： 上班的通勤时间
*/
func getTotalTime(timeW, timeS int) int {
	totalTime := timeW + timeS
	return totalTime
}

func main() {
	var walkTime = 25
	var subwayTime = 12
	totalTime = getTotalTime(walkTime, subwayTime)
	fmt.Println(totalTime)
}
