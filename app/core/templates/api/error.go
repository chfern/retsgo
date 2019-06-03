package api

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
	}
}
