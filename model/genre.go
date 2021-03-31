package model

// Genre describes a genre or category of book
type Genre struct {
	ID   int64  `json:"id" db:"Id"`
	Name string `json:"name" db:"Name"`
}

type GenreRequest struct {
	ID     uint64 `json:"id"`
	Like   string `json:"like"`
	Size   uint64 `json:"size"`
	Offset uint64 `json:"offset"`
	Count  uint64 `json:"count"`
}

type GenreArgs struct {
	Genre    *Genre       `json:"genre"`
	Includes GenreRequest `json:"includes"`
}

type GenresArgs struct {
	Genres   []*Genre     `json:"genres"`
	Includes GenreRequest `json:"includes"`
}
