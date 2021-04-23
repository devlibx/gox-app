
##Run Sample Application
```go
package main

import (
	c "context"
	"fmt"
	app "github.com/harishb2k/gox-app"
	ginapp "github.com/harishb2k/gox-app-gin"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-base/logger"
	"github.com/harishb2k/gox-base/metrics"
	"github.com/harishb2k/gox-base/serialization"
	prometheus "github.com/harishb2k/gox-metrics-prometheus"
)

func main() {

	type TestPojo struct {
		Input string `json:"input"`
		Age   int    `json:"age"`
	}

	// Run Application
	prometheusService := prometheus.NewPrometheusService(metrics.Configuration{})
	crossFunc := gox.NewCrossFunction(
		logger.NewLogger(logger.Configuration{LogLevel: logger.TraceLevel}),
		prometheusService,
	)
	config := app.ApplicationConfig{}
	config.PrometheusConfig.Enabled = true
	router := ginapp.NewGinApp(config, crossFunc)

	router.GET(
		app.NewRouterConfigBuilder("example").
			WithPath("/ping/:name").
			WithHandlerFunc(func(ctx c.Context, request app.Request, response app.Response, requestContext app.RequestContext) {
				name, _ := request.GetPathParam("name")
				data := TestPojo{
					Input: name,
					Age:   10,
				}
				response.Done(app.NewResponseDataBuilder().WithObjectBody(data).Build())
			}).
			Build(),
	)

	router.POST(
		app.NewRouterConfigBuilder("example").
			WithPath("/ping/:name").
			WithHandlerFunc(func(ctx c.Context, request app.Request, response app.Response, requestContext app.RequestContext) {
				name, _ := request.GetPathParam("name")
				body := request.GetParsedRequestBody().(*TestPojo)
				body.Input = fmt.Sprintf("%s-%s", body.Input, name)
				response.Done(app.NewResponseDataBuilder().WithObjectBody(body).Build())
			}).
			WithRequestParser(func(request app.Request) (parsedRequestBody interface{}, err error) {
				body := request.GetRequestBody()
				parsedBody := &TestPojo{}
				err = serialization.JsonBytesToObject(body, parsedBody)
				return parsedBody, err
			}).
			Build(),
	)

	router.Run(app.RunConfig{
		Host: "localhost",
		Port: 9090,
	})
}

```

## Test 
##### Get API
```shell script
curl -X GET http://localhost:9090/get/ping/some_name
 
Output:
{"input":"some_name","age":10}
```
##### Post API
```shell script
curl -X POST http://localhost:9090/post/ping/me \  
  -d '{
    "input": "test",
    "age": 38
}'

Output:
{
    "input": "test-me",
    "age": 38
}
```
