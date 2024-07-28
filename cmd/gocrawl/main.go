package main

import (
	"fmt"
	"os"

	"gocrawl/cmd/crawl"
	"gocrawl/cmd/startproject"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocrawl",
	Short: "Go Crawler",
	Long:  "A web crawling framework written in Go.",
}

func main() {
	rootCmd.AddCommand(startproject.StartprojectCmd)
	rootCmd.AddCommand(crawl.CrawlCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
