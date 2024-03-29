package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds读取反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 当函数返回时关闭文件
	defer file.Close()

	// 将文件解码到一个切片中,这切片的每一项是一个指向Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 这个函数不需要检查错误,调用者会做这件事
	return feeds, err
}
