package tests

import (
	"io/ioutil"
	"testing"

	"imooc.com/crawler/zhenai/parser"
)

func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("./cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	parserResult := parser.ParseCityList(body)

	const resultSize = 470
	count := len(parserResult.Requests)

	if resultSize != count {
		t.Errorf("error num: %d. should be %d", count, resultSize)
	}
}
