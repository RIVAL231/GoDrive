package handlers

import (
	"context"
	"fmt"

	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "http method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//Commented old code that listed files from local uploads directory
	// files, err := os.ReadDir("./uploads")
	// if err != nil {
	// 	http.Error(w, "couldnt read uploads directory", http.StatusInternalServerError)
	// 	return
	// }
	// for _, file := range files {
	// 	fmt.Fprintln(w, file.Name())
	// }
    
    //New code to list files from S3 bucket
	 output, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
        Bucket: aws.String("go-drive-v2"),
		
    })
     if err!=nil{
		http.Error(w, "couldnt list files from s3 bucket"+err.Error(), http.StatusInternalServerError)
		return
	 }
	 for i, object := range output.Contents {
        // log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
		w.Write([]byte(fmt.Sprintf("%d: %s (%d bytes)\n", i+1, aws.ToString(object.Key), object.Size)))
    }


}