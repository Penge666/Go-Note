/*
channel支持for--range的方式进行遍历，请注意两个细节
	1)在遍历时，如果channel没有关闭，则回出现deadlock 的错误
	2)在遍历时，如果channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
*/
package main

import "fmt"

func main() {
	intchan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intchan <- i * 2
	}
	//1)在遍历时，如果channel没有关闭，则回出现deadlock 的错误
	close(intchan)
	for v := range intchan {
		fmt.Println("v=", v)
	}

}
