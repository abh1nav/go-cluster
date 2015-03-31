package api

import (
	"fmt"
	"net/http"
)

// defaultHandler returns the standard Service and Status response.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "ok", "service": "go-cluster"}`)
}
