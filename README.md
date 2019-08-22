# `upload`
#### A CLI util for uploading documents to bucket storage on AWS, GCP, and Azure cloud native platforms.
<br>

## Usage
##### uploads `1904logo.png` to **GCS**
`$ upload 'gs://your-cnp-bucket' 1904logo.png`
<br>
or
<br>
`$ upload gcs 'your-cnp-bucket' 1904logo.png`
<br>
<br>
##### uploads `1904logo.png` to **S3**
`$ upload 's3://your-cnp-bucket' 1904logo.png`
<br>
or
<br>
`$ upload s3 'your-cnp-bucket' 1904logo.png`
<br>
or
<br>
`$ upload s3 'your-cnp-bucket' 1904logo.png -r 'us-east-1'`
<br>
<br>
##### uploads `1904logo.png` to **Azure**
`$ upload 'azblob://your-cnp-bucket' 1904logo.png`
<br>
or
<br>
`$ upload azure 'your-cnp-bucket' 1904logo.png`
<br>
<br>
##### uploads `1904logo.png` to a **Local** bucket directory
_Per documentation, this is primarily for testing purposes but can be used as an alternative for some file transfers._
<br>
`$ upload 'file:///path/to/your-local-bucket' 1904logo.png`
<br>
or
<br>
`$ upload local /path/to/your-local-bucket 1904logo.png`
<br>
<br>
##### uploads `1904logo.png` to _**all**_ CNP storage providers
_Requires that the target buckets for a single command have the same name on all platforms._
<br>
_An exception to the above is when including **Local** by providing the `-l` or `--local` flag which should be 
immediately followed by the path to the target directory for local 'upload' only which can have any name even if 
different from the name specified for this `all` subcommand._
<br>

`$ upload all 'your-cnp-bucket' 1904logo.png`
<br>
or w/ **Local** option:
<br>
`$ upload all 'your-cnp-bucket' 1904logo.png -l /path/to/your-local-bucket`
<br>
or w/ *alternate* specified **AWS** region for **S3**:
<br>
`$ upload all 'your-cnp-bucket' 1904logo.png -r 'us-east-1'`
<br>

#### help
`$ upload`
<br>
`$ upload -h`
<br>
`$ upload --help`
<br>
`$ upload help s3`
<br>
`$ upload s3 --help`
<br>
<br>
## Prerequisites
- For uploading to your preferred or various Cloud platforms, see [this section for uploading an 
image](https://gocloud.dev/tutorials/cli-uploader/#uploading-an-image) within the Go CDK tutorial.
  - For AWS, you will need to set your default region (i.e. `us-east-1`), but if it is appended to the url,
  this `upload` util will use the appended region instead (i.e. `'s3://your-cnp-bucket?region=us-east-1'`).

- You will also likely need some additional setup not included in the tutorial such as [installing and 
configuring the `aws` cli util](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) and setting up 
your credentials and other environment variables for your respective platforms.

  - This is necessary because the Go CDK looks for these credentials based on the storage/bucket url you provide. Under the hood,
    the Go CDK makes use of the respective cli utils and configuration for the targeted platform
- **Note:** If you plan on building the binary yourself (included `upload` binary is up-to-date), you will need to 
	have the Go `>=1.11` development environment setup with Go modules enabled.
	- On this note, Pull Requests and new issue tickets are welcome!
## Building
- For convenience, a Makefile has been provided to cut out repetitive steps when tweaking or customizing the util.
This is used primarily to trim down the produced binary size and implement best practices.
Build this util simply by executing `make build` .

## Invoke from anywhere as a CLI Command:
- Copy the produced binary to a `/bin` somewhere on your path (i.e. `~/bin`).
    - **example:** `$ cp upload ~/bin/upload`
    - *Note:* `make deploy` will build the binary and copy it to your `~/bin` directory for you if it exists.
    
## Credit
- *This is taken from the [Go CDK Tutorial](https://gocloud.dev/tutorials/cli-uploader/). This repo has been modified from the tutorial example to streamline the usage and provide additional explanation.*
- Extended to introduce subcommands for each CNP as well as `help` & `-h` command features using the [Cobra CLI framework](https://github.com/spf13/cobra).
