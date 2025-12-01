package handler

import (
	"company-api/app/api/handler/response"
	"errors"
	"math/rand"
	"net/http"
)

func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request) error {
	if rand.Intn(3) == 0 {
		return errors.New("simulated internal error")
	}

	companies := []response.Company{
		{ID: 1, Name: "Tech Innovators Inc.", Description: "Leading tech company", AmountOfEmployees: 500, Registered: true, Type: "Technology"},
		{ID: 2, Name: "Green Energy Solutions", Description: "Renewable energy provider", AmountOfEmployees: 200, Registered: true, Type: "Energy"},
		{ID: 3, Name: "HealthPlus Corp.", Description: "Healthcare services", AmountOfEmployees: 300, Registered: true, Type: "Healthcare"},
	}

	return h.writeResponse(r.Context(), w, http.StatusOK, companies)
}
