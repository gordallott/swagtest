// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TestPostReaderHandlerFunc turns a function with the right signature into a test post reader handler
type TestPostReaderHandlerFunc func(TestPostReaderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TestPostReaderHandlerFunc) Handle(params TestPostReaderParams) middleware.Responder {
	return fn(params)
}

// TestPostReaderHandler interface for that can handle valid test post reader params
type TestPostReaderHandler interface {
	Handle(TestPostReaderParams) middleware.Responder
}

// NewTestPostReader creates a new http.Handler for the test post reader operation
func NewTestPostReader(ctx *middleware.Context, handler TestPostReaderHandler) *TestPostReader {
	return &TestPostReader{Context: ctx, Handler: handler}
}

/*
	TestPostReader swagger:route POST /test testPostReader

TestPostReader test post reader API
*/
type TestPostReader struct {
	Context *middleware.Context
	Handler TestPostReaderHandler
}

func (o *TestPostReader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTestPostReaderParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
