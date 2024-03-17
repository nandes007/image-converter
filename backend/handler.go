package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "OK",
		"code":   http.StatusOK,
	})
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	imageConverter, err := convertImage("jpeg")
	if err != nil {
		fmt.Printf("Error when passed convert image: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	filePath, err := imageConverter.doConvert()
	if err != nil {
		fmt.Printf("Error convert file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error read byte file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", http.DetectContentType(fileByte))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))

	http.ServeFile(w, r, filePath)
}
