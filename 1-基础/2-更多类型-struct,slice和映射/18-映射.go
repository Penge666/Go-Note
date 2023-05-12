//映射将键映射到值。
//
//映射的零值为 nil 。nil 映射既没有键，也不能添加键。
//
//make 函数会返回给定类型的映射，并将其初始化备用。
//

package main

import "fmt"

type D struct {
	x, y int
}

func main() {
	var mp = make(map[string]D)
	mp["HelloGo"] = D{1, 2}
	fmt.Println(mp["HelloGo"])
}
