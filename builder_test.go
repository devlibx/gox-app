package gox_app

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRouterConfigBuilder(t *testing.T) {

	handlerFunc := func(ctx context.Context, request Request, response Response, requestContext RequestContext) {
	}
	parser := func(Request) (parsedRequestBody interface{}, err error) {
		return "My_Parser", nil
	}
	responseBuilder := func(request Request, response Response, object interface{}) ([]byte, error) {
		return []byte("My_Builder"), nil
	}

	request := NewRouterConfigBuilder("example").
		WithPath("/path").
		WithConsumeType(ContentTypeJson).
		WithProduceType(ContentTypeJson).
		WithHandlerFunc(handlerFunc).
		WithRequestParser(parser).
		WithResponseBuilder(responseBuilder).
		Build()
	assert.NotNil(t, request)
	assert.Equal(t, "example", request.Name)
	assert.Equal(t, "/path", request.Path)
	assert.Equal(t, ContentTypeJson, request.Consumes[0])
	assert.Equal(t, ContentTypeJson, request.Produces[0])
	p, _ := parser(nil)
	assert.Equal(t, p, "My_Parser")
	r, _ := responseBuilder(nil, nil, nil)
	assert.Equal(t, string(r), "My_Builder")

	request = NewRouterConfigBuilder("example").
		WithPath("/path").
		WithConsumeType(ContentTypeJson).
		WithHandlerFunc(handlerFunc).
		WithRequestParser(parser).
		WithResponseBuilder(responseBuilder).
		Build()
	assert.NotNil(t, request)
	assert.Equal(t, "/path", request.Path)
	assert.Equal(t, ContentTypeJson, request.Consumes[0])
	assert.Equal(t, ContentTypeJson, request.Produces[0])
	p, _ = parser(nil)
	assert.Equal(t, p, "My_Parser")
	r, _ = responseBuilder(nil, nil, nil)
	assert.Equal(t, string(r), "My_Builder")

	request = NewRouterConfigBuilder("example").
		WithPath("/path").
		WithProduceType(ContentTypeJson).
		WithHandlerFunc(handlerFunc).
		WithRequestParser(parser).
		WithResponseBuilder(responseBuilder).
		Build()
	assert.NotNil(t, request)
	assert.Equal(t, "/path", request.Path)
	assert.Equal(t, ContentTypeJson, request.Consumes[0])
	assert.Equal(t, ContentTypeJson, request.Produces[0])
	p, _ = parser(nil)
	assert.Equal(t, p, "My_Parser")
	r, _ = responseBuilder(nil, nil, nil)
	assert.Equal(t, string(r), "My_Builder")
}
