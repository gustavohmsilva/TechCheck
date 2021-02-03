package model

// ResponseError is only used to return a logic error to the client
type ResponseError struct {
	Err string `json:"reason"`
}
