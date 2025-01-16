package main

/**
slice
Go 语言切片是对数组的抽象
Go 数组的长度不可发生改变，在特定的场景下这样的集合就不太适用，Go中提供了一种
灵活的，功能强悍的内置类型切片("动态数组")，与数组相比较，切片的长度是不固定的
可以再去进行元素的追加，在追加时可能使得切片的容量变大
*/

// 定义切片
// var identifier []type  切片是不需要说明其长度的

// 使用 make 来去创建切片
// var slice1 []type = make([]type, len)

// 简写
// slice1 := make([]type, len)

// 也可以去指定其容量，其中 capacity 为可选参数
// make([]T, length, capacity)
