/*
请完成goroutine和channel协同工作的案例,具体要求:
	1)开启一个writeData协程,向管道intChan中写入50个整数.
	2)开启一个readData协程，从管道intChan中读取writeData写入的数据。
	3)注意: writeData和readDate操作的是同一个管道
	4)主线程需要等待writeData和readDate协程都完成工作才能退出【管道】

*/
package main

import "fmt"

var (
	intChan  = make(chan int, 51)
	exitChan = make(chan bool, 1)
)

// writeData
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("writeData", i)
	}
	close(intChan)
}

// readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
