//即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
//
//在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
//
//注意: 保存了 nil 具体值的接口其自身并不为 nil。

package main

import "fmt"

type Y interface {
	M()
}
type X struct {
	s string
}

func (x *X) M() {
	fmt.Println("x->M()")
	if x == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("x->M() !nil")
	fmt.Println(x.s)
}
func main() {
	var y Y
	var x *X
	y = x
	y.M()
}
