package gox_app

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponseDataBuilder(t *testing.T) {
	response := NewResponseDataBuilder().
		WithStatusCode(200).
		WithObjectBody("test").
		WithError(errors.New("dummy")).
		Build()
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.Status)
	assert.Equal(t, ContentTypeJson, response.ContentType)
	assert.Equal(t, "test", response.BodyObject)
	assert.Equal(t, "dummy", response.Err.Error())
}
