package gox_app

import "context"

// Parse input request e.g. we may need to parse input request to some specific struct
type RequestParser func(Request) (parsedRequestBody interface{}, err error)

// Helper to give a []byte to be sent
type ResponseBuilder func(request Request, response Response, object interface{}) ([]byte, error)

// A request handler function
type HandlerFunc func(ctx context.Context, request Request, response Response, requestContext RequestContext)
