package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func init() {
	uploadCmd.AddCommand(gcsSubCmd)
}

var gcsSubCmd = &cobra.Command{
	Use:   "gcs 'BUCKET_NAME' FILE",
	Short: "The gcs subcommand auto-constructs the GCP GCS route to the named bucket using the gcloud config.",
	Long:  `The gcs subcommand auto-constructs the GCP GCS route to the named bucket using the gcloud configuration. `,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		gcp(ctx, args[0], args[1])
	},
}

func gcp(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platform:       "GCS",
		platformPrefix: gcpPrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}
