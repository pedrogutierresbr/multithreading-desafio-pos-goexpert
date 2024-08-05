package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pedrogutierresbr/multithreading-desafio-pos-goexpert/internal/infra/webserver/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cep/{cep}", handlers.SearchCepHandler)

	http.ListenAndServe(":8000", r)
}
