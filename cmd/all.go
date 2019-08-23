package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

// TODO: ***** Add subcmd flags to all to enable individually named buckets. *****
// TODO: ***** Add subcmd flags to all to enable exclusion of a platform. *****
func init() {
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

	uploadCmd.AddCommand(allSubCmd)
}

var allSubCmd = &cobra.Command{
	Use:   "all 'BUCKET_NAME' FILE",
	Short: "The 'all' subcommand auto-constructs the routes to all 3 platforms to the named bucket using the respective env vars.",
	Long: `The 'all'' subcommand auto-constructs the routes to AWS S3, GCS, and Azure named bucket/container using the configured ` +
		` environment variables. Bucket/container must have same name across all 3 platforms for standard (non-flagged) usage.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		all(ctx, args[0], args[1])
	},
}

func all(ctx context.Context, bucket string, file string) {
	s3(ctx, bucket, file)
	gcp(ctx, bucket, file)
	azure(ctx, bucket, file)
	if localBucketPath != "" {
		local(ctx, localBucketPath, file)
	}
}
