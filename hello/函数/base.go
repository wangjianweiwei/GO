package main

import "fmt"

//func main() {
//	var a = 3
//	var b = 32
//	var ret int
//
//	ret = max(a, b)
//	fmt.Println(ret)
//}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Google", "Sath")
	fmt.Println(a, b)
}