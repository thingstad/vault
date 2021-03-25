package http

import (
	"embed"
	"io/fs"
	"net/http"
)

// content is our static web server content.
//go:embed web_ui/*
var content embed.FS

// assetFS is a stub for building Vault without a UI.
func assetFS() http.FileSystem {
	// sub to web_ui
	f, err := fs.Sub(content, "web_ui")
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}
