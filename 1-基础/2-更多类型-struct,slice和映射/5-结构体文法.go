//结构体文法通过直接列出字段的值来新分配一个结构体。
//
//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
//
//特殊的前缀 & 返回一个指向结构体的指针。

package main

import "fmt"

type C struct {
	X, Y int
}

var (
	v1 = C{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = C{X: 1}  // Y:0 被隐式地赋予
	v3 = C{}      // X:0 Y:0
	p  = &C{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
