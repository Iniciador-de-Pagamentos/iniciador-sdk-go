package helpers

import (
	"encoding/json"
	"net/url"
	"reflect"
)

// Helper function to compare two values properly
func IsEqual(a, b interface{}) bool {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return false
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return string(aJSON) == string(bJSON)
}

// AreURLQueryParamsEqual checks if two URL query parameters are equal regardless of their order.
func AreURLQueryParamsEqual(params1, params2 url.Values) bool {
	return reflect.DeepEqual(params1, params2)
}
