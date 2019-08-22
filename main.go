// Command upload saves files to blob storage on GCP, AWS, and Azure.
package main

import (
	uploadCmd "github.com/BHunter2889/upload/cmd"
)

func main() {

	// TODO: ***** Add subcmd flags to all to enable individually named buckets. *****
	// TODO: ***** Add subcmd flags to all to enable exclusion of a platform. *****
	// TODO: ***** CONSIDER: Add flags to root upload to enable platform specification and individual naming. *****
	uploadCmd.Execute()
}