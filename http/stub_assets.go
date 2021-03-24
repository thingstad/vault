// +build !ui

package http

import "io/fs"

func init() {
	uiBuiltIn = false
}

// assetFS is a stub for building Vault without a UI.
func assetFS() fs.FS {
	return nil
}
