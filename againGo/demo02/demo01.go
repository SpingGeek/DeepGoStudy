// Package demo02 @Author Gopher
// @Date 2025/1/29 10:04:00
// @Desc
package main

func main() {
	const PI float64 = 3.14
	const PI2 = 3.1415 // 可以省略类型

	const (
		width  = 100
		height = 200
	)

	const i, j = 1, 2 // 多重赋值
	const a, b, c = 1, 2, "foo"

	println(PI, PI2, width, height, i, j, a, b, c)
}
