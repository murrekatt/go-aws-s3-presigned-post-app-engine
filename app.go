package app

import (
	"net/http"
)

// Inits the route with handler.
func init() {
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/success", handleSuccess)
}
