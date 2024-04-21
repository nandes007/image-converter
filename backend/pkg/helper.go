package pkg

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func GetFileName(fileName string) string {
	splitStr := strings.Split(fileName, "/")
	return splitStr[len(splitStr)-1]
}

func GetConvertedDirectory() (string, error) {
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

func ValidateRequest(r *http.Request) error {
	trimmedStr := strings.TrimSpace(r.FormValue("convert_to"))
	if trimmedStr == "" {
		return errors.New("option is required")
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return errors.New("10 MB miximum file size")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return errors.New("file is required")
	}

	defer file.Close()
	return nil
}
