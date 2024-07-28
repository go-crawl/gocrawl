package downloader

import (
	"github.com/gocolly/colly"
)

type Downloader struct {
	Collector *colly.Collector
}

func NewDownloader() *Downloader {
	c := colly.NewCollector()
	return &Downloader{Collector: c}
}

func (d *Downloader) Download(url string) (*colly.Response, error) {
	var resp *colly.Response
	err := d.Collector.Visit(url)
	if err != nil {
		return nil, err
	}
	d.Collector.OnResponse(func(r *colly.Response) {
		resp = r
	})
	return resp, nil
}
