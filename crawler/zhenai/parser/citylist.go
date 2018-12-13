package parser

import (
	"imooc.com/ccmouse/learngo/crawler/engine"
	"regexp"
)

var CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 获取城市列表
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(CityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}