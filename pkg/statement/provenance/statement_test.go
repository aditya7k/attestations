package provenance

import (
	"encoding/json"
	"fmt"
	"github.com/datasweet/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateProvenanceStatement(t *testing.T) {

	// Arrange
	predicateStatementDTO := PredicateStatementDTO{
		configURI:          "https://github.com/example/repo",
		configDigest:       map[string]string{"sha1": "abc123"},
		configEntryPoint:   "build script",
		predicateBuilderId: "example.com/builder",
		predicateBuildType: "https://example.com/build/type",
		subjectName:        "example.com/my-artifact",
		subjectDigest:      map[string]string{"sha256": "abcd1234"},
	}

	statement := buildProvenanceStatement(predicateStatementDTO)

	// Act
	attBytes, err := json.Marshal(statement)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error marshalling statement: %v\n", err))
	}

	// Assert
	var jsonMap map[string]interface{}
	err = json.Unmarshal(attBytes, &jsonMap)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error unmarshalling statement: %v\n", err))
	}

	j := jsonmap.FromMap(jsonMap)

	tests := []struct {
		path          string
		expectedValue string
	}{
		{"predicate.invocation.configSource.uri", predicateStatementDTO.configURI},
		{"predicate.builder.id", predicateStatementDTO.predicateBuilderId},
		{"predicate.invocation.configSource.entryPoint", predicateStatementDTO.configEntryPoint},
		{"subject.name", predicateStatementDTO.subjectName},
		{"subject.digest.sha256", predicateStatementDTO.subjectDigest["sha256"]},
		{"predicate.buildType", predicateStatementDTO.predicateBuildType},
		{"predicateType", predicateType},
		{"_type", inTotoStatementType},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assertJsonString(t, j, tt.expectedValue, tt.path)
		})
	}

}

func assertJsonString(t *testing.T, j *jsonmap.Json, expected string, key string) {
	value := j.Get(key).AsString()
	if value == "" {
		assert.Fail(t, fmt.Sprintf("key not found: %s", key))
	}
	assert.Equal(t, expected, value, "for path: '%s' expected: '%s', actual: '%s'", key, expected, value)
}
