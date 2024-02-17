package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"rateLimiter/internal/external/resources"
	"testing"
)

type sendMailCase struct {
	testName         string
	amountOfRequests int
	mailRequest      resources.SendMailRequest
	expectedError    error
}

func TestMailService_Send(t *testing.T) {
	testCases := []sendMailCase{
		{
			testName:         "Error - No rate limit set for the given notification type",
			amountOfRequests: 1,
			mailRequest: resources.SendMailRequest{
				Type:      "invalid",
				Subject:   "good news!",
				Body:      "this a good news mail",
				Recipient: "example@modak.com",
			},
			expectedError: errors.New("no rate limit set for the given type"),
		},
		{
			testName:         "Success - status mail sent successfully",
			amountOfRequests: 2,
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Subject:   "status update",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			expectedError: nil,
		},
		{
			testName:         "Error - status mail rate limit exceeded",
			amountOfRequests: 3,
			mailRequest: resources.SendMailRequest{
				Type:      "status",
				Subject:   "status update",
				Body:      "this a status mail",
				Recipient: "example@modak.com",
			},
			expectedError: errors.New("rate limit exceeded"),
		},
		{
			testName:         "Success - news mail sent successfully",
			amountOfRequests: 1,
			mailRequest: resources.SendMailRequest{
				Type:      "news",
				Subject:   "good news!",
				Body:      "this a good news mail",
				Recipient: "example@modak.com",
			},
			expectedError: nil,
		},
		{
			testName:         "Error - news mail rate limit exceeded",
			amountOfRequests: 2,
			mailRequest: resources.SendMailRequest{
				Type:      "news",
				Subject:   "good news!",
				Body:      "this a good news mail",
				Recipient: "example@modak.com",
			},
			expectedError: errors.New("rate limit exceeded"),
		},
		{
			testName:         "Success - marketing mail sent successfully",
			amountOfRequests: 3,
			mailRequest: resources.SendMailRequest{
				Type:      "marketing",
				Subject:   "marketing mail",
				Body:      "this a marketing mail",
				Recipient: "example@modak.com",
			},
			expectedError: nil,
		},
		{
			testName:         "Error - marketing mail rate limit exceeded",
			amountOfRequests: 4,
			mailRequest: resources.SendMailRequest{
				Type:      "marketing",
				Subject:   "marketing mail",
				Body:      "this a marketing mail",
				Recipient: "example@modak.com",
			},
			expectedError: errors.New("rate limit exceeded"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {

			mailService := NewMailService()

			for i := 0; i < test.amountOfRequests; i++ {
				err := mailService.Send(context.Background(), test.mailRequest)
				if err != nil {
					require.Equal(t, test.expectedError, err)
				}
			}
		})
	}
}
