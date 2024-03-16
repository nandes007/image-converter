package helper

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func ImageConverterToJpegFromJpg() (string, error) {
	f, err := os.Open("test.jpg")
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	f, err = os.Create("jpg_to_jpeg.jpeg")
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

	return "Successfully converted image to jpeg", nil
}

func ImageConverterToPngFromJpg() (string, error) {
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
	fileName := f.Name()
	err = png.Encode(f, img)
	if err != nil {
		return "", nil
	}

	filePath = fileName

	return filePath, nil
}
