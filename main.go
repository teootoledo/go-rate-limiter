package main

import (
	"rateLimiter/internal/app"
)

//	@title			NotificationService API (Rate Limiting)
//	@version		1.0.0
//	@description	Rate Limiter challenge for Modak

//	@contact.name	Teo Martin Toledo
//	@contact.email	teootoledo@gmail.com

//	@host		localhost:8080
//	@BasePath	/v1

func main() {
	app.NewApp().
		Setup().
		InitServer()
}
