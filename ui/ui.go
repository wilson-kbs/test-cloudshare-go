package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist
var embedUI embed.FS

func GetFiles() fs.FS {
	f, err := fs.Sub(embedUI, "dist")
	if err != nil {
		panic(err)
	}
	return f
}
