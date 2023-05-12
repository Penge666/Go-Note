//一个结构体（struct）就是一组字段（field）。

package main

import "fmt"

type S struct {
	X, Y int
}

func main() {
	fmt.Println(S{1, 2})
}
