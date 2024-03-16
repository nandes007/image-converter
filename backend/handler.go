package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nandes007/image-converter/helper"
)

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath, err := helper.ImageConverterToPngFromJpg()
	if err != nil {
		fmt.Fprintf(w, "Error opening file 1: %v", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(w, "Error opening file 2: %v - %s ", err, filePath)
		return
	}
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "Error read file 3: %v - %s", err, file.Name())
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(fileByte))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))

	http.ServeFile(w, r, filePath)
}
