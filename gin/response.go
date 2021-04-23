package gin_app

import (
	"github.com/gin-gonic/gin"
	app "github.com/harishb2k/gox-app"
)

type ginResponseContext struct {
	*gin.Context
	*app.RouteConfig
}

func (g *ginResponseContext) Done(data *app.ResponseData) {

	// Send status = 200 if missing
	if data.Status == 0 {
		data.Status = 200
	}

	// Send content type as json if missing
	if data.ContentType == "" {
		data.ContentType = app.ContentTypeJson
	}

	// Set status if we missed it
	if data.Err != nil && data.Status == 0 {
		data.Status = 500
	}

	if data.BodyObject != nil {
		if byteData, err := g.ResponseBuilder(nil, nil, data.BodyObject); err != nil {
			g.Writer.WriteHeader(500)
		} else {
			g.Header("content-type", data.ContentTypeAsString())
			g.Writer.WriteHeader(data.Status)
			if byteData != nil {
				_, _ = g.Writer.Write(byteData)
			}
		}
	} else if data.Body != nil {
		g.Header("content-type", data.ContentTypeAsString())
		g.Writer.WriteHeader(data.Status)
		_, _ = g.Writer.Write(data.Body)
	} else {
		g.Header("content-type", data.ContentTypeAsString())
		g.Writer.WriteHeader(data.Status)
	}
}
