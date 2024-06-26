// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1": {
            "post": {
                "description": "Creates a shortened version of a given URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Shortens a URL",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.ShortenURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.ShortenURLResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/{url}": {
            "get": {
                "description": "Redirects to the original URL corresponding to the given short URL.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Resolves a shortened URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL",
                        "name": "url",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Moved Permanently"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/routes.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "routes.ShortenURLRequest": {
            "type": "object",
            "properties": {
                "expiry": {
                    "type": "integer"
                },
                "short": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "routes.ShortenURLResponse": {
            "type": "object",
            "properties": {
                "expiry": {
                    "type": "integer"
                },
                "rate_limit": {
                    "type": "integer"
                },
                "rate_limit_reset": {
                    "type": "integer"
                },
                "short": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "This is a URL shortener API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
