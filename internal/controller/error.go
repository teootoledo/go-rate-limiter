package controller

import (
	"net/http"
	"rateLimiter/internal/external/resources"

	"github.com/gin-gonic/gin"
)

func SetErrorResponseBadRequest(ctx *gin.Context, code int, message, details string) {
	ctx.JSON(code, resources.ErrorResponseHTTPBadRequest{
		StatusCode: http.StatusBadRequest,
		ErrorResponse: resources.ErrorResponse{
			Message: message,
			Details: details,
		},
	})
}

func SetErrorResponseTooManyRequests(ctx *gin.Context, code int, message, details string) {
	ctx.JSON(code, resources.ErrorResponseHTTPBadRequest{
		StatusCode: http.StatusTooManyRequests,
		ErrorResponse: resources.ErrorResponse{
			Message: message,
			Details: details,
		},
	})
}
