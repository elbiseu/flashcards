package responses

import "encoding/json"

type APIError struct {
	errorCode string
	message   string
}

func (a *APIError) JSON() ([]byte, error) {
	data := map[string]interface{}{
		"error_code": a.errorCode,
		"message":    a.message,
	}
	return json.Marshal(data)
}

func newAPIError(errorCode, message string) *APIError {
	return &APIError{
		errorCode: errorCode,
		message:   message,
	}
}

var (
	InternalServerError = newAPIError("10002", "Internal Server Error")
)
