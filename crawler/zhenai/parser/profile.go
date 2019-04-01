package parser

import (
	"regexp"
	"strconv"

	"imooc.com/crawler/engine"
	"imooc.com/crawler/model"
)

var avatarReg = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\(([^\)]*)\);"[^<]*></div>`)
var nickname = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]*)</span>`)
var introduction = regexp.MustCompile(`<div class="m-content-box m-des"[^>]*><span[^>]*>([^<]*)</span><!----></div>`)
var otherInfo = regexp.MustCompile(`<div class="des f-cl"[^>]*>([\p{Han}]*) \| (\d*)岁 \| ([\p{Han}]*) \| ([\p{Han}]*) \| (\d*)cm \| ([\d-]*)元</div>`)

func ParseProfile(contents []byte) engine.ParserResult {
	// 用户模型
	profile := model.Profile{}
	// 图片
	profile.Avatar = extractString(contents, avatarReg)
	// 昵称
	profile.Nickname = extractString(contents, nickname)
	// 个人简介
	profile.Introduction = extractString(contents, introduction)
	// 其它信息
	extractOtherInfo(contents, otherInfo, &profile)

	// 响应返回
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}

// 基本信息
func extractString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	}
	return ""
}

// 其它信息
func extractOtherInfo(contents []byte, reg *regexp.Regexp, p *model.Profile) {
	match := reg.FindAllSubmatch(contents, -1)

	for _, i := range match {
		if len(i) > 6 {
			p.City = string(i[1])
			p.Age, _ = strconv.Atoi(string(i[2]))
			p.Education = string(i[3])
			p.Marriage = string(i[4])
			p.Height, _ = strconv.Atoi(string(i[5]))
			p.Income = string(i[6])
			// fmt.Printf("结果 %s\n", i[1])
			// for _, j := range i {
			// 	fmt.Printf("结果 %s\n", j)
			// }
		}
	}
}
