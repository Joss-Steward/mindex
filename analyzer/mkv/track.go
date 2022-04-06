package mkv

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
	number   int
	kind     int
	codec    string
	language string

	video *Video
	audio *Audio
}

type Video struct {
	PixelWidth    int
	PixelHeight   int
	DisplayWidth  int
	DisplayHeight int
}

type Audio struct {
	SampleFreq int
	Channels   int
}
