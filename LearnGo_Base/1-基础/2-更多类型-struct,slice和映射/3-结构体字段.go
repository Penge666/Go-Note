//结构体字段使用点号来访问。

package main

import "fmt"

type A struct {
	x, y int
}

func main() {
	v := A{1, 2}
	v.y = 3
	fmt.Println(v)
}
