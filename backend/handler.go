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

	filePath, err := imageConverter.doConvert(&fileUploaded{})
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

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "Error retrieve file: %v", err)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(w, "Error retrieve current directory: %v", err)
		return
	}

	dirLocation := filepath.Join(dir, "files")
	if _, err := os.Stat(dirLocation); os.IsNotExist(err) {
		err := os.Mkdir(dirLocation, 0750)
		if err != nil {
			fmt.Fprintf(w, "Error create directory: %v", err)
		}
	}

	filename := handler.Filename
	fileExtension := filepath.Ext(handler.Filename)
	fileLocation := filepath.Join(dirLocation, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(w, "Error open file here: %v - %s", err, fileLocation)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		fmt.Fprintf(w, "Error write file: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "OK",
		"code":          http.StatusOK,
		"target_file":   targetFile,
		"file_location": fileLocation,
		"file_name":     filename,
		"extension":     fileExtension,
	})
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	err := validateRequest(r)
	if err != nil {
		fmt.Printf("Validation error: %v", err)
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
		return
	}
	defer uploadedFile.Close()

	fileToUpload, err := uploadFile(uploadedFile, handler)
	if err != nil {
		fmt.Printf("Error upload file: %v", err)
		w.WriteHeader(http.StatusBadRequest)
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

	// Set the Content-Disposition header to specify the filename and extension
	// fileName := fileToUpload.fileName // Replace with the desired filename
	fileName := getFileName(file.Name())
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	// Set the Content-Type header based on the file type
	contentType := http.DetectContentType(fileByte)
	w.Header().Set("Content-Type", contentType)

	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	fmt.Println("Response Headers:", w.Header())

	// Write the file content to the response body
	_, err = w.Write(fileByte)
	if err != nil {
		fmt.Printf("Error when resturn response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Sorry, something went wrong",
		})
		return
	}

	// w.Header().Set("Content-Type", http.DetectContentType(fileByte))
	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	// w.WriteHeader(http.StatusOK)
	// w.Write(fileByte)

	// http.ServeFile(w, r, filePath)
}
