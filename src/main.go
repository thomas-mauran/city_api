package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomas-mauran/city_api/src/utils"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/_health", func(w http.ResponseWriter, r *http.Request) {
		utils.Response(w, r, 204, "welcome")
		return
	})
	http.ListenAndServe(":3000", r)
}