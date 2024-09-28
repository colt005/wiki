//go:build !dev
// +build !dev

package server

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	return http.FileServerFS(publicFS)
}
