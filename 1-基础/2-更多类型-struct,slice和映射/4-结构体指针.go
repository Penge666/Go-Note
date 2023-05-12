//结构体字段可以通过结构体指针来访问。
//
//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。

package main

import "fmt"

type B struct {
	x, y int
}

func main() {
	v := B{1, 2}
	p := v
	// p:=&v //ok
	p.x = 2
	fmt.Println(p)
}
