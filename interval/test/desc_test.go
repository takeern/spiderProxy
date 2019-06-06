package test

import (
	"testing"
	// "log"
	"spiderProxy/interval/dao"
)

func TestGetDesc(t *testing.T) {
	tests := []struct{
		want		string
		bookName	string
	}{
		{
			"/d/169/169208/",
			"大道朝天",
		},
	}

	for _, test := range tests {
		params, _ := dao.GetBookDesc(test.bookName)
		if string(params[0][1]) != test.want {
			t.Errorf("(%q) = desc(%v)", test.want, params[0][1])
		}
	}
}
