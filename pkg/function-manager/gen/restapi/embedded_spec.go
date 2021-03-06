///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Function Manager\n",
    "title": "Function Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1/function",
  "paths": {
    "/": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "List all existing functions",
        "operationId": "getFunctions",
        "parameters": [
          {
            "type": "string",
            "description": "Function state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter based on tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/getFunctionsOKBody"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Custom error",
            "schema": {
              "$ref": "#/definitions/Error"
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
          "Store"
        ],
        "summary": "Add a new function",
        "operationId": "addFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Function created",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/runs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function runs that are being executed",
        "operationId": "getRuns",
        "responses": {
          "200": {
            "description": "List of function runs",
            "schema": {
              "$ref": "#/definitions/getRunsOKBody"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        }
      ]
    },
    "/{functionName}": {
      "get": {
        "description": "Returns a single function",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Find function by Name",
        "operationId": "getFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
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
          "Store"
        ],
        "summary": "Update a function",
        "operationId": "updateFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Deletes a function",
        "operationId": "deleteFunction",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to work on",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{functionName}/runs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function runs that are being executed",
        "operationId": "getFunctionRuns",
        "responses": {
          "200": {
            "description": "List of function runs",
            "schema": {
              "$ref": "#/definitions/getFunctionRunsOKBody"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
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
          "Runner"
        ],
        "summary": "Run a function",
        "operationId": "runFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful execution (blocking call)",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "202": {
            "description": "Execution started (non-blocking call)",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "400": {
            "description": "User error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Input object validation failed",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "502": {
            "description": "Function error occurred (blocking call)",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        },
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to run",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{functionName}/runs/{runName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function run by its name",
        "operationId": "getRun",
        "responses": {
          "200": {
            "description": "Function Run",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Function or Run not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to retrieve a run for",
          "name": "functionName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "format": "uuid",
          "description": "name of run to retrieve",
          "name": "runName",
          "in": "path",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        }
      ]
    }
  },
  "definitions": {
    "CloudEvent": {
      "type": "object",
      "required": [
        "namespace",
        "event-type",
        "cloud-events-version",
        "source-type",
        "source-id",
        "event-id"
      ],
      "properties": {
        "cloud-events-version": {
          "type": "string"
        },
        "content-type": {
          "type": "string"
        },
        "data": {
          "type": "string",
          "maxLength": 0
        },
        "event-id": {
          "type": "string"
        },
        "event-time": {
          "type": "string",
          "format": "date-time"
        },
        "event-type": {
          "type": "string",
          "maxLength": 128,
          "pattern": "^[\\w\\d\\-\\.]+$"
        },
        "event-type-version": {
          "type": "string"
        },
        "extensions": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "namespace": {
          "type": "string"
        },
        "schema-url": {
          "type": "string"
        },
        "source-id": {
          "type": "string"
        },
        "source-type": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "functionError": {
          "type": "object"
        },
        "message": {
          "type": "string"
        },
        "userError": {
          "type": "object"
        }
      }
    },
    "Function": {
      "type": "object",
      "required": [
        "name",
        "code",
        "image"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "image": {
          "type": "string"
        },
        "main": {
          "type": "string",
          "default": "main"
        },
        "modifiedTime": {
          "type": "integer"
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "schema": {
          "$ref": "#/definitions/Schema"
        },
        "secrets": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "$ref": "#/definitions/functionTags"
        }
      }
    },
    "Run": {
      "type": "object",
      "properties": {
        "blocking": {
          "type": "boolean"
        },
        "event": {
          "$ref": "#/definitions/CloudEvent"
        },
        "executedTime": {
          "type": "integer",
          "readOnly": true
        },
        "finishedTime": {
          "type": "integer",
          "readOnly": true
        },
        "functionId": {
          "type": "string",
          "readOnly": true
        },
        "functionName": {
          "type": "string",
          "readOnly": true
        },
        "input": {
          "type": "object"
        },
        "logs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "name": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "output": {
          "type": "object",
          "readOnly": true
        },
        "reason": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "secrets": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "$ref": "#/definitions/Status"
        },
        "tags": {
          "$ref": "#/definitions/runTags"
        }
      }
    },
    "Schema": {
      "type": "object",
      "properties": {
        "in": {
          "type": "object"
        },
        "out": {
          "type": "object"
        }
      }
    },
    "Status": {
      "type": "string",
      "enum": [
        "CREATING",
        "READY",
        "UPDATING",
        "ERROR",
        "DELETING"
      ]
    },
    "Tag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "functionTags": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Tag"
      },
      "x-go-gen-location": "models"
    },
    "getFunctionRunsOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Run"
      },
      "x-go-gen-location": "operations"
    },
    "getFunctionsOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Function"
      },
      "x-go-gen-location": "operations"
    },
    "getRunsOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Run"
      },
      "x-go-gen-location": "operations"
    },
    "runTags": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Tag"
      },
      "x-go-gen-location": "models"
    }
  },
  "securityDefinitions": {
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    }
  ],
  "tags": [
    {
      "description": "Crud operations on functions",
      "name": "Store"
    },
    {
      "description": "Execution operations on functions",
      "name": "Runner"
    }
  ]
}`))
}
