package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func init() {
	uploadCmd.AddCommand(localSubCmd)
}

var localSubCmd = &cobra.Command{
	Use:   "local PATH_TO_BUCKET FILE",
	Short: "The local subcommand locally 'uploads' FILE to PATH_TO_BUCKET.",
	Long: `The local subcommand 'uploads' FILE to the local filesystem path defined by PATH_TO_BUCKET. ` +
		`This is mainly useful for testing purposes but could be used in place of the 'cp' command (for some file types), for example.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		local(ctx, args[0], args[1])
	},
}

func local(ctx context.Context, bucket string, file string) {
	uploader := uploaderBuilder{
		platform:       "Local",
		platformPrefix: localPrefix,
		bucketName:     bucket,
		fileName:       file,
	}
	uploader.buildAndUpload(ctx)
}
