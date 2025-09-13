package main

import (
	// "fmt"
	// "log"
	"fmt"
	"log"
	"net/http"

	"github.com/rival231/Go-Drive/internal/handlers"
)

func main(){
	http.HandleFunc("/health",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("the server is healthy"))
	})
	http.HandleFunc("/upload",handlers.UploadFile)
	http.HandleFunc("/download",handlers.DownloadFile)
	http.HandleFunc("/list",handlers.ListFiles)

    fmt.Print("The server is running on port 8000")
	err := http.ListenAndServe(":8000",nil)
	if err!=nil{
	log.Fatal("The server couldnt start")
	return
	}
	
}