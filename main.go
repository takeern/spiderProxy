package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func searchBook(bookName string) (resHtml string, err error) {
	url := "https://www.ixdzs.com/bsearch?q=" + bookName
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer res.Body.Close()
	Body, err := ioutil.ReadAll(res.Body)
	return Body, err
}

func main() {
	bookName := "大道朝天"
	Body, err := searchBook(bookName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("body html /n s%", Body)
}