package engine

type Request struct {
	Url        string
	ParserFunc func([]byte, string) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte, string) ParserResult {
	return ParserResult{}
}
