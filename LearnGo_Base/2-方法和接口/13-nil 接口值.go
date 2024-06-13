//nil 接口值既不保存值也不保存具体类型。
//
//为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。

package main

import "fmt"

type UI interface {
	M()
}

func main() {
	var i UI
	Describe(i)
	i.M()
}

func Describe(i UI) {
	fmt.Printf("(%v, %T)\n", i, i)
}
