package test

import (
	"testing"
	// "log"
	"spiderProxy/interval/dao"
)

func TestDownload(t *testing.T) {
	test := struct{
		url	string,
		length int,
	}{
		"https://www.aixdzs.com/down?id=159180&p=1"
		564657,
	}

	len := dao.DownloadBook(test.url)
	log.Fatal(len)
}