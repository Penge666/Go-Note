package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	allChan <- 10
	allChan <- "HelloGo"
	allChan <- Cat{"go", 6}

	<-allChan
	<-allChan
	tmp := <-allChan
	fmt.Println("tmp%T,%v", tmp, tmp)
	// 需要使用类型断言才能实现结构体取值
	res := tmp.(Cat)
	fmt.Println(res.Name)
}
