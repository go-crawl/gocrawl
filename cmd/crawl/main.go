package crawl

import (
	"fmt"
	"gocrawl/internal/engine"

	"github.com/spf13/cobra"
)

var CrawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Run the spider",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crawling...")
		e := engine.NewEngine()
		e.Run()
	},
}
