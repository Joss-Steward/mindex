package crawler

import (
	"fmt"
	"log"
	"mIndex/analyzer"
	"os"
	"path/filepath"
)

type Config struct {
	Hash bool
}

type Crawler struct {
	base   string
	config Config
}

func NewCrawler(base string, config Config) Crawler {
	return Crawler{
		base:   base,
		config: config,
	}
}

func (c Crawler) Crawl() {
	err := filepath.Walk(c.base, c.handleFile)

	if err != nil {
		log.Fatalf("failed to crawl directory: %s", err.Error())
	}
}

func (crawler *Crawler) handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error handling file: %w", err)
	}

	if info.IsDir() {
		return nil
	}

	mf, err := analyzer.Analyze(path, info, crawler.config.Hash)
	if err != nil {
		// log.Printf("skipping '%s': %s", path, err.Error())
	} else {
		fmt.Println(mf.String())
	}

	return nil
}
