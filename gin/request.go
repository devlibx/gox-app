package gin_app

import (
	"github.com/gin-gonic/gin"
	app "github.com/harishb2k/gox-app"
)

type ginRequestContext struct {
	*gin.Context
	routeConfig       app.RouteConfig
	requestBody       []byte
	parsedRequestBody interface{}
}

func (c *ginRequestContext) GetMetricName() string {
	return c.routeConfig.Name
}

func (c *ginRequestContext) GetRequestBody() []byte {
	return c.requestBody
}

func (c *ginRequestContext) GetParsedRequestBody() interface{} {
	return c.parsedRequestBody
}

func (c *ginRequestContext) GetHeaders(name string) ([]string, bool) {
	panic("implement me")
}

func (c *ginRequestContext) GetQuery(name string) ([]string, bool) {
	panic("implement me")
}

func (c *ginRequestContext) GetPathParam(name string) (string, bool) {
	name = c.Param(name)
	if name != "" {
		return name, true
	} else {
		return name, false
	}
}
