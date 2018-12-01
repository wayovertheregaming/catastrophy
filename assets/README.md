# Assets
This directory contains all non .go assets that need to be used for the game.

Using the bin-asset application we can turn these into binary and not have to
distribute them with the compiled game.

This process is all done in the `run.sh` script.  When you want to access the
asset file in code, use `assets.Asset("assets/audio/test.mp3")` (example).  This
will the byte slice and the `os.FileInfo`.

## Pixel
In pixel, to use a picture you have a function like this (taken from the
tutorials):
```go
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
```

This changes to the following with bin-assets:
```go
func loadPicture(path string) (pixel.Picture, error) {
	fileData, _ := assets.Asset(path)
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
```
