package main

import (
	"image"
	"image/jpeg"
	"os"
	"time"
)

type jpegImage struct {
}

func (j *jpegImage) doConvert(fileUploaded *fileUploaded) (string, error) {
	f, err := os.Open(fileUploaded.location + "/" + fileUploaded.fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	dateFormatted := time.Now().Format("20060102030405")
	newFileName := dateFormatted + fileUploaded.onlyFilename + ".jpeg"
	f, err = os.Create(newFileName)
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
