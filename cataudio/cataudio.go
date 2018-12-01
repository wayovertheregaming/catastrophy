package cataudio

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
)

func init() {

}

func mustLoadAudioFile(path string) (beep.StreamSeekCloser, beep.Format) {
	catlog.Debugf("Loading audio file %s", path)

	f, err := assets.Asset(path)
	if err != nil {
		catlog.Fatalf("Cannot load audio file: %v", err)
	}

	fReadCloser := ioutil.NopCloser(bytes.NewReader(f))

	// Get the extension to determine the type
	ext := filepath.Ext(path)

	// Select the correct function to decode based on the extension
	var decodeFunc func(io.ReadCloser) (beep.StreamSeekCloser, beep.Format, error)
	switch ext {
	case "wav":
		decodeFunc = wav.Decode
	case "mp3":
		decodeFunc = mp3.Decode
	case "flac":
		decodeFunc = flac.Decode
	case "ogg":
		decodeFunc = vorbis.Decode
	default:
		catlog.Fatalf("Cannot determine decode function for %s", path)
	}

	s, format, err := decodeFunc(fReadCloser)
	if err != nil {
		catlog.Fatalf("Cannot decode %s file: %v", ext, err)
	}

	return s, format
}
