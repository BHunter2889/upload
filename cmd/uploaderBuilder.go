package cmd // import "github.com/BHunter2889/upload/cmd"

import (
	"context"
	"fmt"
	"log"
)

// Provides a non-exported builder to simplify and streamline the upload process for each/all platforms.
type uploaderBuilder struct {
	bucketName     string
	fileName       string
	platform       string
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
	log.Printf("Preparing for %s upload... \n", u.platform)
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
	log.Printf("Uploading %s to %s %s...", u.fileName, u.platform, u.bucketName)
	if err := upload(ctx, u.bucketUrl, u.fileName); err != nil {
		log.Printf("Failed to upload to %s with error: %v", u.platform, err)
	} else {
		log.Print(" Done.")
	}
	wg.Done()
}

func (u *uploaderBuilder) buildAndUpload(ctx context.Context) {
	u.buildBucketUrl().upload(ctx)
}
