//go:build dev
// +build dev

package server

import (
	"fmt"
	"net/http"
	"os"

)

func public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("./server/public")))
}
