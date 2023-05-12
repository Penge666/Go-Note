//现在我们要把 Abs 和 Scale 方法重写为函数。
//
//同样，我们先试着移除掉第 16 的 *。你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？
package main

import "fmt"

type SSTT struct {
	x, y float64
}

func (sstt SSTT) looklook() {
	fmt.Println(sstt.x, " ", sstt.y)
}
func scale3(sstt *SSTT, f float64) {
	sstt.x = sstt.x * f
	sstt.y = sstt.y * f
}
func main() {
	sstt := SSTT{1, 2}
	//sstt.scale3(10)
	scale3(&sstt, 10)
	sstt.looklook()
}
