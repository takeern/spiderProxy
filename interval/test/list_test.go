package test

import (
	"testing"
	// "log"
	"spiderProxy/interval/dao"
)

func TestGetList(t *testing.T) {
	tests := []struct{
		BookNumber	string
		want		int
	}{
		{
			"/d/169/169208/",
			2,
		},
	}

	for _, test := range tests {
		params, _ := dao.GetBookList(test.BookNumber)
		if len(params) != test.want {
			t.Errorf("(%q) = GetList(%v)", test.want, len(params))
		}
	}
}