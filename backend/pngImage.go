package main

import (
	"image"
	"image/png"
	"os"
)

type pngImage struct {
}

func (p *pngImage) doConvert() (string, error) {
	filePath := "image/"
	f, err := os.Open(filePath + "test.jpg")
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	f, err = os.Create(filePath + "jpg_to_png.png")
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
