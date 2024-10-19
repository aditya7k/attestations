package provenance

import (
	"encoding/json"
	"fmt"
	"github.com/datasweet/jsonmap"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_CreateProvenanceStatement(t *testing.T) {

	// Arrange
	invocationConfigSource := NewInvocationConfigSource(
		withConfigSourceUri("https://github.com/example/repo"),
		withConfigSourceDigest(map[string]string{"sha1": "abc123"}),
		withConfigSourceEntryPoint("build script"))

	predicateBuilder := PredicateBuilder{ID: "example.com/builder"}

	predicateInvocation := PredicateInvocation{ConfigSource: *invocationConfigSource}

	predicate := NewSLSAProvenancePredicate(
		withBuildType("https://example.com/build/type"),
		withPredicateBuilder(predicateBuilder),
		withPredicateInvocation(predicateInvocation))

	subjectStatement := NewSubjectStatement(
		withSubjectName("example.com/my-artifact"),
		withSubjectDigest(map[string]string{"sha256": "abcd1234"}))

	statement := NewProvenanceStatement(
		withPredicate(*predicate),
		withSubject(*subjectStatement))

	// Act
	attBytes, err := json.Marshal(statement)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling statement: %v\n", err)
		return
	}

	// Assert
	var jsonMap map[string]interface{}
	err = json.Unmarshal(attBytes, &jsonMap)
	if err != nil {
		return
	}

	j := jsonmap.FromMap(jsonMap)

	assertJsonString(t, j, "example.com/builder", "predicate.builder.id")
	assertJsonString(t, j, "https://github.com/example/repo", "predicate.invocation.configSource.uri")
	assertJsonString(t, j, "build script", "predicate.invocation.configSource.entryPoint")
	assertJsonString(t, j, "example.com/my-artifact", "subject.name")
	assertJsonString(t, j, "abcd1234", "subject.digest.sha256")
	assertJsonString(t, j, "https://example.com/build/type", "predicate.buildType")
	assertJsonString(t, j, "https://slsa.dev/provenance/v0.1", "predicateType")
	assertJsonString(t, j, "https://in-toto.io/Statement/v0.1", "_type")
}

func assertJsonString(t *testing.T, j *jsonmap.Json, expected string, key string) {
	value := j.Get(key).AsString()
	if value == "" {
		t.Errorf("path %s does not exist", key)
	}
	assert.Equal(t, expected, value, "for path: '%s' expected: '%s', actual: '%s'", key, expected, value)
}
