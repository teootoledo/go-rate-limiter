package resources

type SendMailRequest struct {
	Body      string `json:"body" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Type      string `json:"type" binding:"required" enums:"status,news,marketing"`
}
