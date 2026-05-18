package request

type Company struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	AmountOfEmployees int  `json:"amount_of_employees"`
	Registered      bool   `json:"registered"`
	Type            string `json:"type"`
}
type UpdateCompany struct {
	Name              *string `json:"name,omitempty"`
	Description       *string `json:"description,omitempty"`
	AmountOfEmployees *int    `json:"amount_of_employees,omitempty"`
	Registered        *bool   `json:"registered,omitempty"`
	Type              *string `json:"type,omitempty"`
}