# `upload`

#### A CLI util for uploading documents to bucket storage on AWS, GCP, or Azure cloud platforms.

## Usage
##### uploads `1904logo.png` to **GCS**
`$ upload 'gs://your-cnp-bucket' 1904logo.png` 
<br>
<br>
##### uploads `1904logo.png` to **S3**
`$ upload 's3://your-cnp-bucket' 1904logo.png`
<br>
<br>
##### uploads `1904logo.png` to **Azure**
`$ upload 'azblob://your-cnp-bucket' 1904logo.png`
<br>
<br>
## Prerequisites
- For uploading to your preferred or various Cloud platforms, see [this section for uploading an 
image](https://gocloud.dev/tutorials/cli-uploader/#uploading-an-image) within the Go CDK tutorial.
  - For AWS, you will need to set your default region (i.e. `us-east-1`), but if it is appended to the url,
  this `upload` util will use the appended region instead (i.e. `'s3://your-cnp-bucket?region=us-east-1'`).

- **Note:** You will also likely need some additional setup not included in the tutorial such as [installing and 
configuring the `aws` cli util](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) and setting up 
your credentials and other environment variables for your respective platforms.

  - This is necessary because the Go CDK looks for these credentials based on the storage/bucket url you provide. Under the hood,
    the Go CDK makes use of the respective cli utils and configuration for the targeted platform  
## Building
- For convenience, a Makefile has been provided to cut out repetitive steps when tweaking or customizing the util.
This is used primarily to trim down the produced binary size and implement best practices.
Build this util simply by executing `make build` .

## Add as a standard CLI Command:
- Copy the produced binary to a `/bin` somewhere on your path (i.e. `~/bin`).
    - **example:** `$ cp upload ~/bin/upload`
    
## Credit
- *This is taken from the [Go CDK Tutorial](https://gocloud.dev/tutorials/cli-uploader/). This repo has been modified from the tutorial example to streamline the usage and provide additional explanation.*
