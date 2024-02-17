package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"rateLimiter/internal/external/resources"
)

type MockMailService struct {
	mock.Mock
}

func (mock *MockMailService) Send(context context.Context, mailRequest resources.SendMailRequest) error {
	args := mock.Called(context, mailRequest)
	return args.Error(0)
}
