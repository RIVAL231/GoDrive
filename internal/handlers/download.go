package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
    presignClient := s3.NewPresignClient(s3Client)
	presignedUrl, err := presignClient.PresignGetObject(context.Background(),
  &s3.GetObjectInput{
   Bucket: aws.String("go-drive-v2"),
   Key:    aws.String(fileName),
  },
  s3.WithPresignExpires(time.Minute*15))
  if err != nil {
	  http.Error(w, "couldnt generate presigned url"+err.Error(), http.StatusInternalServerError)
	  return
  }

	http.Redirect(w, r, presignedUrl.URL, http.StatusFound)
	fmt.Println("File downloaded successfully")
}