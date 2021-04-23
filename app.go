package gox_app

import "github.com/harishb2k/gox-base/lock"

type PrometheusConfig struct {
	Enabled bool
	Path    string
}

type ApplicationConfig struct {
	PrometheusConfig PrometheusConfig
	Harish           string
}

// Contains information to launch application
type RunConfig struct {
	Host string
	Port int
}

// Content type
type ContentType string

func (c ContentType) String() string {
	return string(c)
}

// Some of the content types supported
const (
	ContentTypeJson ContentType = "application/json"
)

// Complete request context will all data
type RequestContext struct {
	RequestParser   RequestParser
	ResponseBuilder ResponseBuilder
}

// Everything required to handle a request
type RouteConfig struct {
	Name            string
	Path            string               // Path to match
	Consumes        []ContentType        // What content-type is handled by this route
	Produces        []ContentType        // What content-type is returned by this route
	HandlerFunc     HandlerFunc          // The actual handler method to process this request
	RequestParser   RequestParser        // Parse request data
	ResponseBuilder ResponseBuilder      // Build response to be sent to client
	Lock            lock.DistributedLock // Lock to be taken when we process this request
}

// Provides a router to setup rest end-point and methods
type Router interface {

	// Run a router
	Run(runConfig RunConfig)

	// All HTTP methods
	GET(config RouteConfig)
	POST(config RouteConfig)
	PUT(config RouteConfig)
	DELETE(config RouteConfig)
}
