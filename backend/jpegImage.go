package main

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"
)

type jpegImage struct {
}

func (j *jpegImage) doConvert(fileUploaded *fileUploaded) (string, error) {
	f, err := os.Open(fileUploaded.fullPathFile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	dateFormatted := time.Now().Format("20060102030405")
	newFileName := dateFormatted + fileNameWithoutExtSliceNotation(fileUploaded.fileName) + ".jpeg"
	fileLocation, err := getConvertedDirectory()
	if err != nil {
		return "", err
	}

	f, err = os.Create(filepath.Join(fileLocation, newFileName))
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
