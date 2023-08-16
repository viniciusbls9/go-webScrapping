package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func handlerSaveProducts(w http.ResponseWriter, r *http.Request, content []byte, productName string) {
	filePath := "products/"+ productName +".json"

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create directory: %v", err))
		return
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to write JSON file: %v", err))
		return
	}
}