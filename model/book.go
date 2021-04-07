package model

// Book  is a single book
type Book struct {
	ID                 uint64 `json:"id" db:"id"`
	Author             *User  `json:"author" db:"Author"`
	ISBN               string `json:"isbn" db:"ISBN"`
	Name               string `json:"name" db:"Name"`
	Edition            uint64 `json:"edition" db:"Edition"`
	LatestEdition      bool   `json:"latest_edition" db:"LatestEdition"`
	PredecessorEdition *Book  `json:"predecessor_edition" db:"PredecessorEdition"`
	Image              string `json:"image" db:"Image"`
}

type BookRequest struct {
	ID     uint64 `json:"id"`
	Author string `json:"author"`
	ISBN   uint64 `json:"isbn"`
	Name   string `json:"name"`
	Like   string `json:"like"`
	Size   uint64 `json:"size"`
	Offset uint64 `json:"offset"`
	Count  uint64 `json:"count"`
}

type BookArgs struct {
	Book     *Book
	Includes BookRequest
}

type BooksArgs struct {
	Books    []*Book
	Includes BookRequest
}
