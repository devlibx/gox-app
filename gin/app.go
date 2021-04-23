package gin_app

import (
	c "context"
	"fmt"
	"github.com/gin-gonic/gin"
	app "github.com/harishb2k/gox-app"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-base/logger"
	"github.com/harishb2k/gox-base/metrics"
	"github.com/harishb2k/gox-base/serialization"
	"github.com/pkg/errors"
	"net/http"
)

type ginApp struct {
	app.ApplicationConfig
	engine *gin.Engine
	app.Router
	gox.CrossFunction
}

func NewGinApp(config app.ApplicationConfig, crossFunction gox.CrossFunction) app.Router {
	engine := gin.Default()

	// Setup Prometheus
	if config.PrometheusConfig.Enabled {
		path := "/metrics"
		if len(config.PrometheusConfig.Path) != 0 {
			path = config.PrometheusConfig.Path
		}
		engine.GET(path, func(context *gin.Context) {
			crossFunction.(metrics.Service).HttpHandler().ServeHTTP(context.Writer, context.Request)
		})
	}

	return &ginApp{
		engine:            engine,
		CrossFunction:     crossFunction,
		ApplicationConfig: config,
	}
}

func (a *ginApp) GET(routeConfig app.RouteConfig) {

	// Setup defaults in routes
	setupDefaultsInRouteConfig(&routeConfig)

	// Register a counter
	counterName := fmt.Sprintf("GET_%s", routeConfig.Name)
	_ = a.RegisterCounter(counterName, counterName+" Help", []string{"type"})

	// Setup GET on Gin
	a.engine.GET(routeConfig.Path, func(ginContext *gin.Context) {
		a.WithFields(logger.Fields{"path": routeConfig.Path}).Trace("GET")
		a.Counter(counterName).WithLabels(metrics.Labels{"type": "GET"}).Inc()

		// We will use Gin specific request
		request := &ginRequestContext{Context: ginContext, routeConfig: routeConfig}
		response := &ginResponseContext{Context: ginContext, RouteConfig: &routeConfig}

		// Run handler function
		routeConfig.HandlerFunc(c.Background(), request, response, app.RequestContext{})
	})
}

func (a *ginApp) POST(routeConfig app.RouteConfig) {

	// Setup defaults in routes
	setupDefaultsInRouteConfig(&routeConfig)

	// Register a counter
	counterName := fmt.Sprintf("POST_%s", routeConfig.Name)
	_ = a.RegisterCounter(counterName, counterName+" Help", []string{"type"})

	// Setup GET on Gin
	a.engine.POST(routeConfig.Path, func(ginContext *gin.Context) {
		a.WithFields(logger.Fields{"path": routeConfig.Path}).Trace("POST")
		a.Counter(counterName).WithLabels(metrics.Labels{"type": "POST"}).Inc()

		var err error

		// We will use Gin specific request
		request := &ginRequestContext{Context: ginContext, routeConfig: routeConfig}
		response := &ginResponseContext{Context: ginContext, RouteConfig: &routeConfig}

		// Read raw data from Gin request
		if request.requestBody, err = ginContext.GetRawData(); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to read raw body from request")).
					Build(),
			)
			return
		}

		// Try to parse request body and set it in request to be used
		if request.parsedRequestBody, err = routeConfig.RequestParser(request); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to parse request body to required object")).
					Build(),
			)
		} else {
			a.WithFields(logger.Fields{"path": routeConfig.Path}).Trace("POST Body=", request.GetParsedRequestBody())
			routeConfig.HandlerFunc(c.Background(), request, response, app.RequestContext{})
		}
	})
}

func (a *ginApp) PUT(routeConfig app.RouteConfig) {
	// Setup defaults in routes
	setupDefaultsInRouteConfig(&routeConfig)

	// Setup GET on Gin
	a.engine.PUT(routeConfig.Path, func(ginContext *gin.Context) {
		a.WithFields(logger.Fields{"path": routeConfig.Path}).Trace("POST")
		var err error

		// We will use Gin specific request
		request := &ginRequestContext{Context: ginContext, routeConfig: routeConfig}
		response := &ginResponseContext{Context: ginContext, RouteConfig: &routeConfig}

		// Read raw data from Gin request
		if request.requestBody, err = ginContext.GetRawData(); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to read raw body from request")).
					Build(),
			)
			return
		}

		// Try to parse request body and set it in request to be used
		if request.parsedRequestBody, err = routeConfig.RequestParser(request); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to parse request body to required object")).
					Build(),
			)
		} else {
			a.WithFields(logger.Fields{"path": routeConfig.Path}).Trace("POST Body=", request.GetParsedRequestBody())
			routeConfig.HandlerFunc(c.Background(), request, response, app.RequestContext{})
		}
	})
}

func (a *ginApp) DELETE(routeConfig app.RouteConfig) {
	// Setup defaults in routes
	setupDefaultsInRouteConfig(&routeConfig)

	// Setup GET on Gin
	a.engine.DELETE(routeConfig.Path, func(ginContext *gin.Context) {
		var err error

		// We will use Gin specific request
		request := &ginRequestContext{Context: ginContext, routeConfig: routeConfig}
		response := &ginResponseContext{Context: ginContext, RouteConfig: &routeConfig}

		// Read raw data from Gin request
		if request.requestBody, err = ginContext.GetRawData(); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to read raw body from request")).
					Build(),
			)
			return
		}

		// Try to parse request body and set it in request to be used
		if request.parsedRequestBody, err = routeConfig.RequestParser(request); err != nil {
			response.Done(
				app.NewResponseDataBuilder().
					WithStatusCode(http.StatusBadRequest).
					WithError(errors.Wrap(err, "failed to parse request body to required object")).
					Build(),
			)
		} else {
			routeConfig.HandlerFunc(c.Background(), request, response, app.RequestContext{})
		}
	})
}

func (a *ginApp) Run(runConfig app.RunConfig) {
	_ = a.engine.Run(fmt.Sprintf("%s:%d", runConfig.Host, runConfig.Port))
}

func setupDefaultsInRouteConfig(routeConfig *app.RouteConfig) {

	// Setup a default request parser
	if routeConfig.RequestParser == nil {
		routeConfig.RequestParser = func(request app.Request) (interface{}, error) {
			return request.GetRequestBody(), nil
		}
	}

	// Setup a default response parser parser
	if routeConfig.ResponseBuilder == nil {
		routeConfig.ResponseBuilder = func(request app.Request, response app.Response, object interface{}) ([]byte, error) {
			return serialization.ToBytes(object)
		}
	}
}
