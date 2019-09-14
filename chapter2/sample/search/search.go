package search

import (
	"log"
	"sync"
)

/*
可以看到每一个文件都是以package关键字开头,随后跟着包名.
文件夹search下的每一个代码文件都是用search作为包名.
*/

/*
与导入第三方的包不同, 从标准库中导入代码时,只需要给出要导入的包名,编译器查找包的时候总是回到GOROOT和GOPATH环境变量应用的位置去查找
*/

// 注册用于搜索的匹配器映射
var matchers = make(map[string]Matchers)
/*
matchers 变量没有声明在任何函数作用域中, 所以会被当成包级变量,这个变量使用关键字var声明
而且声明为Matchers类型的映射(map), 这个映射以string类型值作为键, Matchers类型值作为值,
Matchers类型在代码文件matcher.go中声明.

这里有一个地方要注意, 变量名matchers是以小写字母开头的,
在go语言中标识符要么从包里公开, 要么不从包里公开,当代码导入一个包时,程序可以直接访问这个保重任意公开的标识符,
这些标识符都是以大写字母开头, 以小写字母开头的标识符都是不公开的, 不能被其他包中的代码直接访问, 但是,
其他包可以间接访问不公开的标识符,
例如: 一个函数可以返回一个未公开类型的值, 那么这个函数的任何调用者, 哪怕调用者不是在这个包里声明的,都可以访问这个值
这行变量的声明还能使用复制运算符和特殊的内置函数make初始化了变量
*/

// Run执行搜索的逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个无缓冲的通道, 接受匹配后的结果
	results := make(chan *Result)
	// 构造一个waitGroup, 以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待等待处理每一个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每一个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// 启动一个goroutine来执行搜索
		go func(matcher Matcherm, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}
	// 启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		waitGroup.Wait()
		close(results)
	}()
	// 启动函数,显示返回的结果,并且在最后一个结果显示完后返回
	Display(results)
}
