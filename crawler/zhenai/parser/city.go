package parser

import (
	"regexp"

	"imooc.com/crawler/engine"
)

var cityPat = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]*)</a></th>`
var gender = `<tr><td[^>]*><span[^>]*>性别：</span>([^<]*)</td>[^tr]*tr>`

func ParseCity(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(cityPat)
	regText := reg.FindAllSubmatch(contents, -1)

	// 响应返回
	result := engine.ParserResult{}
	for _, match := range regText {
		// 用户
		result.Items = append(result.Items, "user:"+string(match[2]))
		// 用户链接
		result.Requests = append(result.Requests, engine.Request{Url: string(match[1]), ParserFun: ParseProfile})
	}

	return result
}
