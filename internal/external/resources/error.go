package resources

// ErrorResponse Error model
// @Description Error base info.
type ErrorResponse struct {
	Message string `json:"message" example:"status code message"`
	Details string `json:"details" example:"error details"`
}

// ErrorResponseHTTPBadRequest Error model
// @Description Error response for http BadRequest.
type ErrorResponseHTTPBadRequest struct {
	ErrorResponse ErrorResponse `json:"error-response"`
	StatusCode    int           `json:"status-code" example:"400"`
}

// ErrorResponseHTTPTooManyRequests Error model
// @Description Error response for http TooManyRequests.
type ErrorResponseHTTPTooManyRequests struct {
	ErrorResponse ErrorResponse `json:"error-response"`
	StatusCode    int           `json:"status-code" example:"429"`
}
