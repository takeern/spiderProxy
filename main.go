package main

import (
	"fmt"

	"spiderProxy/interval/dao"
)

// type Bookdesc struct {
// 	BookName	string
// 	BookNumber	string
// 	BookState	string
// 	BookIntro	string
// }

// func searchBook(bookName string, tryTimes int) (resHtml string, err error) {
// 	if (tryTimes > 3) {
// 		return "", errors.New("try many times")
// 	}
// 	url := "https://www.aixdzs.com/bsearch?q=" + url.QueryEscape(bookName)
// 	res, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println("http err", err.Error())
// 		tryTimes ++
// 		return searchBook(bookName, tryTimes)
// 	}

// 	defer res.Body.Close()
// 	Body, err := ioutil.ReadAll(res.Body)

// 	return string(Body), err
// }

// func getBookDesc(html string) ([]Bookdesc, error) {
// 	var resp []Bookdesc
// 	if (len(html) == 0) {
// 		return resp, errors.New("string length = 0")
// 	}

// 	re := regexp.MustCompile(`b_name"><a href="(.{1,}?)"[\s\S]{1,}?"_blank">(.{1,}?)<[\s\S]{1,}?p">(.{1,}?)<[\s\S]{1,}?b_intro">([\s\S]{1,}?)<`)
// 	params := re.FindAllSubmatch([]byte(html), -1)
// 	for _, param := range params {
// 		resp = append(resp, Bookdesc{
// 			BookName: string(param[2]),
// 			BookNumber: string(param[1]),
// 			BookIntro: string(param[4]),
// 			BookState: string(param[3]),
// 		})
// 	}
// 	return resp, nil
// }

func main() {

	bookName := "大"
	res, err := dao.GetBookDesc(bookName)
	if err != nil {
		fmt.Println(err)
	}
	// res, err1 := getBookDesc(html)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }
	// fmt.Println("%+v\n", res.BooksDesc)
	for index, item := range res.BooksDesc {
		fmt.Printf("%+v\n", index)
		fmt.Printf("%+v\n", item)
	}
}