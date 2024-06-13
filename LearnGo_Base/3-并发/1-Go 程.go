// 来看个例子：需求:要求统计1-9000000000的数字中，哪些是素数?
// 使用goroutine
// Go主线程(有程序员直接称为线程/也可以理解成进程):一个Go线程上，可以起多个协程，你可以这样理解，协程是轻量级的线程[编译器做优化]。

//请编写一个程序，完成如下功能:
//1)在主线程(可以理解成进程)中，开启一个goroutine，该协程每隔Ⅰ秒输出"hello,world"
//2)在主线程中也每隔一秒输出"hello,golang"，输出10次后，退出程序
//3)要求主线程和goroutine同时执行.
//4)画出主线程和协程执行流程图

package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	go test()
	for i := 1; i <= 10; i++ {
		fmt.Println("main() ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
