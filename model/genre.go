package model

// Genre describes a genre or category of book
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
