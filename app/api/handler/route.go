package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(handler Handler) *chi.Mux {
	r := chi.NewRouter()
	w := handler.ErrorWrapper

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/health", http.StatusMovedPermanently)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/companies", w(handler.GetCompanies))
		r.Post("/companies", w(handler.StoreCompany))
		r.Get("/companies/{id}", w(handler.GetCompanyByID))
		r.Put("/companies/{id}", w(handler.UpdateCompany))
	})

	return r
}
