package tests

import (
	"fmt"
	"io/ioutil"
	"testing"

	"imooc.com/crawler/zhenai/parser"
)

func TestProfile(t *testing.T) {
	body, err := ioutil.ReadFile("./profile_test_data.html")
	// body, err := fetcher.Fetch("http://album.zhenai.com/u/1595163909")
	if err != nil {
		t.Errorf("获取文件错误: %v", err)
	}

	parserResult := parser.ParseProfile(body)
	fmt.Printf("%+v \n", parserResult.Items)
}
