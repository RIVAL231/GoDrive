package handlers

import (
	"context"
	"fmt"
	"time"

	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rival231/Go-Drive/internal/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)
var s3Client *s3.Client
func init(){
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-north-1"),
	)
	if err != nil {
		// log.Fatalf("Unable to load AWS config, %v", err)
		fmt.Println("Unable to load AWS config, ", err)
		return
	}
	s3Client = s3.NewFromConfig(cfg)
}
  
func UploadFile(w http.ResponseWriter, r *http.Request) {
	 fileCollection := db.GetCollection("Go-Drive","files")
	username := r.Context().Value("username").(string)
// Use username to associate/filter files
   if r.Method != http.MethodPost {
	   http.Error(w, "http method not allowed", http.StatusMethodNotAllowed)
	   return
   }

   file, header, err := r.FormFile("file")
   if err != nil {
	   http.Error(w, "couldnt read file", http.StatusBadRequest)
	   return
   }
   defer file.Close()
   Key := header.Filename + "_"+ username
   _,err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
	Bucket: aws.String("go-drive-v2"),
	Key:    aws.String(header.Filename + "_"+ username+".pdf"),
	Body:   file,
   })
   if err != nil {
	   http.Error(w, "couldnt upload file"+err.Error(), http.StatusInternalServerError)
	   return
   }
   ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
  fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", "go-drive-v2", "eu-north-1", Key)
  fileCollection.InsertOne(ctx,bson.M{
	"username":username,
	"filename":header.Filename,
	"url":fileURL,
   })
   fmt.Println("File uploaded successfully")
   w.Write([]byte("File uploaded successfully"))
}