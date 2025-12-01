package response

type Response struct{}


type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered bool   `json:"registered"`
	Type       string `json:"type"`
}