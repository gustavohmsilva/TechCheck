package tech

import "time"

// BookError contains a single error that was informed by a user about a certain
// book
type BookError struct {
	ID          int       `json:"id" db:"id"`
	Book        *Book     `json:"book" db:"Book"`
	Page        int       `json:"page" db:"Page"`
	Chapter     int       `json:"chapter" db:"Chapter"`
	Paragraph   int       `json:"paragraph" db:"Paragraph"`
	Line        int       `json:"line" db:"Line"`
	Description string    `json:"description" db:"Description"`
	ReportedBy  *User     `json:"reported_by" db:"ReportedBy"`
	ReportedAt  time.Time `json:"reported_at" db:"ReportedAt"`
	Accepted    bool      `json:"accepted" db:"Accepted"`
	AcceptedAt  time.Time `json:"accepted_at" db:"AcceptedAt"`
	Severity    int       `json:"severity" db:"Severity"`
	CheckIssued bool      `json:"check_issued" db:"CheckIssued"`
}
