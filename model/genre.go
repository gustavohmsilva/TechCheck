package model

// Genre describes a genre or category of book
type Genre struct {
	ID   int64  `json:"id" db:"Id"`
	Name string `json:"name" db:"Name"`
}

type GenreRequest struct {
	Like   string `json:"like"`
	Size   uint64 `json:"size"`
	Offset uint64 `json:"offset"`
}

type GenreArgs struct {
	Genre
	Request GenreRequest
}
