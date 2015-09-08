# go-aws-s3-presigned-post-app-engine

Example showing AWS S3 pre-signed POST.

## Run locally

1. You'll need to have the Go App Engine SDK installed and working.
2. Clone this repo.
3. Create an AWS S3 test bucket with public-read

    {
      "Version": "2012-10-17",
      "Id": "Policy1234567890",
      "Statement": [{
        "Sid": "Stmt1234567890",
        "Effect": "Allow",
        "Principal": {
          "AWS": "*"
        },
        "Action": "s3:GetObject",
        "Resource": "arn:aws:s3:::bucketnamehere/*"
      }]
    }

4. Configure the AWS credentials in `s3.go`

    const regionName = "TODO"
    const bucketName = "TODO"
    const accessKeyID = "TODO"
    const secretAccessKey = "TODO"

5. run the app locally

    goapp serve

6. Open `localhost:8080` with your browser.

You should now see the HTML form so you can select a file to upload. Assuming
all setup is correct, the file will be uploaded to the S3 bucket.
