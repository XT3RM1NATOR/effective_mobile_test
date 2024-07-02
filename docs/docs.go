// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": ".",
            "url": "."
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update task details",
                "parameters": [
                    {
                        "description": "Updated task details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Add a new task",
                "parameters": [
                    {
                        "description": "Task details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.AddTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task created successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{taskId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete task by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid task ID",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get tasks by user ID with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.TaskResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid user ID or pagination parameters",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Add a new user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.AddUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User added successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload or user ID",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.AddTaskRequest": {
            "type": "object",
            "properties": {
                "taskName": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.AddUserRequest": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.TaskResponse": {
            "type": "object",
            "properties": {
                "endTime": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                },
                "taskId": {
                    "type": "integer"
                },
                "taskName": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UpdateTaskRequest": {
            "type": "object",
            "properties": {
                "isFinished": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "taskId": {
                    "type": "integer"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "github_com_XT3RM1NAT0R_time-tracker_internal_delivery_models.UserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passportNumber": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI 2.0"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Time Tracker",
	Description:      "This is the backend server for the test assignment.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}