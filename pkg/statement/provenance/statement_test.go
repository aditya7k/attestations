package provenance

import (
	"attestations/pkg/test_util"
	"github.com/datasweet/jsonmap"
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
	attBytes := test_util.ToJson(t, statement)

	// Assert
	jsonMap := test_util.FromJson(t, attBytes)

	json := jsonmap.FromMap(jsonMap)

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
			test_util.AssertJsonString(t, json, tt.expectedValue, tt.path)
		})
	}

}
