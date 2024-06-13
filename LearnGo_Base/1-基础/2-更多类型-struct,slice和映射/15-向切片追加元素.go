//为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。内建函数的文档对此函数有详细的介绍。
//
//func append(s []T, vs ...T) []T
//append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。
//
//append 的结果是一个包含原切片所有元素加上新添加元素的切片。
//
//当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
//
//（要了解关于切片的更多内容，请阅读文章 Go 切片：用法和本质。）

package main

import "fmt"

func main() {
	var s []int
	printSlice3(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice3(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice3(s)

	// 可以一次性添加多个元素
	// 要增加切片的容量必须创建一个新的、更大容量的切片，然后将原有切片的内容复制到新的切片。 整个技术是一些支持动态数组语言的常见实现。下面的例子将切片 s 容量翻倍，先创建一个2倍 容量的新切片 t ，复制 s 的元素到 t ，然后将 t 赋值给 s ：

	s = append(s, 2, 3, 4)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
