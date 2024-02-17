package service

import (
	"context"
	"rateLimiter/internal/external/resources"
	"rateLimiter/internal/logs"
	"sync"
)

type Mail interface {
	Send(ctx context.Context, mail resources.SendMailRequest) error
}

type mailService struct {
	logger     logs.Logger
	mu         sync.Mutex
	recipients map[string]*RateLimit
}

// Send - emails the mail content to the recipient
func (m *mailService) Send(ctx context.Context, mail resources.SendMailRequest) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	recipient, exists := m.recipients[mail.Recipient]
	if !exists {
		recipient = NewRateLimit()
		m.recipients[mail.Recipient] = recipient
	}

	if err := recipient.AllowNotification(mail.Type); err != nil {
		m.logger.Error(ctx, "error in MailService#Send: rate limit exceeded", err)
		return err
	}

	m.logger.Info(ctx, "sending mail to recipient")

	return nil
}

func NewMailService() Mail {
	logger := logs.New("Mail Service")

	return &mailService{
		logger:     logger,
		recipients: make(map[string]*RateLimit),
	}
}
