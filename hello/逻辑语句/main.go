package main

import "fmt"

//func main() {
//	var num = 10
//
//	if num < 3 {
//		fmt.Println("da yu")
//	}else {
//		fmt.Println("xiao yu")
//	}
//}

// if 嵌套

/*
func main() {
	var num int = 0
	if num > 1 {
		fmt.Println("大于1")
		if num > 2 {
			fmt.Println("大于2")
		} else {
			fmt.Println("小于2")
		}
	}else {
		fmt.Println("小于1")
	}
}
*/

/*
// switch
func main() {
	var grade = "B"
	var marks = 70

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60, 50:
		grade = "C"
	default:
		grade = "D"
	}
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("笨蛋")

	}
}
*/

// type switch
/*
func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T", i)
	case int:
		fmt.Printf("x 的类型：%T", i)
	case float64:
		fmt.Printf("x 的类型 %T", i)
	case func(int) float64:
		fmt.Printf("x 的类型：%T", i)
	case bool, string:
		fmt.Printf("x 是 bool或string 类型")
	default:
		fmt.Printf("未知类型")

	}
}
*/
/*
func main() {

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case true:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
*/

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received", i1, "from c1")
	case c2 <- i2:
		fmt.Println("sent", i2, "to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received", i3, "from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
}
