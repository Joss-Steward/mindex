package analyzer

import (
	"fmt"
	"time"
)

type MediaFile struct {
	Path    string
	Created time.Time
	Size    int64
	Title   string
	Hash    []byte

	AudioTracks    []Audio
	VideoTracks    []Video
	SubtitleTracks []Subtitle
}

type Track struct {
	Codec    string
	Language string
	Default  bool
	Forced   bool
	Number   int
}

type Audio struct {
	Track
	Channels   int
	SampleFreq int
}

type Video struct {
	Track
	Height int
	Width  int
}

type Subtitle struct {
	Track
}

func NewMediaFile(path string, created time.Time, size int64) *MediaFile {
	f := &MediaFile{
		Path:    path,
		Created: created,
		Size:    size,
		Title:   "",

		AudioTracks:    make([]Audio, 0),
		VideoTracks:    make([]Video, 0),
		SubtitleTracks: make([]Subtitle, 0),
	}

	return f
}

func NewTrack(
	codec *string,
	language *string,
	defaultFlag *bool,
	forcedFlag *bool,
	number *int64,
) Track {
	track := Track{
		Codec:    "UNDEF",
		Language: "und",
		Default:  false,
		Forced:   false,
		Number:   0,
	}

	if codec != nil {
		track.Codec = *codec
	}

	if language != nil {
		track.Language = *language
	}

	if defaultFlag != nil {
		track.Default = *defaultFlag
	}

	if forcedFlag != nil {
		track.Forced = *forcedFlag
	}

	if number != nil {
		track.Number = int(*number)
	}

	return track
}

func NewAudio(track Track, channels *int, sampleFreq *int) Audio {
	audio := Audio{
		Track:      track,
		Channels:   0,
		SampleFreq: 0,
	}

	if channels != nil {
		audio.Channels = *channels
	}

	if sampleFreq != nil {
		audio.SampleFreq = *sampleFreq
	}

	return audio
}

func NewVideo(track Track, height *int, width *int) Video {
	video := Video{
		Track:  track,
		Height: 0,
		Width:  0,
	}

	if height != nil {
		video.Height = *height
	}

	if width != nil {
		video.Width = *width
	}

	return video
}

func NewSubtitle(track Track) Subtitle {
	return Subtitle{Track: track}
}

func (mf MediaFile) String() string {
	result := fmt.Sprintf("%s\n", mf.Path)

	result = result + fmt.Sprintf("'%s' Size: %s\n", mf.Title, byteCountIEC(mf.Size))
	result = result + fmt.Sprintf("Hash: %x\n", mf.Hash)

	result = result + "Video:\n"
	for _, v := range mf.VideoTracks {
		result = result + "  " + v.String() + "\n"
	}

	result = result + "Audio:\n"
	for _, a := range mf.AudioTracks {
		result = result + "  " + a.String() + "\n"
	}

	result = result + "Subtitles:\n"
	for _, s := range mf.SubtitleTracks {
		result = result + "  " + s.String() + "\n"
	}

	return result
}

func (v Video) String() string {
	dflt := " "

	if v.Default || v.Forced {
		dflt = "*"
	}

	codec := fmt.Sprintf("[%s]", v.Codec)
	return fmt.Sprintf("%s%d %3s %5dx%d %12s", dflt, v.Number, v.Language, v.Width, v.Height, codec)
}

func (a Audio) String() string {
	dflt := " "

	if a.Default || a.Forced {
		dflt = "*"
	}

	codec := fmt.Sprintf("[%s]", a.Codec)
	return fmt.Sprintf("%s%d %3s %2d channels %8dhz%12s", dflt, a.Number, a.Language, a.Channels, a.SampleFreq, codec)
}

func (s Subtitle) String() string {
	dflt := " "

	if s.Default || s.Forced {
		dflt = "*"
	}

	return fmt.Sprintf("%s%d %s", dflt, s.Number, s.Language)
}
