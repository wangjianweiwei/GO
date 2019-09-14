package main

import (
	"fmt"
)

func main() {
	pipe := make(chan int, 3) // 规定make的是一个管道，存放的数据类型是int，管道的长度为3，类似队列(FIFO)
	// 往管道中放数据
	pipe <- 1
	pipe <- 2
	pipe <- 3

	var value int

	value = <-pipe // 从管道中获取数据并赋值给value

	fmt.Println(value)
	fmt.Println(len(pipe))
}
