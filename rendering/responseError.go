package rendering

// ResponseError is only used to return a logic error to the client
type ResponseError struct {
	Errors []string `json:"errors"`
}
