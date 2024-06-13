//切片文法类似于没有长度的数组文法。
//
//这是一个数组文法：
//
//[3]bool{true, true, false}
//下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
//
//[]bool{true, true, false}
package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5}
	fmt.Println(a)

	b := []bool{true, false, true, true}
	fmt.Println(b)

	c := []struct {
		i int
		b bool
		t int
	}{
		{2, true, 1},
		{3, false, 2},
		{5, true, 3},
	}
	fmt.Println(c)
}
