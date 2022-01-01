package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := NewStore("s3")

	router := NewRouter(RouterConfig{
		Base:              "/",
		TusdStoreComposer: store,
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}
}
