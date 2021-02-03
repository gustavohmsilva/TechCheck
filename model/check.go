package model

import "time"

// Check is issued when a author approve a revision given by any user.
type Check struct {
	ID          int        `json:"id" db:"id"`
	Bank        *Bank      `json:"bank" db:"Bank"`
	IssueDate   time.Time  `json:"issue_date" db:"IssueDate"`
	Beneficiary *User      `json:"beneficiary" db:"Beneficiary"`
	Value       int        `json:"value" db:"Value"`
	Issuer      *User      `json:"issuer" db:"Issuer"`
	OverError   *BookError `json:"over_error" db:"OverError"`
}
