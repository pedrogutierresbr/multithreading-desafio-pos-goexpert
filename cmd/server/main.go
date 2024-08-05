package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pedrogutierresbr/multithreading-desafio-pos-goexpert/internal/infra/webserver/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cep/{cep}", handlers.SearchCepHandler)

	http.ListenAndServe(":8000", r)
}
