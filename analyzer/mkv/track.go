package mkv

import (
	"errors"
)

// Error Types
var (
	ErrorAudioEndWithoutAudio     = errors.New("video element end token found outside of video element")
	ErrorAudioEndWithoutTrack     = errors.New("audio element end token found outside of track element")
	ErrorAudioNestedStart         = errors.New("audio element start token found nested inside audio element")
	ErrorAudioStartWithoutTrack   = errors.New("audio element start token found outside of track")
	ErrorAudioVideoElementOverlap = errors.New("audio element start token found in track with audio element")

	ErrorVideoAudioElementOverlap = errors.New("video element start token found in track with audio element")
	ErrorVideoEndWithoutTrack     = errors.New("video element end token found outside of track element")
	ErrorVideoEndWithoutVideo     = errors.New("video element end token found outside of video element")
	ErrorVideoNestedStart         = errors.New("video element start token found nested inside video element")
	ErrorVideoStartWithoutTrack   = errors.New("video element start token found outside of track element")

	ErrorDuplicateTrackIDToken   = errors.New("track ID token found but track ID is already set")
	ErrorDuplicateTrackTypeToken = errors.New("track type token found but track type is already set")
	ErrorDuplicateCodecToken     = errors.New("track codec token found but track codec is already set")
	ErrorDuplicateLanguageToken  = errors.New("track language token found but track language is already set")
	ErrorDuplicatedDefaultFlag   = errors.New("track default flag token found but track default is already set")
	ErrorDuplicatedForcedFlag    = errors.New("track forced flag token found but track default is already set")

	ErrorTrackIDWithoutTrack     = errors.New("track ID token found without track element")
	ErrorTrackTypeWithoutTrack   = errors.New("track type token found without track element")
	ErrorCodecWithoutTrack       = errors.New("track codec token found without track element")
	ErrorLanguageWithoutTrack    = errors.New("track language token found without track element")
	ErrorDefaultFlagWithoutTrack = errors.New("track default flag token found without track element")
	ErrorForcedFlagWithoutTrack  = errors.New("track forced flag token found without track element")
)

// https://www.matroska.org/technical/elements.html
const (
	TypeVideo    = 1
	TypeAudio    = 2
	TypeComplex  = 3
	TypeLogo     = 16
	TypeSubtitle = 17
	TypeButtons  = 18
	TypeControl  = 32
	TypeMetadata = 33
)

type Track struct {
	audioActive bool
	videoActive bool

	ID        *int64
	TrackType *int64
	Codec     *string
	Language  *string
	Default   *bool
	Forced    *bool

	Video *Video
	Audio *Audio
}

func NewTrack() *Track {
	return &Track{
		audioActive: false,
		videoActive: false,
	}
}

func (t *Track) BeginVideoElement() error {
	if t == nil {
		return ErrorVideoStartWithoutTrack
	}

	if t.Video != nil {
		return ErrorVideoNestedStart
	}

	if t.Audio != nil {
		return ErrorVideoAudioElementOverlap
	}

	t.videoActive = true
	t.Video = &Video{}

	return nil
}

func (t *Track) CanSetVideo() bool {
	return t.videoActive
}

func (t *Track) EndVideoElement() error {
	if t == nil {
		return ErrorVideoEndWithoutTrack
	}

	if t.Video == nil {
		return ErrorVideoEndWithoutVideo
	}

	t.videoActive = false
	return nil
}

func (t *Track) BeginAudioElement() error {
	if t == nil {
		return ErrorAudioStartWithoutTrack
	}

	if t.Audio != nil {
		return ErrorAudioNestedStart
	}

	if t.Video != nil {
		return ErrorAudioVideoElementOverlap
	}

	t.audioActive = true
	t.Audio = &Audio{}

	return nil
}

func (t *Track) CanSetAudio() bool {
	return t.audioActive
}

func (t *Track) EndAudioElement() error {
	if t == nil {
		return ErrorAudioEndWithoutTrack
	}

	if t.Audio == nil {
		return ErrorAudioEndWithoutAudio
	}

	t.audioActive = false
	return nil
}

func (t *Track) SetID(id int64) error {
	if t == nil {
		return ErrorTrackIDWithoutTrack
	}

	if t.ID != nil {
		return ErrorDuplicateTrackIDToken
	}

	t.ID = &id
	return nil
}

func (t *Track) SetType(trackType int64) error {
	if t == nil {
		return ErrorTrackTypeWithoutTrack
	}

	if t.TrackType != nil {
		return ErrorDuplicateTrackTypeToken
	}

	t.TrackType = &trackType
	return nil
}

func (t *Track) SetCodec(codec string) error {
	if t == nil {
		return ErrorCodecWithoutTrack
	}

	if t.Codec != nil {
		return ErrorDuplicateCodecToken
	}

	t.Codec = &codec
	return nil
}

func (t *Track) SetLanguage(language string) error {
	if t == nil {
		return ErrorLanguageWithoutTrack
	}

	if t.Language != nil {
		return ErrorDuplicateLanguageToken
	}

	t.Language = &language
	return nil
}

func (t *Track) SetFlagDefault(flag bool) error {
	if t == nil {
		return ErrorDefaultFlagWithoutTrack
	}

	if t.Default != nil {
		return ErrorDuplicatedDefaultFlag
	}

	t.Default = &flag
	return nil
}

func (t *Track) SetFlagForced(flag bool) error {
	if t == nil {
		return ErrorForcedFlagWithoutTrack
	}

	if t.Forced != nil {
		return ErrorDuplicatedForcedFlag
	}

	t.Forced = &flag
	return nil
}
