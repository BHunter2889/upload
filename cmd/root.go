package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"gocloud.dev/blob"
	// Import blank GoCDK blob impls for desired platform.
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"io/ioutil"
	"log"
	"sync"
)

const (
	urlTemplate = "%s://%s"
	s3Prefix    = "s3"
	gcpPrefix   = "gs"
	azurePrefix = "azblob"
	localPrefix = "file"
)

var (
	ctx             context.Context
	awsRegion       string
	localBucketPath string
	wg              sync.WaitGroup
)

// TODO: ***** CONSIDER: Add flags to root upload to enable platform specification and individual naming. *****
func init() {
	ctx = context.Background()
}

func Execute() {
	if err := uploadCmd.Execute(); err != nil {
		log.Fatalf("Could not execute 'upload' command: %s", err)
	}
}

var uploadCmd = &cobra.Command{
	Use:   "upload 'BUCKET_URL' FILE",
	Short: "upload is a util to upload to any cloud platform storage bucket provider",
	Long: `A convenient cli util to upload objects to bucket storage in AWS, GCP, or Azure. ` +
		`Written by bhunter2889 in Go, based off the Go CDK upload tutorial.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := upload(ctx, args[0], args[1]); err != nil {
			log.Fatalf("Upload to %s failed with error: %s", args[0], err)
		}
	},
}

func upload(ctx context.Context, bucketURL string, file string) (error error) {
	//	Open Bucket Connection
	bucket, err := blob.OpenBucket(ctx, bucketURL)
	if err != nil {
		error = err
		log.Printf("Failed to setup bucket: %s", err)
		return
	}
	defer bucket.Close()

	//	Prepare File for upload.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		error = err
		log.Printf("Failed to read file: %s", err)
		return
	}

	// To write the file to the bucket, you can do this:
	writer, err := bucket.NewWriter(ctx, file, nil)
	if err != nil {
		error = err
		log.Printf("Failed to obtain writer: %s", err)
		return
	}

	_, err = writer.Write(data)
	if err != nil {
		error = err
		log.Printf("Failed to write to bucket: %s", err)
		return
	}
	if err := writer.Close(); err != nil {
		error = err
		log.Printf("Failed to close: %s", err)
		return
	}

	//	Or alternatively use this shortcut at the expense of explicit error handling:
	//	err = bucket.WriteAll(ctx, file, data, nil);
	//if err != nil {
	//	log.Fatalf("Writer Failed: ", err)
	//}
	return
}
