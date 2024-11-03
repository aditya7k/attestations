package provenance

import (
	"attestations/pkg/test_util"
	"github.com/datasweet/jsonmap"
	"testing"
)

func Test_CreateProvenanceStatement(t *testing.T) {

	// Arrange
	predicateStatementDTO := PredicateStatementDTO{
		ConfigURI:          "https://github.com/example/repo",
		ConfigDigest:       map[string]string{"sha1": "abc123"},
		ConfigEntryPoint:   "build script",
		PredicateBuilderId: "example.com/builder",
		PredicateBuildType: "https://example.com/build/type",
		SubjectName:        "example.com/my-artifact",
		SubjectDigest:      map[string]string{"sha256": "abcd1234"},
	}

	// Act
	statement := BuildProvenanceStatement(predicateStatementDTO)
	statementBytes := test_util.ToJson(t, statement)

	// Assert
	statementJsonMap := test_util.FromJson(t, statementBytes)
	statementJson := jsonmap.FromMap(statementJsonMap)

	tests := []struct {
		jsonPath      string
		expectedValue string
	}{
		{"predicate.invocation.configSource.uri", predicateStatementDTO.ConfigURI},
		{"predicate.builder.id", predicateStatementDTO.PredicateBuilderId},
		{"predicate.invocation.configSource.entryPoint", predicateStatementDTO.ConfigEntryPoint},
		{"subject.name", predicateStatementDTO.SubjectName},
		{"subject.digest.sha256", predicateStatementDTO.SubjectDigest["sha256"]},
		{"predicate.buildType", predicateStatementDTO.PredicateBuildType},
		{"predicateType", predicateType},
		{"_type", inTotoStatementType},
	}

	for _, tt := range tests {
		t.Run(tt.jsonPath, func(t *testing.T) {
			test_util.AssertJsonString(t, statementJson, tt.expectedValue, tt.jsonPath)
		})
	}
}
