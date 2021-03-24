package model

// UserType Represent a Type of User, being for the present moment "Reviewer" or
// "Author"
type UserType struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"Name"`
}

type UserTypeRequest struct {
	ID     uint64 `json:"id"`
	Like   string `json:"like"`
	Size   uint64 `json:"size"`
	Offset uint64 `json:"offset"`
	Count  uint64 `json:"count"` // TODO: REMEMBER TO IMPLEMENT A SIMPLE COUNT TO AID IN THE PAGINATION
}

type UserTypeArgs struct {
	UserType *UserType       `json:"user_type"`
	Includes UserTypeRequest `json:"includes"`
}

type UserTypesArgs struct {
	UserTypes []*UserType     `json:"user_types"`
	Includes  UserTypeRequest `json:"includes"`
}
