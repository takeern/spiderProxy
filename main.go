package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"net/url"
)

type SearchResp struct {
	BookName	string
	BookNumber	string
	BookState	bool
	BookIntro	string
}

func searchBook(bookName string) (resHtml string, err error) {
	url := "https://www.aixdzs.com/bsearch?q=" + url.QueryEscape(bookName)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	defer res.Body.Close()
	Body, err := ioutil.ReadAll(res.Body)
	return string(Body), err
}

func praseHtml(html string) SearchResp {
	s := SearchResp{
		BookName: "大道",
		BookNumber: "ddsds",
		BookState: true,
		BookIntro: "dsdsds",
	}
	return s
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	bookName := "大道朝天"
	_, err := searchBook(bookName)
	if err != nil {
		fmt.Println(err)
	}
	s := praseHtml("dsdss")
	fmt.Println("body html /n s%", s)
}