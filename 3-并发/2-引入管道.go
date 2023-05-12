//需求:现在要计算1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。最后显示出来。要求使用goroutine完成

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 21)
	lock  sync.Mutex
)

//Write函数就是计算n!，让将这个结果放入到myMap

func Write(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go Write(i)
	}
	// 休眠10s
	time.Sleep(time.Second * 10)
	for i, v := range myMap {
		fmt.Println("map[%d]=%d\n", i, v)
	}
}
