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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Employees": {
            "post": {
                "description": "Create employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create employee",
                "operationId": "create-employee",
                "parameters": [
                    {
                        "description": "Employee entity",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/bloodtypes": {
            "get": {
                "description": "list bloodtype entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List bloodtype entities",
                "operationId": "list-bloodtype",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Bloodtype"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create bloodtype",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create bloodtype",
                "operationId": "create-bloodtype",
                "parameters": [
                    {
                        "description": "Bloodtype entity",
                        "name": "bloodtype",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Bloodtype"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Bloodtype"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/bloodtypes/{id}": {
            "get": {
                "description": "get bloodtype by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a bloodtype entity by ID",
                "operationId": "get-bloodtype",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bloodtype ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Bloodtype"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees": {
            "get": {
                "description": "list employee entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List employee entities",
                "operationId": "list-employee",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Employee"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "description": "get employee by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a employee entity by ID",
                "operationId": "get-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/genders": {
            "get": {
                "description": "list gender entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List gender entities",
                "operationId": "list-gender",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Gender"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create gender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create gender",
                "operationId": "create-gender",
                "parameters": [
                    {
                        "description": "Gender entity",
                        "name": "gender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Gender"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Gender"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/genders/{id}": {
            "get": {
                "description": "get gender by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a gender entity by ID",
                "operationId": "get-gender",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Gender ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Gender"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/nametitles": {
            "get": {
                "description": "list nametitle entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List nametitle entities",
                "operationId": "list-nametitle",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.NameTitle"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create nametitle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create nametitle",
                "operationId": "create-nametitle",
                "parameters": [
                    {
                        "description": "NameTitle entity",
                        "name": "nametitle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.NameTitle"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.NameTitle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/nametitles/{id}": {
            "get": {
                "description": "get nametitle by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a nametitle entity by ID",
                "operationId": "get-nametitle",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "NameTitle ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.NameTitle"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/patients": {
            "get": {
                "description": "list patient entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List patient entities",
                "operationId": "list-patient",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Patient"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create patient",
                "operationId": "create-patient",
                "parameters": [
                    {
                        "description": "Patient entity",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Patient"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/statuss": {
            "get": {
                "description": "list status entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List status entities",
                "operationId": "list-status",
                "parameters": [
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
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Status"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create status",
                "operationId": "create-status",
                "parameters": [
                    {
                        "description": "Status entity",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Status"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Status"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/statuss/{id}": {
            "get": {
                "description": "get status by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a status entity by ID",
                "operationId": "get-status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Status ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Status"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Bloodtype": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BloodtypeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.BloodtypeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.BloodtypeEdges": {
            "type": "object",
            "properties": {
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Patient"
                    }
                }
            }
        },
        "ent.Employee": {
            "type": "object",
            "properties": {
                "Email": {
                    "description": "Email holds the value of the \"Email\" field.",
                    "type": "string"
                },
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "User_id": {
                    "description": "UserID holds the value of the \"User_id\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the EmployeeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.EmployeeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.EmployeeEdges": {
            "type": "object",
            "properties": {
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Patient"
                    }
                }
            }
        },
        "ent.Gender": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the GenderQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.GenderEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.GenderEdges": {
            "type": "object",
            "properties": {
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Patient"
                    }
                }
            }
        },
        "ent.NameTitle": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the NameTitleQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.NameTitleEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.NameTitleEdges": {
            "type": "object",
            "properties": {
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Patient"
                    }
                }
            }
        },
        "ent.Patient": {
            "type": "object",
            "properties": {
                "Address": {
                    "description": "Address holds the value of the \"Address\" field.",
                    "type": "string"
                },
                "Allergic": {
                    "description": "Allergic holds the value of the \"Allergic\" field.",
                    "type": "string"
                },
                "Congenital": {
                    "description": "Congenital holds the value of the \"Congenital\" field.",
                    "type": "string"
                },
                "Idcard": {
                    "description": "Idcard holds the value of the \"Idcard\" field.",
                    "type": "string"
                },
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PatientQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PatientEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PatientEdges": {
            "type": "object",
            "properties": {
                "bloodtype": {
                    "description": "Bloodtype holds the value of the bloodtype edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Bloodtype"
                },
                "employee": {
                    "description": "Employee holds the value of the employee edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Employee"
                },
                "gender": {
                    "description": "Gender holds the value of the gender edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Gender"
                },
                "nametitle": {
                    "description": "Nametitle holds the value of the nametitle edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.NameTitle"
                },
                "status": {
                    "description": "Status holds the value of the status edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Status"
                }
            }
        },
        "ent.Status": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StatusQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.StatusEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.StatusEdges": {
            "type": "object",
            "properties": {
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Patient"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API Patient",
	Description: "This is a sample server for SUT SE 2563",
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