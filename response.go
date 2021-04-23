package gox_app

import (
	"fmt"
	"net/http"
)

// Response data to send out
type ResponseData struct {
	Status      int
	BodyObject  interface{}
	Body        []byte
	Err         error
	ContentType ContentType
}

// Send back response
//go:generate mockgen -source=response.go -destination=mocks/mock_resposne.go -package=mock_gox_app
type Response interface {
	Done(data *ResponseData)
}

// Get the content type as string
func (r *ResponseData) ContentTypeAsString() string {
	return fmt.Sprintf("%s", r.ContentType)
}

/*type responseDataBuilder interface {
	WithStatusCode(int) responseDataBuilder
	WithBody([]byte) responseDataBuilder
	WithObjectBody(interface{}) responseDataBuilder
	WithError(error) responseDataBuilder
	WithContentType(ContentType) responseDataBuilder
	Build() *ResponseData
}
*/
type internalResponseDataBuilder struct {
	responseData *ResponseData
}

func (i *internalResponseDataBuilder) WithStatusCode(status int) *internalResponseDataBuilder {
	i.responseData.Status = status
	return i
}

func (i *internalResponseDataBuilder) WithBody(body []byte) *internalResponseDataBuilder {
	i.responseData.Body = body
	return i
}

func (i *internalResponseDataBuilder) WithObjectBody(body interface{}) *internalResponseDataBuilder {
	i.responseData.BodyObject = body
	return i
}

func (i *internalResponseDataBuilder) WithError(err error) *internalResponseDataBuilder {
	i.responseData.Err = err
	return i
}

func (i *internalResponseDataBuilder) WithContentType(contentType ContentType) *internalResponseDataBuilder {
	i.responseData.ContentType = contentType
	return i
}

func (i *internalResponseDataBuilder) Build() *ResponseData {
	if i.responseData.ContentType == "" {
		i.responseData.ContentType = ContentTypeJson
	}
	if i.responseData.Status == 0 {
		i.responseData.Status = http.StatusOK
	}
	return i.responseData
}

func NewResponseDataBuilder() *internalResponseDataBuilder {
	return &internalResponseDataBuilder{responseData: &ResponseData{ContentType: ContentTypeJson, Status: http.StatusOK}}
}
