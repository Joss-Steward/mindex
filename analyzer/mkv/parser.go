package mkv

import (
	"errors"
	"fmt"
	"os"

	"github.com/remko/go-mkvparse"
)

var (
	ErrorNestedTracks         = errors.New("track element start token found inside track element")
	ErrorTrackEndWithoutTrack = errors.New("track element end token found outside of track element")
)

type parser struct {
	mkvparse.DefaultHandler

	mkvInfo *Info
	track   *Track
}

func Parse(path string) (*Info, error) {
	p := NewParser()

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s': %w", path, err)
	}

	defer file.Close()

	err = mkvparse.ParseSections(file, mkvparse.NewHandlerChain(p), mkvparse.InfoElement, mkvparse.TracksElement)
	if err != nil {
		return nil, fmt.Errorf("could not parse '%s': %w", path, err)
	}

	return p.mkvInfo, nil
}

func NewParser() *parser {
	p := &parser{
		mkvInfo: new(Info),
	}

	return p
}

func (p *parser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TitleElement:
		return p.mkvInfo.SetTitle(value)
	case mkvparse.LanguageElement:
		return p.track.SetLanguage(value)
	case mkvparse.CodecNameElement:
		return p.track.SetCodec(value)
	}

	return nil
}

func (p *parser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TrackNumberElement:
		return p.track.SetID(value)
	case mkvparse.TrackTypeElement:
		return p.track.SetType(value)
	case mkvparse.PixelHeightElement:
		if p.track.videoActive {
			return p.track.Video.SetPixelHeight(int(value))
		}

		return ErrorVideoElementNotStarted
	case mkvparse.PixelWidthElement:
		if p.track.videoActive {
			return p.track.Video.SetPixelWidth(int(value))
		}

		return ErrorVideoElementNotStarted
	case mkvparse.DisplayHeightElement:
		if p.track.videoActive {
			return p.track.Video.SetDisplayHeight(int(value))
		}

		return ErrorVideoElementNotStarted
	case mkvparse.DisplayWidthElement:
		if p.track.videoActive {
			return p.track.Video.SetDisplayWidth(int(value))
		}

		return ErrorVideoElementNotStarted
	case mkvparse.ChannelsElement:
		if p.track.audioActive {
			return p.track.Audio.SetChannels(int(value))
		}

		return ErrorAudioElementNotStarted
	}

	return nil
}

func (p *parser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.SamplingFrequencyElement:
		if p.track.audioActive {
			return p.track.Audio.SetSampleFreq(int(value))
		}

		return ErrorAudioElementNotStarted
	}

	return nil
}

func (p *parser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	switch id {
	case mkvparse.SegmentElement:
		return true, nil
	case mkvparse.TracksElement:
		return true, nil
	case mkvparse.TrackEntryElement:
		if p.track != nil {
			return false, ErrorNestedTracks
		}

		p.track = NewTrack()
		return true, nil
	case mkvparse.VideoElement:
		return true, p.track.BeginVideoElement()
	case mkvparse.AudioElement:
		return true, p.track.BeginAudioElement()
	}

	return false, nil
}

func (p *parser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	switch id {
	case mkvparse.TrackEntryElement:
		if p.track == nil {
			return ErrorTrackEndWithoutTrack
		}

		p.mkvInfo.Tracks = append(p.mkvInfo.Tracks, p.track)
		p.track = nil

		return nil
	case mkvparse.VideoElement:
		return p.track.EndVideoElement()
	case mkvparse.AudioElement:
		return p.track.EndAudioElement()
	}

	return nil
}
