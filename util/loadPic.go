package util

import (
	"bytes"
	"image"
	"io"

	// Required for png decoding
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/wayovertheregaming/catastrophy/assets"
	"github.com/wayovertheregaming/catastrophy/catlog"
)

// LoadPic will load an image from an asset file and get the pixel Picture
// data from it
func LoadPic(path string) pixel.Picture {
	img, _, err := image.Decode(GetReaderFromFile(path))
	if err != nil {
		catlog.Fatal(err)
	}

	return pixel.PictureDataFromImage(img)
}

// GetReaderFromFile creates an io reader from filepath
func GetReaderFromFile(path string) io.Reader {
	f, err := assets.Asset(path)
	if err != nil {
		catlog.Fatal(err)
	}

	return bytes.NewReader(f)
}

// LoadSprite will load a sprite from a give file path and rectangle
// The `r` parameter is the location of the desired sprite from the file
func LoadSprite(path string, r pixel.Rect) (*pixel.Sprite, pixel.Picture) {
	pic := LoadPic(path)
	return pixel.NewSprite(pic, r), pic
}
