package gox_app

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyRequest struct {
	Request
	body []byte
}

func (d dummyRequest) GetRequestBody() []byte {
	return d.body
}

func TestNewJsonRequestParser_WhenPointerIsPassed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type testStruct struct {
		Name string
		age  int
	}
	jsonString := `{
		"name": "user",
		"age": 10
	}`

	mockRequest := dummyRequest{body: []byte(jsonString)}
	rp := NewJsonRequestParser(&testStruct{})
	obj, err := rp(mockRequest)
	assert.NoError(t, err)
	assert.Equal(t, "user", obj.(*testStruct).Name)
	assert.Equal(t, 0, obj.(*testStruct).age)
}
