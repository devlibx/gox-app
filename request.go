package gox_app

// Provider of request information e.g. path param, headers, or query parameters
//go:generate mockgen -source=request.go -destination=mocks/mock_request.go -package=mock_gox_app
type Request interface {
	GetMetricName() string
	GetPathParam(name string) (string, bool) // Get a path param
	GetHeaders(name string) ([]string, bool) // Get all headers with given name
	GetQuery(name string) ([]string, bool)   // Get all query param with given name
	GetRequestBody() []byte                  // Get the original request body
	GetParsedRequestBody() interface{}       // Get the request which is same a original request body OR it is parsed by your request parser
}

// Extended provider of request information e.g. path param, headers, or query parameters
type RequestContextExt interface {
	Request

	GetPathParamAsInt(name string) (int, bool)
	GetPathParamAsFloat(name string) (float32, bool)
	GetPathParamAsBool(name string) (bool, bool)
	GetPathParamAsString(name string) (string, bool)

	GetOneHeaderAsInt(name string) (int, bool)
	GetOneHeaderAsFloat(name string) (int, bool)
	GetOneHeaderAsBool(name string) (int, bool)
	GetOneHeaderAsString(name string) (int, bool)

	GetOneQueryAsInt(name string) (int, bool)
	GetOneQueryAsFloat(name string) (int, bool)
	GetOneQueryAsBool(name string) (int, bool)
	GetOneQueryAsString(name string) (int, bool)
}
