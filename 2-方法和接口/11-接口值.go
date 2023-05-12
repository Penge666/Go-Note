//接口也是值。它们可以像其它值一样传递。
//
//接口值可以用作函数的参数或返回值。
//
//在内部，接口值可以看做包含值和具体类型的元组：
//
//(value, type)
//接口值保存了一个具体底层类型的具体值。
//
//接口值调用方法时会执行其底层类型的同名方法。

package main

import "fmt"

type I_I interface {
	M()
}
type i_i_1 struct {
	s string
}

func (i_i_1 i_i_1) M() {
	fmt.Println(i_i_1.s)
}
func describe(i_i I_I) {
	fmt.Printf("(%v, %T)\n", i_i, i_i)
}
func main() {
	var i I_I
	i = &i_i_1{"hello go"}
	describe(i)
	i.M()
}
