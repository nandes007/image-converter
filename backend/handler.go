package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "OK",
		"code":   http.StatusOK,
	})
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	err := validateRequest(r)
	if err != nil {
		fmt.Printf("Validation error: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}

	convertTo := r.FormValue("convert_to")
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Error retrieve file: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
		return
	}
	defer uploadedFile.Close()

	fileToUpload, err := uploadFile(uploadedFile, handler)
	if err != nil {
		fmt.Printf("Error upload file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	imageConverter, err := convertImage(convertTo)
	if err != nil {
		fmt.Printf("Error when passed convert image: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	filePath, err := imageConverter.doConvert(fileToUpload)
	if err != nil {
		fmt.Printf("Error convert file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
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
		w.Header().Set("Content-Type", "application/json")
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	fileName := getFileName(file.Name())
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	contentType := http.DetectContentType(fileByte)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	_, err = w.Write(fileByte)
	if err != nil {
		fmt.Printf("Error when resturn response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}
}
