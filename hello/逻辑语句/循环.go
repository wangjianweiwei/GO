package main

import "fmt"

/*
func main() {
	for true {
		fmt.Println(1)
	}
}
*/

func main() {
	var b int = 5
	var a int

	//numbers := [6]int{1, 2, 3, 5}

	//for a := 0; a < 10; a++ {
	//	fmt.Println(a)
	//}

	for a < b {
		if a > 2 {
			break
		}
		a++
		fmt.Println(a)
	}

	//for i, x := range numbers {
	//	fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	//
	//}
}
