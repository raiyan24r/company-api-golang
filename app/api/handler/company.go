package handler

import (
	"company-api/app/api/handler/request"
	"company-api/app/api/handler/response"
	"company-api/business/database"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request) error {
	companies, err := h.App.DbRepo.GetAllCompanies()
	if err != nil {
		h.App.Log.Error("Failed to fetch companies", zap.Error(err))
		return err
	}

	if companies == nil {
		companies = []database.Company{}
	}

	responseCompanies := make([]response.Company, len(companies))
	for i, c := range companies {
		responseCompanies[i] = response.Company{
			ID:                c.ID,
			Name:              c.Name,
			Description:       c.Description,
			AmountOfEmployees: c.AmountOfEmployees,
			Registered:        c.Registered,
			Type:              c.Type,
		}
	}

	return h.writeResponse(r.Context(), w, http.StatusOK, responseCompanies)
}

func (h *Handler) StoreCompany(w http.ResponseWriter, r *http.Request) error {
	var companyRequest request.Company
	if err := json.NewDecoder(r.Body).Decode(&companyRequest); err != nil {
		h.App.Log.Error("Failed to decode request body", zap.Error(err))
		return err
	}

	if err := h.Validate.Struct(companyRequest); err != nil {
		h.writeValidationErrorResponse(r.Context(), w, err)
		return nil
	}

	dbCompany := database.Company{
		Name:              companyRequest.Name,
		Description:       companyRequest.Description,
		AmountOfEmployees: companyRequest.AmountOfEmployees,
		Registered:        companyRequest.Registered,
		Type:              companyRequest.Type,
	}

	id, err := h.App.DbRepo.CreateCompany(dbCompany)
	if err != nil {
		h.App.Log.Error("Failed to create company", zap.Error(err))
		return err
	}

	return h.writeResponse(r.Context(), w, http.StatusCreated, map[string]interface{}{"id": id, "status": "company created"})
}

func (h *Handler) GetCompanyByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	company, err := h.App.DbRepo.GetCompanyByID(id)
	if err != nil {
		h.App.Log.Error("Failed to fetch company", zap.Error(err))
		return err
	}

	if company == nil {
		h.writeErrorResponse(r.Context(), w, http.StatusNotFound, "company not found")
		return nil
	}

	responseCompany := response.Company{
		ID:                company.ID,
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              company.Type,
	}

	return h.writeResponse(r.Context(), w, http.StatusOK, responseCompany)
}

func (h *Handler) UpdateCompany(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	// Check if company exists
	company, err := h.App.DbRepo.GetCompanyByID(id)
	if err != nil {
		h.App.Log.Error("Failed to fetch company", zap.Error(err))
		return err
	}

	if company == nil {
		h.writeErrorResponse(r.Context(), w, http.StatusNotFound, "company not found")
		return nil
	}

	// Parse request body
	var updateReq request.UpdateCompany
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		h.App.Log.Error("Failed to decode request body", zap.Error(err))
		return err
	}

	if updateReq.Name == nil &&
		updateReq.Description == nil &&
		updateReq.AmountOfEmployees == nil &&
		updateReq.Registered == nil &&
		updateReq.Type == nil {
		_ = h.writeResponse(r.Context(), w, http.StatusBadRequest, ValidationErrorResponse{
			Message: "validation failed",
			Errors: map[string]string{
				"body": "at least one field must be provided for update",
			},
		})
		return nil
	}

	if err := h.Validate.Struct(updateReq); err != nil {
		h.writeValidationErrorResponse(r.Context(), w, err)
		return nil
	}

	// Apply updates (only update fields that are provided)
	if updateReq.Name != nil {
		company.Name = *updateReq.Name
	}
	if updateReq.Description != nil {
		company.Description = *updateReq.Description
	}
	if updateReq.AmountOfEmployees != nil {
		company.AmountOfEmployees = *updateReq.AmountOfEmployees
	}
	if updateReq.Registered != nil {
		company.Registered = *updateReq.Registered
	}
	if updateReq.Type != nil {
		company.Type = *updateReq.Type
	}

	// Update in database
	if err := h.App.DbRepo.UpdateCompany(id, *company); err != nil {
		h.App.Log.Error("Failed to update company", zap.Error(err))
		return err
	}

	responseCompany := response.Company{
		ID:                company.ID,
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              company.Type,
	}

	return h.writeResponse(r.Context(), w, http.StatusOK, responseCompany)
}
