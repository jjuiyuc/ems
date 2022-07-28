// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "get token by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Show a token",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/PasswordResetByToken": {
            "put": {
                "description": "set a new password by having the token from email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Reset the password",
                "parameters": [
                    {
                        "description": "Token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "New password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/passwordLost": {
            "put": {
                "description": "get email by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Send an email for reset the password",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get user by token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Show the detailed information about an individual user",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Input user's access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/services.ProfileResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/{gwid}/devices/battery/energy-info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get battery by token, gateway UUID and startTime",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "energy resources"
                ],
                "summary": "Show the detailed information and current state about a battery",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Input user's access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Gateway UUID",
                        "name": "gwid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "description": "Example : UTC time in ISO-8601",
                        "name": "startTime",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/services.BatteryEnergyInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/{gwid}/devices/battery/power-state": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get battery by token, gateway UUID, resolution, startTime and endTime",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "energy resources"
                ],
                "summary": "Show today's hourly power state of a battery",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Input user's access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Gateway UUID",
                        "name": "gwid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "example": "UTC time in ISO-8601",
                        "name": "endTime",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "hour"
                        ],
                        "type": "string",
                        "name": "resolution",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date-time",
                        "example": "UTC time in ISO-8601",
                        "name": "startTime",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/app.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/services.BatteryPowerStateResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "routers.ZoomableQuery": {
            "type": "object",
            "required": [
                "endTime",
                "resolution",
                "startTime"
            ],
            "properties": {
                "endTime": {
                    "type": "string",
                    "format": "date-time",
                    "example": "UTC time in ISO-8601"
                },
                "resolution": {
                    "type": "string",
                    "enum": [
                        "hour"
                    ]
                },
                "startTime": {
                    "type": "string",
                    "format": "date-time",
                    "example": "UTC time in ISO-8601"
                }
            }
        },
        "services.BatteryEnergyInfoResponse": {
            "type": "object",
            "properties": {
                "batteryConsumedEnergyAC": {
                    "type": "number"
                },
                "batteryConsumedLifetimeEnergyAC": {
                    "type": "number"
                },
                "batteryLifetimeOperationCycles": {
                    "type": "number"
                },
                "batteryOperationCycles": {
                    "type": "number"
                },
                "batteryPower": {
                    "type": "number"
                },
                "batteryProducedEnergyAC": {
                    "type": "number"
                },
                "batteryProducedLifetimeEnergyAC": {
                    "type": "number"
                },
                "batterySoC": {
                    "type": "number"
                },
                "capcity": {
                    "type": "number"
                },
                "model": {
                    "type": "string"
                },
                "powerSources": {
                    "type": "string"
                },
                "voltage": {
                    "type": "number"
                }
            }
        },
        "services.BatteryPowerStateResponse": {
            "type": "object",
            "properties": {
                "batteryAveragePowerACs": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "onPeakTime": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "timestamps": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "services.GatewayInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "gatewayID": {
                    "type": "string"
                }
            }
        },
        "services.ProfileResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "expirationDate": {
                    "type": "string"
                },
                "gateways": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.GatewayInfo"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "lockedAt": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "passwordLastChanged": {
                    "type": "string"
                },
                "passwordRetryCount": {
                    "type": "string"
                },
                "pwdTokenExpiry": {
                    "type": "string"
                },
                "resetPWDToken": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "DER-EMS API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
