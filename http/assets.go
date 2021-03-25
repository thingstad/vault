package http

import (
	"embed"

	"github.com/ryboe/q"
)

// content is our static web server content.
//go:embed web_ui/*
var content embed.FS

// assetFS is a stub for building Vault without a UI.
func assetFS() embed.FS {
	return content
}

func DebugEmbed() {
	q.Q("--> trying to read")
	d, err := content.ReadDir("web_ui")
	if err != nil {
		panic(err)
	}
	for i, c := range d {
		q.Q("\t%d: %s, %t\n", i, c.Name(), c.IsDir())
	}
}
