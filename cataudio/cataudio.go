package cataudio

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
)

const (
	// dirPrefix is the directory path for all audio files
	dirPrefix = "assets/audio"

	sampleRate = 44800
)

var (
	// bufferSize is amount of the time the audio file is loaded before it starts
	// playing
	bufferSize = time.Second / 10

	// allAudioFiles holds all playable audio files
	allAudioFiles = map[string]*audio{
		"Menu.mp3":            {loops: -1},
		"PlayingInGarden.mp3": {loops: -1},
		"PlayingInHouse.mp3":  {loops: -1},
		"ShadowRealm.mp3":     {loops: -1},
		"BirdsChirping.mp3":   {loops: 0},
		"CatDrinking.mp3":     {loops: 0},
		"CatPeeing.mp3":       {loops: 0},
		"CatPoop.mp3":         {loops: 0},
		"CatWalking.mp3":      {loops: 0},
		"DoorClosed.mp3":      {loops: 0},
		"DoorOpen.mp3":        {loops: 0},
		"Excited.mp3":         {loops: 0},
		"Snoring.mp3":         {loops: 0},
		"TvNoise.mp3":         {loops: 0},
		"TvStatic.mp3":        {loops: 0},
		"WhistlingFall.mp3":   {loops: 0},
	}

	// speakerRunning will prevent playing to an uninitialised speaker, in case of error
	speakerRunning = false
)

type audio struct {
	streamer beep.Streamer
	format   beep.Format
	loops    int
}

func (a *audio) play() {
	speaker.Play(a.streamer)
}

func init() {
	catlog.Debug("Doing cataudio init")

	if err := speaker.Init(sampleRate, beep.SampleRate(sampleRate).N(bufferSize)); err != nil {
		catlog.Infof("Failed to init speaker, no audio will play: %v", err)
		return
	}
	speakerRunning = true

	for filename, a := range allAudioFiles {
		// Create the full path to the file
		fullPath := filepath.Join(dirPrefix, filename)
		// Get the streamseaker and format
		ss, f := mustLoadAudioFile(fullPath)

		// Loop, converting to a streamer
		s := beep.Loop(a.loops, ss)

		a.streamer = s
		a.format = f

		catlog.Debugf("Initialised audio file %s", fullPath)
	}
}

// Play will attempt to play a sound
func Play(filename string) {
	if !speakerRunning {
		catlog.Debug("Not playing file, speaker not initialised")
		return
	}

	catlog.Debugf("Attempting to play %s", filename)

	// Check if audio file exists
	if aud, ok := allAudioFiles[filename]; ok {
		// If it does, play then return
		aud.play()
		return
	}

	catlog.Info("Could not find %s while trying to play audio, will not play", filename)
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
	case ".wav":
		decodeFunc = wav.Decode
	case ".mp3":
		decodeFunc = mp3.Decode
	case ".flac":
		decodeFunc = flac.Decode
	case ".ogg":
		decodeFunc = vorbis.Decode
	default:
		catlog.Fatalf("Cannot determine decode function for %s.  Extension '%s' not found", path, ext)
	}

	s, format, err := decodeFunc(fReadCloser)
	if err != nil {
		catlog.Fatalf("Cannot decode %s file: %v", ext, err)
	}

	return s, format
}
