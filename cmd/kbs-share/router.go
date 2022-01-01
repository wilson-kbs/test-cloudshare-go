package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/wilson-kbs/test-cloudshare-go/pkg/handler"
	"github.com/wilson-kbs/test-cloudshare-go/pkg/web"
	"github.com/wilson-kbs/test-cloudshare-go/ui"
	"net/http"
	"strings"
)

type RouterConfig struct {
	Base              string
	TusdStoreComposer *tusd.StoreComposer
}

func NewRouter(config RouterConfig) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	web.ParseIndexFile(web.Config{
		Base: config.Base,
	})

	if config.Base != "" && config.Base != "/" {
		r.Route(config.Base, func(r chi.Router) {
			generateRoutes(r, config)
		})
	} else {
		r.Route("/", func(r chi.Router) {
			generateRoutes(r, config)
		})
	}

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return r
}

func generateRoutes(r chi.Router, config RouterConfig) {

	r.Get("/", indexWebHandlerFunc)

	handler.FileServer(r, "/", ui.GetFiles())

	uploadServer(r, config)
	r.NotFound(indexWebHandlerFunc)
}

func uploadServer(r chi.Router, config RouterConfig) {
	basePath := config.Base + "files"
	tusdConfig := tusd.Config{
		BasePath:              basePath,
		StoreComposer:         config.TusdStoreComposer,
		NotifyCompleteUploads: true,
	}
	uploadConfig := handler.UploadConfig{TusdConfig: tusdConfig}

	uploadHandler := handler.NewUploadHandler(r, basePath, uploadConfig)

	go func() {
		for {
			event := <-uploadHandler.CompleteUploads
			fmt.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()
}

func indexWebHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	_, err := w.Write(web.GetIndex())
	if err != nil {
		panic(err)
		return
	}
}
