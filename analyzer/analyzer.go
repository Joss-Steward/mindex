package analyzer

import (
	"fmt"
	"mIndex/analyzer/mkv"
	"os"
	"path/filepath"
)

type FileHandler func(string, os.FileInfo) (*MediaFile, error)

var ExtensionTable = map[string]FileHandler{
	".mkv":  HandleMKV,
	".mka":  HandleMKV,
	".webm": HandleMKV,
}

func Analyze(path string, info os.FileInfo, hash bool) (*MediaFile, error) {
	extension := filepath.Ext(path)

	handler, ok := ExtensionTable[extension]

	if ok {
		mf, err := handler(path, info)
		if err != nil {
			return nil, fmt.Errorf("could not parse mkv file '%s': %w", path, err)
		}

		if hash {
			h, err := hashFile(path)
			if err != nil {
				return nil, fmt.Errorf("could not hash file '%s': %w", path, err)
			}

			mf.Hash = h
		}

		return mf, nil
	}

	return nil, fmt.Errorf("unknown file type: '%s'", extension)
}

func HandleMKV(path string, info os.FileInfo) (*MediaFile, error) {
	mkvInfo, err := mkv.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse file '%s': %w", path, err)
	}

	mf := NewMediaFile(path, info.ModTime(), info.Size())

	if mkvInfo.Title != nil {
		mf.Title = *mkvInfo.Title
	}

	for _, t := range mkvInfo.Tracks {
		trackBase := NewTrack(t.Codec, t.Language, t.Default, t.Forced, t.ID)

		switch *t.TrackType {
		case mkv.TypeVideo:
			if t.Video == nil {
				return nil, fmt.Errorf("video track type found without video information")
			}

			v := NewVideo(trackBase, t.Video.PixelHeight, t.Video.PixelWidth)

			mf.VideoTracks = append(mf.VideoTracks, v)
			continue
		case mkv.TypeAudio:
			if t.Audio == nil {
				return nil, fmt.Errorf("audio track type found without audio information")
			}

			a := NewAudio(trackBase, t.Audio.Channels, t.Audio.SampleFreq)

			mf.AudioTracks = append(mf.AudioTracks, a)
			continue
		case mkv.TypeSubtitle:
			s := NewSubtitle(trackBase)

			mf.SubtitleTracks = append(mf.SubtitleTracks, s)
			continue
		}
	}

	return mf, nil
}
