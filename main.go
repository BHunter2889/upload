// Command upload saves files to blob storage on GCP, AWS, and Azure.
package main

import (
	"context"
	"github.com/spf13/cobra"
	"gocloud.dev/blob"
	// Import blank GoCDK blob impls for desired platform.
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"io/ioutil"
	"log"
	"os"
)

var ctx context.Context

func main() {
	ctx = context.Background();
	uploadCmd := &cobra.Command{
		Use:   "upload BUCK",
		Short: "upload is a util to upload to any cloud platform storage bucket provider",
		Long:  `A convenient cli util to upload objects to bucket storage in AWS, GCP, or Azure. 
				Written by BHunter2889 in Go, based off the Go CDK upload tutorial.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			upload(ctx, args[0], args[1])
		},
	}

	err := uploadCmd.Execute()
	if err != nil {
		log.Fatal("upload command failed to execute: ", err)
	}

	//	Define Input
	if len(os.Args) != 3 {
		log.Fatal("usage: upload BUCKET_URL FILE")
	}

}

func upload(ctx context.Context, bucketURL string, file string) {
	//	Open Bucket Connection
	bucket, err := blob.OpenBucket(ctx, bucketURL)
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
	writer, err := bucket.NewWriter(ctx, file, nil)
	if err != nil {
		log.Fatalf("Failed to obtain writer: %s", err)
	}

	_, err = writer.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to bucket: %s", err)
	}
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to close: %s", err)
	}
	//	Or alternatively use this shortcut at the expense of explicit error handling:
	//	err = bucket.WriteAll(ctx, file, data, nil);
	//if err != nil {
	//	log.Fatalf("Writer Failed: ", err)
	//}
}
