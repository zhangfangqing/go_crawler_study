package parser

import (
	"go_crawler_study/crawler/engine"
	"regexp"
)

const cityRegexp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) (engine.ParseResult) {
	re := regexp.MustCompile(cityRegexp)
	matchersArr := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matchersArr {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: engine.NilParser})
	}
	return result
}
