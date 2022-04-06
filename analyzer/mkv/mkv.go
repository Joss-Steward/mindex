package mkv

import (
	"fmt"
	"os"

	"github.com/remko/go-mkvparse"
)

type MKV struct {
	Title  string
	Tracks []Track
}

type parser struct {
	mkvparse.DefaultHandler

	mkv   *MKV
	track *Track
}

func Parse(path string) (*MKV, error) {
	p := newParser()

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s': %w", path, err)
	}

	defer file.Close()

	err = mkvparse.ParseSections(file, mkvparse.NewHandlerChain(&p), mkvparse.InfoElement, mkvparse.TracksElement)
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s': %w", path, err)
	}

	return p.mkv, nil
}

func newParser() parser {
	p := parser{
		mkv: new(MKV),
	}

	return p
}

func (p *parser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TitleElement:
		p.mkv.Title = value
	}

	return nil
}

func (p *parser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
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

func (p *parser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	fmt.Printf("Master Begin")

	switch id {
	case mkvparse.CuesElement:
		return false, nil
	case mkvparse.ClusterElement:
		return false, nil
	default:
		fmt.Printf("%d- %s:\n", info.Level, mkvparse.NameForElementID(id))
		return true, nil
	}
}

func (p *parser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	fmt.Printf("Master End")

	return nil
}
