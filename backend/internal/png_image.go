package internal

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/nandes007/image-converter/pkg"
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
	newFileName := dateFormatted + pkg.FileNameWithoutExtSliceNotation(fileUploaded.fileName) + ".png"
	fileLocation, err := pkg.GetConvertedDirectory()
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
