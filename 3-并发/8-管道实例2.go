/*
任务：要求统计1-200000 的数字中，哪些是素数﹖这个问题在本章开篇就提出了，现在我们有goroutine和 channel 的知识后，就可以完成了[测试数据: 80000]
思路：使用并发/并行的方式，将统计素数的任务分配给多个(4个)goroutine去完成，完成任务时间短。

*/

package main

import (
	"fmt"
	"time"
)

var (
	inputChan = make(chan int, 2000)
	primeChan = make(chan int, 5000)
	exit_Chan = make(chan bool, 4)
)

//putchan 用于存数
func putchan(inputChan chan int) {
	for i := 1; i <= 8000; i++ {
		inputChan <- i
	}
	close(inputChan)
}
func primeNum(inputChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-inputChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	exit_Chan <- true
}
func main() {
	go putchan(inputChan)
	for i := 0; i < 4; i++ {
		go primeNum(inputChan, primeChan, exit_Chan)
	}

	for i := 0; i < 4; i++ {
		<-exit_Chan
	}
	close(primeChan)

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
