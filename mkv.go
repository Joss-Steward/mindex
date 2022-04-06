package main

import (
	"fmt"
	"os"

	"github.com/remko/go-mkvparse"
)

type MyParser struct {
	mkvparse.DefaultHandler

	title *string
}

func (p *MyParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TitleElement:
		p.title = &value
	}

	return nil
}

func (p *MyParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TrackNumberElement:
		fmt.Printf("Track (int): %d\n", value)
	case mkvparse.PixelHeightElement:
		fmt.Printf("Display Height (int): %d\n", value)
	case mkvparse.PixelWidthElement:
		fmt.Printf("Display Width (int): %d\n", value)
	}
	return nil
}

func getTags(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s': %w", path, err)
	}

	defer file.Close()

	titleh := MyParser{}

	err = mkvparse.ParseSections(file, mkvparse.NewHandlerChain(&titleh), mkvparse.InfoElement, mkvparse.TracksElement)
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s': %w", path, err)
	}

	// Print (sorted) tags
	if titleh.title != nil {
		fmt.Printf("- title: %q\n", *titleh.title)
	}

	return nil, nil
}
