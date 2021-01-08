package parser

import (
	"gowork/crawler/engine"
	"regexp"
)

// typeListRe 影视类别匹配规则
const typeListRe = `<a href="(/html/gndy(/[0-9a-z]+)/)index\.html">([^>]*[^<])</a>`

// ParseTypeList 解析影视类别
func ParseTypeList(contents []byte, url string) engine.ParserResult {
	re := regexp.MustCompile(typeListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	// tmpMap 去重
	tmpMap := make(map[string]interface{})
	for _, m := range matches {
		if _, ok := tmpMap[string(m[3])]; !ok {
			//result.Items = append(result.Items, string(m[3]))
			result.Requests = append(result.Requests, engine.Request{
				Url:        url + string(m[1]),
				ParserFunc: ParsePageList,
			})
			tmpMap[string(m[3])] = nil
		}
	}
	return result
}
