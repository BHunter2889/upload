package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func init() {
	uploadCmd.AddCommand(azSubCmd)
}

var azSubCmd = &cobra.Command{
	Use:   "azure 'BUCKET_NAME' FILE",
	Short: "The azure subcommand auto-constructs the Azure route to the named container using the configured azure env vars.",
	Long: `The azure subcommand auto-constructs the Azure route to the named container using the ` +
		`AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_KEY environment variable configuration.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		azure(ctx, args[0], args[1])
	},
}

func azure(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platform:       "Azure",
		platformPrefix: azurePrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}