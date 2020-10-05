package httputils

import "net/http"

// IsHttpError returns true if a status code is 400 or above
func IsHttpError(code int) bool {
	return code > 399
}

// IsHttpSuccess returns true if a status code is of the 2xx range (200 - 299)
func IsHttpSuccess(code int) bool {
	return code >= http.StatusOK && code <= 299
}
