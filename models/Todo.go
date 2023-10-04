package models

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	// Add more fields as needed for your user model
}
