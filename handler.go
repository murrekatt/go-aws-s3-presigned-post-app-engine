package app

import (
	"html/template"
	"net/http"

	"appengine"

	"github.com/twinj/uuid"
	"github.com/murrekatt/go-s3presigned-post"
)

const (
	regionName = "FILL-ME-IN"
	bucketName = "FILL-ME-IN"
	accessKeyID = "FILL-ME-IN"
	secretAccessKey = "FILL-ME-IN"
)

// HTML form for pre-signed POST
const htmlDocument = `
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
</head>
<body>
  <h1>AWS S3 File Uploader</h1>
  <form action="{{.Action}}" method="post" enctype="multipart/form-data">
    <input type="hidden" name="key" value="{{.Key}}" />
    <input type="hidden" name="acl" value="public-read" />
    <input type="hidden" name="X-Amz-Credential" value="{{.Credential}}" />
    <input type="hidden" name="X-Amz-Algorithm" value="AWS4-HMAC-SHA256" />
    <input type="hidden" name="X-Amz-Date" value="{{.Date}}" />
    <h3>Tags</h3>
    <input type="hidden" name="Policy" value="{{.Policy}}" />
    <input type="hidden" name="X-Amz-Signature" value="{{.Signature}}" />
    <h3>File</h3>
    <input type="file"   name="file" />
    <!-- The elements after this will be ignored -->
    <input type="submit" name="submit" value="Upload to Amazon S3" />
  </form>
</body>
</html>
`

// Handles upload requests.
func handleUpload(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// parse HTML template
	t, err := template.New("presign").Parse(htmlDocument)
	if err != nil {
		c.Errorf("Error parsing template: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// a unique key to upload
	id := uuid.NewV4().String()

	// AWS S3 credentials
	creds := &s3.Credentials{
		Region: regionName,
		Bucket: bucketName,
		AccessKeyID: accessKeyID,
		SecretAccessKey: secretAccessKey,
	}

	// create pre-signed POST details
	post, err := s3.NewPresignedPOST(id, creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// render HTML form
	t.Execute(w, post)
}
