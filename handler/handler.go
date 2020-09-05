package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello %s", chi.URLParam(r, "name"))
}
