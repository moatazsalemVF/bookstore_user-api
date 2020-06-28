package errors

import "net/http"

//RestError is the global error format for JSON
type RestError struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
}

//NewBadRequestError is a generation function to return bad request error
func NewBadRequestError(message string) *RestError {
	restErr := RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	return &restErr
}

//NewInternalServerError is a generation function to return bad request error
func NewInternalServerError(message string) *RestError {
	restErr := RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	return &restErr
}

//NewNotFoundError is a generation function to return not found error
func NewNotFoundError(message string) *RestError {
	restErr := RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
	return &restErr
}
