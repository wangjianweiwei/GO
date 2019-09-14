package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go testPrint(i) // 启动一个线程
	}
	time.Sleep(time.Second) // 让主线程等待子线程
}

func testPrint(num int) {
	fmt.Println(num)
}
