// +build ui

package http

import (
	"embed"
	"io/fs"
)

// content is our static web server content.
//go:embed web_ui
var content embed.FS

// assetFS is a stub for building Vault without a UI.
func assetFS() fs.FS {
	return content
}
