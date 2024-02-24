package customhttp

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	DECODE_ERROR       = "Failed to decode body"
	UNKNOWN_TYPE_ERROR = "Unknown type passed for decoding"
)

// Decode flexibly decodes either response or request from web
func Decode(res interface{}, data any) error {
	var dec *json.Decoder

	switch v := res.(type) {
	case *http.Response:
		dec = json.NewDecoder(v.Body)
	case *http.Request:
		dec = json.NewDecoder(v.Body)
	default:
		return errors.New(UNKNOWN_TYPE_ERROR)
	}

	if err := dec.Decode(data); err != nil {
		return errors.New(DECODE_ERROR + ": " + err.Error())
	}
	return nil
}
