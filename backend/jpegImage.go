package main

import (
	"image"
	"image/jpeg"
	"os"
)

type jpegImage struct {
}

func (j *jpegImage) doConvert() (string, error) {
	filePath := "image/"
	f, err := os.Open(filePath + "test.jpg")
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	f, err = os.Create(filePath + "jpg_to_jpeg.jpeg")
	if err != nil {
		return "", err
	}

	opt := jpeg.Options{
		Quality: 30,
	}

	err = jpeg.Encode(f, img, &opt)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
