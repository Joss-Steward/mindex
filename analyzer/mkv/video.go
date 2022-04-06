package mkv

import "errors"

// Error Types
var (
	ErrorVideoElementNotStarted = errors.New("video element parameter was found outside of video element")

	ErrorDuplicatePixelWidth    = errors.New("video pixel width token was found but pixel width is already set")
	ErrorDuplicatePixelHeight   = errors.New("video pixel height token was found but pixel height is already set")
	ErrorDuplicateDisplayWidth  = errors.New("video display width token was found but display width is already set")
	ErrorDuplicateDisplayHeight = errors.New("video display height token was found but display height is already set")
)

type Video struct {
	PixelWidth    *int
	PixelHeight   *int
	DisplayWidth  *int
	DisplayHeight *int
}

func (v *Video) SetPixelWidth(width int) error {
	if v == nil {
		return ErrorVideoElementNotStarted
	}

	if v.PixelWidth != nil {
		return ErrorDuplicatePixelWidth
	}

	v.PixelWidth = &width
	return nil
}

func (v *Video) SetPixelHeight(height int) error {
	if v == nil {
		return ErrorVideoElementNotStarted
	}

	if v.PixelHeight != nil {
		return ErrorDuplicatePixelHeight
	}

	v.PixelHeight = &height
	return nil
}

func (v *Video) SetDisplayWidth(width int) error {
	if v == nil {
		return ErrorVideoElementNotStarted
	}

	if v.DisplayWidth != nil {
		return ErrorDuplicateDisplayWidth
	}

	v.DisplayWidth = &width
	return nil
}

func (v *Video) SetDisplayHeight(height int) error {
	if v == nil {
		return ErrorVideoElementNotStarted
	}

	if v.DisplayHeight != nil {
		return ErrorDuplicateDisplayHeight
	}

	v.DisplayHeight = &height
	return nil
}
