package parser

import (
	"GoCrawler/engine"
	"regexp"
)

// pageListRe 分类下的影视分页列表
const pageListRe = `<option value='(list_\d+_\d+\.html)'( selected)?>(\d+)</option>`

// movieListRe 每个影视分页下的影视列表匹配规则
const movieListRe = `<a href="(/html/[0-9a-z]+/[0-9a-z]+/[0-9a-z]+/[0-9a-z]+\.html)" class="ulink">([^>]*[^<])</a>`

// ParsePageList 解析分页列表
func ParsePageList(contents []byte, url string) engine.ParserResult {
	re := regexp.MustCompile(pageListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matches {
		//result.Items = append(result.Items, string(m[3]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        url + string(m[1]),
			ParserFunc: ParseMovieList,
		})
	}
	return result
}

// ParseMovieList 解析影视列表
func ParseMovieList(contents []byte, url string) engine.ParserResult {
	re := regexp.MustCompile(movieListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	// tmpMap 去重
	tmpMap := make(map[string]interface{})
	for _, m := range matches {
		if _, ok := tmpMap[string(m[2])]; !ok {
			//result.Items = append(result.Items, string(m[2]))
			result.Requests = append(result.Requests, engine.Request{
				Url:        PrefixUrl + string(m[1]),
				ParserFunc: ParseMovieInfo,
			})
			tmpMap[string(m[2])] = nil
		}
	}
	return result
}
