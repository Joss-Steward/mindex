package mkv

import "errors"

// Error Types
var (
	ErrorAudioElementNotStarted = errors.New("audio element parameter was found outside of audio element")

	ErrorDuplicateSampleFreq = errors.New("audio sample freq token was found but sample freq is already set")
	ErrorDuplicateChannels   = errors.New("audio channels token was found but channels is already set")
)

type Audio struct {
	SampleFreq *int
	Channels   *int
}

func (a *Audio) SetSampleFreq(frequency int) error {
	if a == nil {
		return ErrorAudioElementNotStarted
	}

	if a.SampleFreq != nil {
		return ErrorDuplicateSampleFreq
	}

	a.SampleFreq = &frequency
	return nil
}

func (a *Audio) SetChannels(channels int) error {
	if a == nil {
		return ErrorAudioElementNotStarted
	}

	if a.Channels != nil {
		return ErrorDuplicateChannels
	}

	a.Channels = &channels
	return nil
}
