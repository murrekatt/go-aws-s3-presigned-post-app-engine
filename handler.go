package app

import (
	"html/template"
	"net/http"

	"appengine"

	"github.com/twinj/uuid"
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
    <input type="input"  name="x-amz-meta-tag" value="" size="50" /><br />
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
	// create pre-signed POST details
	policy := NewPolicy(id, regionName)
	post, err := NewPresignedPOST(policy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// render HTML form
	t.Execute(w, post)
}
