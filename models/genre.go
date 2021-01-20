package models

// Genre describes a genre or category of book
type Genre struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"Name"`
}
