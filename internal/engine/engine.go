package engine

import (
	"gocrawl/internal/downloader"
	"gocrawl/internal/pipelines"
	"gocrawl/internal/scheduler"
	"gocrawl/internal/spider"
)

type Engine struct {
	Scheduler  *scheduler.Scheduler
	Downloader *downloader.Downloader
	Spider     *spider.Spider
	Pipelines  *pipelines.Pipelines
}

func NewEngine() *Engine {
	return &Engine{
		Scheduler:  scheduler.NewScheduler(),
		Downloader: downloader.NewDownloader(),
		Spider:     spider.NewSpider(),
		Pipelines:  pipelines.NewPipelines(),
	}
}

func (e *Engine) Run() {
	for {
		req := e.Scheduler.NextRequest()
		if req == nil {
			break
		}
		resp, err := e.Downloader.Download(req.URL)
		if err != nil {
			continue
		}
		items := e.Spider.Parse(resp)
		for _, item := range items {
			e.Pipelines.Process(item)
		}
	}
}
