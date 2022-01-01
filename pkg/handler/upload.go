package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	tusd "github.com/tus/tusd/pkg/handler"
	"net/http"
	"strings"
)

type UploadConfig struct {
	TusdConfig tusd.Config
}

type UploadHandler struct {
	*tusd.UnroutedHandler
	chi.Router
}

func NewUploadHandler(r chi.Router, path string, config UploadConfig) *UploadHandler {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Post(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	handler, err := tusd.NewUnroutedHandler(config.TusdConfig)

	if err != nil {
		panic(err)
	}

	router := r.With(handler.Middleware)
	//r.Use(handler.Middleware)
	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
			handler.ServeHTTP(w, r)
		})
	})

	router.Post(path, http.HandlerFunc(handler.PostFile))
	router.Head(path + "{id:[-+a-z0-9]+}", http.HandlerFunc(handler.HeadFile))
	router.Patch(path + "{id:[-+a-z0-9]+}", http.HandlerFunc(handler.PatchFile))
	router.Get(path + "{id:[-+a-z0-9]+}", http.HandlerFunc(handler.GetFile))

	if config.TusdConfig.StoreComposer.UsesTerminater {
		router.Delete(path + "{id:[-+a-z0-9]+}", http.HandlerFunc(handler.DelFile))
	}

	return &UploadHandler{handler, r}
}
