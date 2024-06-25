package goai

import (
	"github.com/xrc360/golang/internal/json"
	"github.com/xrc360/golang/text/gstr"
	"github.com/xrc360/golang/util/gconv"
)

// ExternalDocs is specified by OpenAPI/Swagger standard version 3.0.
type ExternalDocs struct {
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

func (ed *ExternalDocs) UnmarshalValue(value interface{}) error {
	var valueBytes = gconv.Bytes(value)
	if json.Valid(valueBytes) {
		return json.UnmarshalUseNumber(valueBytes, ed)
	}
	var (
		valueString = string(valueBytes)
		valueArray  = gstr.Split(valueString, "|")
	)
	ed.URL = valueArray[0]
	if len(valueArray) > 1 {
		ed.Description = valueArray[1]
	}
	return nil
}
