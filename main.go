// Command upload saves files to blob storage on GCP, AWS, and Azure.
package main

import (
	"context"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
	"io/ioutil"
	"log"
	"os"
)

func main() {
//	Define Input
	if len(os.Args) != 3 {
		log.Fatal("usage: upload BUCKET_URL FILE")
	}
	bucketURL := os.Args[1]
	file := os.Args[2]
	_, _ = bucketURL, file
	ctx := context.Background()

//	Open Bucket Connection
	bucket, err := blob.OpenBucket(context.Background(), bucketURL)
	if err != nil {
		log.Fatalf("Failed to setup bucket: %s", err)
	}
	defer bucket.Close()

//	Prepare File for upload.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

// To write the file to the bucket, you can do this:
	//writer, err := bucket.NewWriter(ctx, file, nil)
	//if err != nil {
	//	log.Fatalf("Failed to obtain writer: %s", err)
	//}
	//
	//_, err = writer.Write(data)
	//if err != nil {
	//	log.Fatalf("Failed to write to bucket: %s", err)
	//}
	//if err := writer.Close(); err != nil {
	//	log.Fatalf("Failed to close: %s", err)
	//}

//	Or simply use this shortcut:
	err = bucket.WriteAll(ctx, file, data, nil);
	if err != nil {
		log.Fatalf("Writer Failed: ", err)
	}
}
