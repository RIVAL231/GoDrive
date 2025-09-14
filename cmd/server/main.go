package main

import (
	// "fmt"
	// "log"
	"fmt"
	"log"
	"net/http"

	"github.com/rival231/Go-Drive/internal/db"
	"github.com/rival231/Go-Drive/internal/handlers"
)

func main(){
    db.ConnectDB("mongodb+srv://admin:admin@cluster0.xatwqkk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	http.HandleFunc("/health",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("the server is healthy"))
	})

	http.HandleFunc("/users", handlers.CreateUser)
	http.HandleFunc("/login",handlers.UserAuthMiddleware(http.HandlerFunc(handlers.UserLogin)).ServeHTTP)
	http.HandleFunc("/upload", handlers.JWTAuthMiddleware(http.HandlerFunc(handlers.UploadFile)).ServeHTTP)
    http.HandleFunc("/download", handlers.JWTAuthMiddleware(http.HandlerFunc(handlers.DownloadFile)).ServeHTTP)
    http.HandleFunc("/list", handlers.JWTAuthMiddleware(http.HandlerFunc(handlers.ListFiles)).ServeHTTP)

    fmt.Print("The server is running on port 8000")
	err := http.ListenAndServe(":8000",nil)
	if err!=nil{
	log.Fatal("The server couldnt start")
	return
	}
	
}