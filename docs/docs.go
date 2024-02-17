// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Teo Martin Toledo",
            "email": "teootoledo@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/mail": {
            "post": {
                "description": "Allows to email a recipient with a subject and body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "Send mail",
                "parameters": [
                    {
                        "description": "Mail details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/resources.SendMailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resources.SendMailRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponseHTTPBadRequest"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/resources.ErrorResponseHTTPTooManyRequests"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "resources.ErrorResponse": {
            "description": "Error base info.",
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "error details"
                },
                "message": {
                    "type": "string",
                    "example": "status code message"
                }
            }
        },
        "resources.ErrorResponseHTTPBadRequest": {
            "description": "Error response for http BadRequest.",
            "type": "object",
            "properties": {
                "error-response": {
                    "$ref": "#/definitions/resources.ErrorResponse"
                },
                "status-code": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "resources.ErrorResponseHTTPTooManyRequests": {
            "description": "Error response for http TooManyRequests.",
            "type": "object",
            "properties": {
                "error-response": {
                    "$ref": "#/definitions/resources.ErrorResponse"
                },
                "status-code": {
                    "type": "integer",
                    "example": 429
                }
            }
        },
        "resources.SendMailRequest": {
            "type": "object",
            "required": [
                "body",
                "recipient",
                "subject",
                "type"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "status",
                        "news",
                        "marketing"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "NotificationService API (Rate Limiting)",
	Description:      "Rate Limiter challenge for Modak",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}