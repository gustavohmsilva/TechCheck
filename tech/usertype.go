package tech

// UserType Represent a Type of User, being for the present moment "Reviewer" or
// "Author"
type UserType struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"Name"`
}
