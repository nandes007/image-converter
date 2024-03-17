package main

import "fmt"

type iConverterFactory interface {
	doConvert() (string, error)
}

func convertImage(extensionType string) (iConverterFactory, error) {
	if extensionType == "png" {
		return &pngImage{}, nil
	}
	if extensionType == "jpeg" {
		return &jpegImage{}, nil
	}
	return nil, fmt.Errorf("wrong extension type passed")
}
