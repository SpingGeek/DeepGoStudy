// Package main @Author Spring
// @Date 2025/1/16 09:23:00
// @Desc
package main

import "unsafe"

// slice 底层
type slice struct {

	// 底层数组指针 或者说是指向一块连续内存的空间的起点
	array unsafe.Pointer

	// 长度
	len int

	// 容量
	cap int
}

type _type struct {
	size  uintptr
	align uintptr
	kind  uint8
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	less  func(unsafe.Pointer, unsafe.Pointer) bool
	hash  func(unsafe.Pointer) uintptr
}

// growSlice 内存对齐
func growSlice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice {

	// oldPtr: 旧切片的底层数组指针
	// newLen: 新切片的预估长度 （oldLen + num）
	// oldCap: 旧切片的容量
	// num: append 的元素的个数
	// et: 切片的元素类型

	// 得到原来切片的长度
	oldLen := newLen - num

	newCap := oldCap
	doubleCap := newCap + newCap

	// 如果新的切片的长度 大于 旧的切片的容量的两倍， 那么就需要去扩容成为新的切片的长度
	if newLen > doubleCap {
		newCap = newLen
	} else {
		// 扩容策略分流 阈值 256
		const threshold = 256

		// 如果旧的切片的容量是小于 256 ， 那么新的切片的容量就是旧的切片的容量的两倍
		if oldCap < threshold {
			newCap = doubleCap
		} else {
			// 扩容 1.25
			for 0 < newCap && newCap < newLen {
				newCap += (newCap + 3*threshold) / 4
				oldLen = oldLen
			}
		}
	}
	return slice{
		array: unsafe.Pointer(oldPtr),
		len:   oldLen,
		cap:   newCap,
	}
}
