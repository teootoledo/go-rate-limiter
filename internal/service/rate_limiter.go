package service

import (
	"errors"
	"golang.org/x/time/rate"
	"rateLimiter/internal/constants"
	"time"
)

type RateLimit struct {
	status    *rate.Limiter
	news      *rate.Limiter
	marketing *rate.Limiter
}

// AllowNotification - allows a notification to be sent
func (rl *RateLimit) AllowNotification(notificationType string) error {
	isAllowed := false
	switch notificationType {
	case constants.EmailTypeStatus:
		isAllowed = rl.status.Allow()
	case constants.EmailTypeNews:
		isAllowed = rl.news.Allow()
	case constants.EmailTypeMarketing:
		isAllowed = rl.marketing.Allow()
	default:
		return errors.New("no rate limit set for the given type")
	}

	if !isAllowed {
		return errors.New("rate limit exceeded")
	}

	return nil
}

func NewRateLimit() *RateLimit {

	return &RateLimit{
		status:    rate.NewLimiter(rate.Every(time.Minute), 2),
		news:      rate.NewLimiter(rate.Every(time.Hour*24), 1),
		marketing: rate.NewLimiter(rate.Every(time.Hour), 3),
	}
}
