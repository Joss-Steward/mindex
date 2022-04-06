package mkv

import "errors"

// Error Values
var (
	ErrorDuplicateTitle = errors.New("title token found but title is already set")
)

type Info struct {
	Title  *string
	Tracks []*Track
}

func NewMKV() *Info {
	mkvInfo := &Info{
		Tracks: make([]*Track, 0),
	}

	return mkvInfo
}

func (i *Info) SetTitle(title string) error {
	if i.Title != nil {
		return ErrorDuplicateTitle
	}

	i.Title = &title
	return nil
}
