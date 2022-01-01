package handler

import (
	"github.com/go-chi/chi/v5"
	"io/fs"
	"net/http"
	"strings"
)

func FileServer(r chi.Router, path string, root fs.FS) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	files, err := fs.ReadDir(root, ".")
	if err != nil {
		panic(err)
	}

	for _, v := range files {
		if v.IsDir() {
			dir, err := fs.Sub(root, v.Name())
			if err != nil {
				panic(err)
			}
			dirPath := path + v.Name() + "/*"
			r.Get(dirPath, func(w http.ResponseWriter, r *http.Request) {
				rctx := chi.RouteContext(r.Context())
				pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
				handler := http.StripPrefix(pathPrefix, http.FileServer(http.FS(dir)))
				handler.ServeHTTP(w, r)
			})
		} else {
			filePath := path + v.Name()
			r.Get(filePath, func(w http.ResponseWriter, r *http.Request) {
				rctx := chi.RouteContext(r.Context())
				suffix := "/" + v.Name()
				pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), suffix)
				handler := http.StripPrefix(pathPrefix, http.FileServer(http.FS(root)))
				handler.ServeHTTP(w, r)
			})
		}
	}
}
