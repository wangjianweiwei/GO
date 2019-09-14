package main

import (
	"dev/day1/package_example/calc"
	"fmt"
)

func main() {
	value2 := calc.Add(100, 200) // 大写字母开头的表示公共的，才可以被导出使用
	value1 := calc.Sub(200, 100)

	fmt.Println("sum:", value2)
	fmt.Println("sub:", value1)
}

// 在项目目录下执行打包 /Users/mac/wangjianwei/project/GO
// go build -o bin/calc.bin dev/day1/package_example/main
