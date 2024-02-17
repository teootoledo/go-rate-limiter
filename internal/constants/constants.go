package constants

const (

	// Server constants
	ServerPort = 8080

	// Logger constants
	RequestID       = "request_id"
	EnvLogLevel     = "LOG_LEVEL"
	EnvLogStdOut    = "LOG_STDOUT"
	EnvLogFormatter = "LOG_FORMATTER"

	// Error constants
	BadRequestMessage      = "bad request"
	TooManyRequestsMessage = "too many requests"

	// Path constants
	MailBasePath = "mail"

	// Email type constants
	EmailTypeStatus    = "status"
	EmailTypeNews      = "news"
	EmailTypeMarketing = "marketing"
)
