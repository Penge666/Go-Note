/*
channle介绍
	1) channle本质就是一个数据结构-队列【示意图】
	2)数据是先进先出【FIFO : first in first out 】
	3)线程安全，多 goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
	4) channel有类型的，一个string 的channel只能存放string 类型数据。
定义
	var 变量名 chan 数据类型
注意事项1
	channel是引用类型
	channel必须初始化才能写入数据，即 make后才能使用管道是有类型的，
	intChan只能写入整数int
注意事项2
	1) channel中只能存放指定的数据类型
	2) channle的数据放满后，就不能再放入了
	3) 如果从channel取出数据后，可以继续放入
	4) 在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock

*/
package main

import "fmt"

func main() {
	// 定义管道变量
	var intChan chan int
	intChan = make(chan int, 3)

	// 查看intChan
	fmt.Println("intChan的值%v以及intChan本身的地址%p", intChan, &intChan)

	//向管道写入数据
	intChan <- 10
	num := 21
	intChan <- num
	intChan <- 50
	// 查看管道的len以及cap
	fmt.Println("Channel len=%v cap=%v", len(intChan), cap(intChan))

	// 从管道中读取数据

	num2 := <-intChan
	fmt.Println(num2)
	fmt.Println("Channel len=%v cap=%v", len(intChan), cap(intChan))
}
