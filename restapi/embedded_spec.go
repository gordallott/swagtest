// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test",
    "version": "2.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/test": {
      "post": {
        "operationId": "testPostReader",
        "parameters": [
          {
            "type": "string",
            "format": "datetime",
            "name": "aDate",
            "in": "query"
          },
          {
            "name": "payload",
            "in": "body",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "the version of the invoked instance",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test",
    "version": "2.0.0"
  },
  "basePath": "/api",
  "paths": {
    "/test": {
      "post": {
        "operationId": "testPostReader",
        "parameters": [
          {
            "type": "string",
            "format": "datetime",
            "name": "aDate",
            "in": "query"
          },
          {
            "name": "payload",
            "in": "body",
            "schema": {
              "type": "string",
              "format": "binary"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "the version of the invoked instance",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
}
