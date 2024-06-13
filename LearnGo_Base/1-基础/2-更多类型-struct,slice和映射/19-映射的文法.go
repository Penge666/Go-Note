//映射的文法与结构体相似，不过必须有键名。

package main

import "fmt"

type E struct {
	x, y int
}

func main() {
	var mp = map[string]E{
		"A": {1, 2},
		"B": {2, 3},
	}
	fmt.Println(mp)
}
