package handlers

import (
	"fmt"
	"net/http"
	"os"

)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "http method not allowed", http.StatusMethodNotAllowed)
		return
	}
	files, err := os.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "couldnt read uploads directory", http.StatusInternalServerError)
		return
	}
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
	
}