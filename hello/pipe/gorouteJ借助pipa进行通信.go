package main

import (
	"fmt"
	"time"
)

var pipa chan int

func add(a, b int) {
	var sum int
	sum = a + b
	time.Sleep(5 * time.Second) // 在这里sleep验证管道中没值时去取值会不会阻塞
	pipa <- sum                 // 将计算的结果放到管道中
}

func main() {
	pipa = make(chan int, 2)
	go add(10, 20) // 使用goroute调用add函数
	var value int
	value = <-pipa // 从管道中去除add函数放进去的值, 当从管道中取不出值时会阻塞在这里，这是主线程就会一直阻塞
	fmt.Println(value)
}
