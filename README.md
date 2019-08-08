# `upload`
*This is taken from the [Go CDK Tutorial](https://gocloud.dev/tutorials/cli-uploader/). This repo has been modified from the tutorial example to streamline the usage and provide additional explanation.*

#### This is a CLI util for uploading documents to bucket storage on AWS, GCP, or Azure.

## Usage
#### uploads `1904logo.png` to GCS
$ upload 'gs://go-cloud-bucket' 1904logo.png

#### uploads `1904logo.png` to S3
$ upload 's3://go-cloud-bucket' 1904logo.png

#### uploads `1904logo.png` to Azure
$ upload 'azblob://go-cloud-bucket' 1904logo.png

## Prerequisites
For uploading to your preferred or various Cloud platforms, see [this section for uploading an image](https://gocloud.dev/tutorials/cli-uploader/#uploading-an-image) within the Go CDK tutorial.

**Note:** You will also likely need some additional setup not included in the tutorial such as [installing and configuring the `aws` cli util](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) and setting up your credentials and other environment variables for your respective platforms.

The above is necessary because the Go CDK looks for these credentials based on the storage/bucket url you provide.

## Building

## Add as a standard CLI Command: