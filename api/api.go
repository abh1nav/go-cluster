package api

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// New creates a new router that will get mounted at the web root.
func New() http.Handler {
	r := web.New()
	r.Get("/", defaultHandler)
	return r
}
