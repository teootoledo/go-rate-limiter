package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rateLimiter/internal/constants"
	"rateLimiter/internal/external/resources"
	"rateLimiter/internal/logs"
	"rateLimiter/internal/service"
)

type Mail interface {
	Send(ctx *gin.Context)
}

type mailController struct {
	logger      logs.Logger
	mailService service.Mail
}

// Send godoc
// @Summary Send mail
// @Description Allows to email a recipient with a subject and body
// @Tags mail
// @Accept json
// @Produce json
// @Param request body resources.SendMailRequest true "Mail details"
// @Success 200 {object} resources.SendMailRequest
// @Failure 400 {object} resources.ErrorResponseHTTPBadRequest "Bad Request"
// @Failure 429 {object} resources.ErrorResponseHTTPTooManyRequests "Too Many Requests"
// @Router /mail [post]
func (mc *mailController) Send(ctx *gin.Context) {
	var request resources.SendMailRequest
	if err := ctx.BindJSON(&request); err != nil {
		mc.logger.Error(ctx, "error in MailController#Send: bad request", err)
		SetErrorResponseBadRequest(ctx, http.StatusBadRequest, constants.BadRequestMessage, err.Error())
		return
	}

	err := mc.mailService.Send(ctx, request)
	if err != nil {
		mc.logger.Error(ctx, "error in MailController#Send: too many requests", err)
		SetErrorResponseTooManyRequests(ctx, http.StatusTooManyRequests, constants.TooManyRequestsMessage, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, request)
}

// NewMailController - returns a new mail controller
func NewMailController(mailService service.Mail) Mail {
	logger := logs.New("Mail Controller")

	return &mailController{
		mailService: mailService,
		logger:      logger,
	}
}
