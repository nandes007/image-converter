package main

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func getConvertedDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dirLocation := filepath.Join(dir, "converted_images")
	if _, err := os.Stat(dirLocation); os.IsNotExist(err) {
		err := os.Mkdir(dirLocation, 0750)
		if err != nil {
			return "", err
		}
	}

	return dirLocation, nil
}

func uploadFile(uploadedFile multipart.File, handler *multipart.FileHeader) (*fileUploaded, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dirLocation := filepath.Join(dir, "images")
	if _, err := os.Stat(dirLocation); os.IsNotExist(err) {
		err := os.Mkdir(dirLocation, 0750)
		if err != nil {
			return nil, err
		}
	}

	filename := handler.Filename
	fileLocation := filepath.Join(dirLocation, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return nil, err
	}

	fullPathFile := dirLocation + "/" + filename

	return &fileUploaded{
		fileName:     filename,
		fullPathFile: fullPathFile,
	}, nil
}
