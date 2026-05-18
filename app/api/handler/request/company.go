package request

type Company struct {
	Name              string `json:"name" validate:"required,min=2,max=50"`
	Description       string `json:"description" validate:"omitempty,max=3000"`
	AmountOfEmployees int    `json:"amount_of_employees" validate:"required,gte=1"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type" validate:"required,min=2,max=300"`
}

type UpdateCompany struct {
	Name              *string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Description       *string `json:"description,omitempty" validate:"omitempty,max=3000"`
	AmountOfEmployees *int    `json:"amount_of_employees,omitempty" validate:"omitempty,gte=1"`
	Registered        *bool   `json:"registered,omitempty"`
	Type              *string `json:"type,omitempty" validate:"omitempty,min=2,max=300"`
}