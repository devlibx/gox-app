package gin_app

import (
	"bytes"
	c "context"
	"fmt"
	app "github.com/harishb2k/gox-app"
	"github.com/harishb2k/gox-base"
	"github.com/harishb2k/gox-base/serialization"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostRequest(t *testing.T) {
	router := NewGinApp(app.ApplicationConfig{},gox.NewNoOpCrossFunction()).(*ginApp)

	ts := httptest.NewServer(router.engine)
	defer ts.Close()

	router.POST(
		app.NewRouterConfigBuilder("example").
			WithPath("/ping/:name").
			WithHandlerFunc(func(ctx c.Context, request app.Request, response app.Response, requestContext app.RequestContext) {
				name, _ := request.GetPathParam("name")
				body := request.GetParsedRequestBody().(*getTestAutoGenerated)
				body.Input = fmt.Sprintf("%s-%s", body.Input, name)
				response.Done(app.NewResponseDataBuilder().WithObjectBody(body).Build())
			}).
			WithRequestParser(func(request app.Request) (parsedRequestBody interface{}, err error) {
				body := request.GetRequestBody()
				parsedBody := &getTestAutoGenerated{}
				err = serialization.JsonBytesToObject(body, parsedBody)
				return parsedBody, err
			}).
			Build(),
	)

	// Make a GET call
	payload := getTestAutoGenerated{
		Input: "123",
		Age:   11,
	}
	reader := bytes.NewReader(serialization.ToBytesSuppressError(payload))
	resp, err := http.Post(fmt.Sprintf("%s/ping/user_1", ts.URL), app.ContentTypeJson.String(), reader)

	// Do basic check that we got 200 and correct content-type
	assert.NoError(t, err, "failed to make http call to /ping")
	assert.Equal(t, 200, resp.StatusCode)
	val, ok := resp.Header["Content-Type"]
	assert.True(t, ok)
	assert.Equal(t, app.ContentTypeJson.String(), val[0])

	// Check body
	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "failed to read body")

	obj := getTestAutoGenerated{}
	err = serialization.JsonBytesToObject(data, &obj)
	assert.NoError(t, err, "failed to parse body")
	assert.Equal(t, "123-user_1", obj.Input)
	assert.Equal(t, 11, obj.Age)
}
