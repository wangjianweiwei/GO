package search

// 实现默认的匹配器
type defaultMathcher struct{}

// init函数将默认匹配器注册到程序中
func init() {
	var matcher defaultMathcher

	Regiser("default", matcher)
}

// Search实现默认的匹配器行为
func (m defaultMathcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
