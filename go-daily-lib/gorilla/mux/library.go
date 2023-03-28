package mux

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// https://darjun.github.io/2021/07/19/godailylib/gorilla/mux/
type Book struct {
	ISBN string `json:"isbn"`
	Name string `json:"name"`
	Authors []string `json:"authors"`
	Press string `json:"press"`
	PublishedAt string `json:"published_at"`
}

var (
	mapBooks map[string]*Book
	slcBooks []*Book
)

// 从文件中加载数据
func init() {
	mapBooks = make(map[string]*Book)
	slcBooks = make([]*Book, 0, 1)

	data, err := ioutil.ReadFile("../data/books.json")
	if err != nil {
		log.Fatalf("failed to read book.json:%v", err)
	}

	err = json.Unmarshal(data, &slcBooks)
	if err != nil {
		log.Fatalf("failed to unmarshal books:%v", err)
	}

	for _, book := range slcBooks {
		mapBooks[book.ISBN] = book
	}
}

// 返回整个列表