package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"test-project/handler"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{name}", handler.Hello)
	log.Panic(http.ListenAndServe(":80", r))
}
