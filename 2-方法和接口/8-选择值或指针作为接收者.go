//使用指针接收者的原因有二：
//
//首先，方法能够修改其接收者指向的值。
//
//其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
//
//在本例中，Scale 和 Abs 接收者的类型为 *Vertex，即便 Abs 并不需要修改其接收者。
//
//通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。（我们会在接下来几页中明白为什么。）
//

package main

import "fmt"

type TT struct {
	x, y float64
}

func (tt *TT) scale_func(f float64) {
	tt.x = tt.x * f
	tt.y = tt.y * f
}
func (tt TT) Look() {
	fmt.Println(tt.x, " ", tt.y)
}
func main() {
	tt := &TT{3, 4}
	tt.scale_func(10)
	tt.Look()
}
