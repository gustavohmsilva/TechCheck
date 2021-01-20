package models

// Book  is a single book
type Book struct {
	ID                 int    `json:"id" db:"id"`
	Author             *User  `json:"author" db:"Author"`
	ISBN               string `json:"isbn" db:"ISBN"`
	Name               string `json:"name" db:"Name"`
	Edition            int    `json:"edition" db:"Edition"`
	LatestEdition      bool   `json:"latest_edition" db:"LatestEdition"`
	PredecessorEdition *Book  `json:"predecessor_edition" db:"PredecessorEdition"`
	Image              string `json:"image" db:"Image"`
}
