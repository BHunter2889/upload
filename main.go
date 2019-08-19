// Command upload saves files to blob storage on GCP, AWS, and Azure.
package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"gocloud.dev/blob"
	// Import blank GoCDK blob impls for desired platform.
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"io/ioutil"
	"log"
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
)

func main() {
	ctx = context.Background()

	uploadCmd := &cobra.Command{
		Use:   "upload 'BUCKET_URL' FILE",
		Short: "upload is a util to upload to any cloud platform storage bucket provider",
		Long: `A convenient cli util to upload objects to bucket storage in AWS, GCP, or Azure. ` +
			`Written by bhunter2889 in Go, based off the Go CDK upload tutorial.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			upload(ctx, args[0], args[1])
		},
	}

	s3SubCmd := &cobra.Command{
		Use:   "s3 'BUCKET_NAME' FILE",
		Short: "The s3 subcommand auto-constructs the AWS S3 route to the named bucket using the AWS env vars.",
		Long: `The s3 subcommand auto-constructs the AWS S3 route to the named bucket using the configured ` +
			`(or optionally provided) AWS environment variables.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			s3(ctx, args[0], args[1])
		},
	}

	s3SubCmd.Flags().StringVarP(
		&awsRegion,
		"region",
		"r",
		"",
		"The AWS Region where the named bucket is located. "+
			"Overrides any default region previously set.")

	gcsSubCmd := &cobra.Command{
		Use:   "gcs 'BUCKET_NAME' FILE",
		Short: "The gcs subcommand auto-constructs the GCP GCS route to the named bucket using the gcloud config.",
		Long:  `The gcs subcommand auto-constructs the GCP GCS route to the named bucket using the gcloud configuration. `,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			gcp(ctx, args[0], args[1])
		},
	}

	azSubCmd := &cobra.Command{
		Use:   "azure 'BUCKET_NAME' FILE",
		Short: "The azure subcommand auto-constructs the Azure route to the named container using the configured azure env vars.",
		Long: `The azure subcommand auto-constructs the Azure route to the named container using the ` +
			`AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_KEY environment variable configuration.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			azure(ctx, args[0], args[1])
		},
	}

	localSubCmd := &cobra.Command{
		Use:   "local PATH_TO_BUCKET FILE",
		Short: "The local subcommand locally 'uploads' FILE to PATH_TO_BUCKET.",
		Long: `The local subcommand 'uploads' FILE to the local filesystem path defined by PATH_TO_BUCKET. ` +
			`This is mainly useful for testing purposes but could be used in place of the 'cp' command, for example.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			local(ctx, args[0], args[1])
		},
	}

	allSubCmd := &cobra.Command{
		Use:   "all 'BUCKET_NAME' FILE",
		Short: "The 'all' subcommand auto-constructs the routes to all 3 platforms to the named bucket using the respective env vars.",
		Long: `The 'all'' subcommand auto-constructs the routes to AWS S3, GCS, and Azure named bucket/container using the configured ` +
			` environment variables. Bucket/container must have same name across all 3 platforms for standard (non-flagged) usage.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			all(ctx, args[0], args[1])
		},
	}

	allSubCmd.Flags().StringVarP(
		&awsRegion,
		"region",
		"r",
		"",
		"The AWS Region where the named bucket is located. "+
			"Overrides any default region previously set.")

	allSubCmd.Flags().StringVarP(
		&localBucketPath,
		"local",
		"l",
		"",
		"Also upload locally. The provided argument is the `PATH_TO_BUCKET` locally. "+
			"See `upload local --help` for the same usage.")

	// TODO: ***** Add subcmd flags to all to enable individually named buckets. *****
	// TODO: ***** Add subcmd flags to all to enable exclusion of a platform. *****
	// TODO: ***** CONSIDER: Add flags to root upload to enable platform specification and individual naming. *****

	uploadCmd.AddCommand(s3SubCmd, gcsSubCmd, azSubCmd, localSubCmd)
	err := uploadCmd.Execute()
	if err != nil {
		log.Fatal("upload command failed to execute: ", err)
	}
}

func s3(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platformPrefix: s3Prefix,
		awsRegion:      awsRegion,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}

func gcp(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platformPrefix: gcpPrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}

func azure(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platformPrefix: azurePrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}

func local(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platformPrefix: localPrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}

// TODO: Verify safety...
func all(ctx context.Context, bucket string, file string) {
	go s3(ctx, bucket, file)
	go gcp(ctx, bucket, file)
	go azure(ctx, bucket, file)
	if localBucketPath != "" {
		go local(ctx, bucket, file)
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

type uploaderBuilder struct {
	bucketName     string
	fileName       string
	platformPrefix string
	bucketUrl      string
	awsRegion      string
}

func (u *uploaderBuilder) toBucket(bucketName string) *uploaderBuilder {
	u.bucketName = bucketName
	return u
}

func (u *uploaderBuilder) withFile(fileName string) *uploaderBuilder {
	u.fileName = fileName
	return u
}

func (u *uploaderBuilder) usingPlatform(bucketUrlPrefix string) *uploaderBuilder {
	u.platformPrefix = bucketUrlPrefix
	return u
}

func (u *uploaderBuilder) inAWSRegion(region string) *uploaderBuilder {
	u.awsRegion = region
	return u
}

func (u *uploaderBuilder) buildBucketUrl() *uploaderBuilder {
	if u.awsRegion != "" {
		regionQuery := fmt.Sprintf("?region=%s", awsRegion)
		u.bucketUrl = fmt.Sprintf(urlTemplate, s3Prefix, u.bucketName+regionQuery)
	} else if u.platformPrefix == localPrefix {
		u.bucketUrl = fmt.Sprintf(urlTemplate, localPrefix, "/"+u.bucketName)
	} else {
		u.bucketUrl = fmt.Sprintf(urlTemplate, s3Prefix, u.bucketName)
	}
	return u
}

func (u *uploaderBuilder) upload(ctx context.Context) {
	upload(ctx, u.bucketUrl, u.fileName)
}

func (u *uploaderBuilder) buildAndUpload(ctx context.Context) {
	u.buildBucketUrl().upload(ctx)
}
