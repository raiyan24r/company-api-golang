package handler

import (
	"math/rand"
	"net/http"
)

func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request)  {
	// 1 in 3 chance to simulate an internal error
	if rand.Intn(3) == 0 { // values: 0,1,2
		h.Log.Error("simulated internal error in GetCompanies")
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"random failure"}`))
		return
	}

	// Return list of dummy companies
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`[
		{"id": 1, "name": "Company A"},
		{"id": 2, "name": "Company B"}
	]`))
}