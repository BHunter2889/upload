// Command upload saves files to blob storage on GCP, AWS, and Azure.
package main

import (
	uploadCmd "github.com/BHunter2889/upload/cmd"
)

// Execute root command on startup.
func main() {
	uploadCmd.Execute()
}
