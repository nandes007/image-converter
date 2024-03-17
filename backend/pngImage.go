package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

type pngImage struct {
}

func (p *pngImage) doConvert(fileUploaded *fileUploaded) (string, error) {
	f, err := os.Open(fileUploaded.fullPathFile)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	dateFormatted := time.Now().Format("20060102030405")
	newFileName := dateFormatted + fileNameWithoutExtSliceNotation(fileUploaded.fileName) + ".png"
	fileLocation, err := getConvertedDirectory()
	if err != nil {
		return "", err
	}

	f, err = os.Create(filepath.Join(fileLocation, newFileName))
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
