//可以将下标或值赋予 _ 来忽略它。
//
//for i, _ := range pow
//for _, value := range pow
//若你只需要索引，忽略第二个变量即可。
//
//for i := range pow
package main

import "fmt"

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = 1 << uint(i)
	}
	for _, val := range a {
		fmt.Println(val)
	}
}
