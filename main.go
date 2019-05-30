package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"net/url"
)

func searchBook(bookName string) (resHtml string, err error) {
	url := "https://www.aixdzs.com/bsearch?q=" + url.QueryEscape(bookName)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer res.Body.Close()
	Body, err := ioutil.ReadAll(res.Body)
	return Body, err
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	bookName := "大道朝天"
	Body, err := searchBook(bookName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("body html /n s%", Body)
}