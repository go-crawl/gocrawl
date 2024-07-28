package spider

import (
	"github.com/gocolly/colly"
)

type Spider struct {
}

func NewSpider() *Spider {
	return &Spider{}
}

func (s *Spider) Parse(resp *colly.Response) []interface{} {
	items := []interface{}{}
	// 解析逻辑
	return items
}
