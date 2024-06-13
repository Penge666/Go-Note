//使用select可以解决从管道取数据的阻塞问题
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个int管道
	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intchan <- i
	}
	// 定义一个string管道
	stringchan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	//问题，在实际开发中，可能我们不好确定什么关闭该管道.
	//
	//可以使用select方式可以解决

	for {
		select {
		//注意:这里，如果intChan一直没有关闭，不会一直阻塞而deadlock，会自动到下一个case匹配
		case v := <-intchan:
			fmt.Println("从intChan读取的数据%d", v)
			time.Sleep(time.Second)
		case v := <-stringchan:
			fmt.Println("从stringChan读取的数据%s", v)
			time.Sleep(time.Second)
		default:
			fmt.Println("GG")
			time.Sleep(time.Second)
		}
	}

}
