//使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel 写数据了，但是仍然可以从该channel读取数据

package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// error:不能·写入
	//panic: send on closed channel
	// intChan <- 300
	num := <-intChan
	fmt.Println(num)
}
