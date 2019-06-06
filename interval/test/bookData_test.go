package test

import (
	"testing"
	// "log"
	"spiderProxy/interval/dao"
)

func TestGetData(t *testing.T) {
	tests := []struct{
		BookNumber	string
		BookHref	string
		want		int
	}{
		{
			"/d/169/169208/",
			"p2.html",
			100,
		},
	}

	for _, test := range tests {
		bookData, _ := dao.GetBookData(test.BookHref, test.BookNumber)
		if len(bookData) != test.want {
			t.Errorf("(%q) = GetData(%v)", test.want, bookData)
		}
	}
}