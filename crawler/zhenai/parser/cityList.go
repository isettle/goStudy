package parser

import (
	"regexp"

	"imooc.com/crawler/engine"
)

const cityListPat = `<a[^href]*href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityListPat)
	regText := reg.FindAllSubmatch(contents, -1)

	// 响应返回
	result := engine.ParserResult{}
	for _, match := range regText {
		// 城市名
		result.Items = append(result.Items, "city:"+string(match[2]))
		// 城市链接
		result.Requests = append(result.Requests, engine.Request{Url: string(match[1]), ParserFun: ParseCity})
	}

	return result
}
