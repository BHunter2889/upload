package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func init() {
	s3SubCmd.Flags().StringVarP(
		&awsRegion,
		"region",
		"r",
		"",
		"The AWS Region where the named bucket is located. "+
			"Overrides any default region previously set.")

	uploadCmd.AddCommand(s3SubCmd)
}

var s3SubCmd = &cobra.Command{
	Use:   "s3 'BUCKET_NAME' FILE",
	Short: "The s3 subcommand auto-constructs the AWS S3 route to the named bucket using the AWS env vars.",
	Long: `The s3 subcommand auto-constructs the AWS S3 route to the named bucket using the configured ` +
		`(or optionally provided) AWS environment variables.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		s3(ctx, args[0], args[1])
	},
}

func s3(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platform:       "S3",
		platformPrefix: s3Prefix,
		awsRegion:      awsRegion,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}
