package gox_app

import (
	"encoding/json"
	"github.com/harishb2k/gox-base/serialization"
	"github.com/pkg/errors"
)

func NewJsonRequestParser(obj interface{}) RequestParser {
	return func(request Request) (parsedRequestBody interface{}, err error) {
		err = serialization.JsonBytesToObject(request.GetRequestBody(), obj)
		if err != nil {
			if _, ok := err.(*json.InvalidUnmarshalError); ok {
				return obj, errors.Wrap(err, "it seems you did not pass object pointer but object itself")
			}
		}
		return obj, err
	}
}
