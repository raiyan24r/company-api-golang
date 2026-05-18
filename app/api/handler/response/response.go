package response

type Response struct{}

type Company struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type"`
}