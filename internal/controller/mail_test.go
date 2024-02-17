package controller_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/baloo.v3"
	"net/http"
	"net/http/httptest"
	"rateLimiter/internal/app"
	"rateLimiter/internal/controller"
	"rateLimiter/internal/external/resources"
	"rateLimiter/internal/mock/service"
	"testing"
)

type MailControllerSuite struct {
	suite.Suite
}

func TestMailControllerSuite(t *testing.T) {
	suite.Run(t, new(MailControllerSuite))
}

type sendMailCase struct {
	testName           string
	mailRequest        resources.SendMailRequest
	serviceError       error
	expectedStatusCode int
	assert             func(res *http.Response, req *http.Request) error
}

func (suite *MailControllerSuite) TestSendMail() {
	testCases := []sendMailCase{
		{
			testName: "Error - Missing mail type",
			mailRequest: resources.SendMailRequest{
				Subject:   "status update",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Error - Missing mail subject",
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Error - Missing mail body",
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Subject:   "status update",
				Recipient: "example@modak.com",
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Error - Missing mail recipient",
			mailRequest: resources.SendMailRequest{
				Type:    "status",
				Subject: "status update",
				Body:    "this a status mail",
			},
			expectedStatusCode: http.StatusBadRequest,
			assert:             controller.AssertBadRequest,
		},
		{
			testName: "Error - No rate limit set for the given type",
			mailRequest: resources.SendMailRequest{
				Type:      "invalid",
				Subject:   "good news!",
				Body:      "this a good news mail",
				Recipient: "example@modak.com",
			},
			serviceError:       errors.New("no rate limit set for the given type"),
			expectedStatusCode: http.StatusTooManyRequests,
			assert:             controller.AssertTooManyRequests,
		},
		{
			testName: "Error - Rate limit exceeded",
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Subject:   "status update",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			serviceError:       errors.New("rate limit exceeded"),
			expectedStatusCode: http.StatusTooManyRequests,
			assert:             controller.AssertTooManyRequests,
		},
		{
			testName: "Success - Mail sent",
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Subject:   "status update",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			serviceError:       nil,
			expectedStatusCode: http.StatusOK,
			assert:             controller.AssertSuccess,
		},
	}

	for _, test := range testCases {
		suite.T().Run(test.testName, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			testApp := app.NewApp()
			testApp.Setup()

			mockMailService := new(service.MockMailService)
			testApp.MailService = mockMailService
			mailController := controller.NewMailController(testApp.MailService)
			testApp.MailController = mailController

			if test.expectedStatusCode != http.StatusBadRequest {
				mockMailService.On("Send", mock.Anything, test.mailRequest).Return(test.serviceError)
			}

			// Gin attaches dependencies by value so after changing dependencias at app lvl it is required to re-generate gin's engine
			testApp.ConfigureRoutes()
			testServer := httptest.NewServer(testApp.Engine)

			request := baloo.New(testServer.URL).
				Post("/v1/mail").
				JSON(test.mailRequest).
				Expect(t)

			_ = request.
				Status(test.expectedStatusCode).
				AssertFunc(test.assert).
				Done()
		})
	}
}
