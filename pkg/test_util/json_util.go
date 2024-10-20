package test_util

import (
	"encoding/json"
	"fmt"
	"github.com/datasweet/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FromJson(t *testing.T, attBytes []byte) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(attBytes, &jsonMap)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error unmarshalling: %v\n", err))
	}
	return jsonMap
}

func ToJson(t *testing.T, v any) []byte {
	attBytes, err := json.Marshal(v)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error marshalling: %v\n", err))
	}
	return attBytes
}

func AssertJsonString(t *testing.T, j *jsonmap.Json, expected string, key string) {
	value := j.Get(key).AsString()
	if value == "" {
		assert.Fail(t, fmt.Sprintf("key not found: %s", key))
	}
	assert.Equal(t, expected, value, "for path: '%s' expected: '%s', actual: '%s'", key, expected, value)
}
