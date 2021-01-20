package models

import "time"

// Bank describes a institution created by a author (user) to issue checks for
// users that reviewed his books
type Bank struct {
	ID      int       `json:"id" db:"id"`
	Name    string    `json:"name" db:"Name"`
	Image   string    `json:"image" db:"Image"`
	Since   time.Time `json:"since" db:"Since"`
	Founder *User     `json:"founder" db:"Founder"`
}
