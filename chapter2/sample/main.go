package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	// 前面的"_"是为了让go语言对包做初始化操作,但是并使用包里的标识符,为了让程序的可读性更强,
	// go编译器不允许声明导入某个包却又不用, 下划线能够使编译器接受这类导入,并调用对应包内的所有代码文件里定义的init函数.
	// 对这个程序来说,这样做的目的是调用matchers包中的rss.go代码文件中的init函数,注册RSS匹配器, 以便后用.
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init在main之前调用
// 这个init函数会将标准库日志类的输出, 从默认的标准错误(stderr),设置为标准输出(stdout)
func init() {
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	// 使用特定的项做搜索
	/*这一行调用了search包里的Run函数, 这个函数包含程序的核心逻辑,需要传入一个字符串作为搜索项
	一旦Run函数退出,程序就会终止*/
	search.Run("president")
}

/*
每个可执行的go程序都有两个明显的特征:
1. 一个特征是第14行声明的名为main的函数,构建程序在构建可执行文件时,需要找到这个已经声明的main函数,把它作为程序的入口
2. 另一个是程序的第一行的包名main
*/

/*
可以看到main函数保存在名为main的包里, 如果main函数不在main包里, 构建工具就不会生成可执行文件
go语言的每一个代码文件都属于一个包, main.go也不例外, 包这个特性对于go来说很重要,
简单的理解:
一个包定义一组编译过得代码,包的名字类似于命名空间,可以用来间接访问包内声明的标识符,这个特征可以把不同的包中定义的同名标识符区分开来
*/
// TODO

