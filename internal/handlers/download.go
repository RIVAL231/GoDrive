package handlers

import(
	"fmt"
	"net/http"
	
	"os"
)

func DownloadFile( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "http method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fileName := r.URL.Query().Get("filename")
    if fileName == ""{
		http.Error(w, "filename is required", http.StatusBadRequest)
		return
	}

	filePath := "./uploads/" + fileName
	_, err := os.Stat(filePath)
	if os.IsNotExist(err){
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
	fmt.Println("File downloaded successfully")
}