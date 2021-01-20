package models

// User is a single user
type User struct {
	ID         int       `json:"id" db:"ID"`
	Username   string    `json:"username" db:"Username"`
	Passphrase string    `json:"passphrase" db:"Passphrase"`
	FullName   string    `json:"full_name" db:"FullName"`
	Type       *UserType `json:"user_type" db:"UserType"`
	Bio        string    `json:"bio" db:"Bio"`
	Image      string    `json:"image" db:"Image"`
}
