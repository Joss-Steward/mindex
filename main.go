package main

import (
	"flag"
	"mIndex/crawler"
)

func main() {
	hashPtr := flag.Bool("hash", false, "compute hash for media file deduplication")

	flag.Parse()

	config := &crawler.Config{
		Hash: *hashPtr,
	}

	crawler := crawler.NewCrawler("/Volumes/media/movies", *config)

	crawler.Crawl()
}
