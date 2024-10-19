package statement

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
		withConfigSourceUri(),
		withConfigSourceDigest(),
		withConfigSourceEntryPoint())

	predicateBuilder := PredicateBuilder{ID: "example.com/builder"}

	predicateInvocation := PredicateInvocation{ConfigSource: *invocationConfigSource}

	predicate := NewSLSAProvenancePredicate(
		withBuildType(),
		withPredicateBuilder(predicateBuilder),
		withPredicateInvocation(predicateInvocation))

	subjectStatement := NewSubjectStatement(
		withSubjectName(),
		withSubjectDigest())

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

func NewInvocationConfigSource(options ...func(_ *InvocationConfigSource)) *InvocationConfigSource {
	invocationConfigSource := &InvocationConfigSource{}
	for _, o := range options {
		o(invocationConfigSource)
	}
	return invocationConfigSource
}

func withConfigSourceUri() func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.URI = "https://github.com/example/repo"
	}
}

func withConfigSourceDigest() func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.Digest = map[string]string{"sha1": "abc123"}
	}
}

func withConfigSourceEntryPoint() func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.EntryPoint = "build script"
	}
}

func NewSLSAProvenancePredicate(options ...func(_ *SLSAProvenancePredicate)) *SLSAProvenancePredicate {
	predicate := &SLSAProvenancePredicate{}
	for _, o := range options {
		o(predicate)
	}
	return predicate
}

func withBuildType() func(predicate *SLSAProvenancePredicate) {
	return func(predicate *SLSAProvenancePredicate) {
		predicate.BuildType = "https://example.com/build/type"
	}
}

func withPredicateBuilder(predicateBuilder PredicateBuilder) func(predicate *SLSAProvenancePredicate) {
	return func(predicate *SLSAProvenancePredicate) {
		predicate.Builder = predicateBuilder
	}
}

func withPredicateInvocation(predicateInvocation PredicateInvocation) func(predicate *SLSAProvenancePredicate) {
	return func(predicate *SLSAProvenancePredicate) {
		predicate.Invocation = predicateInvocation
	}
}

func NewSubjectStatement(options ...func(_ *SubjectStatement)) *SubjectStatement { //nolint:interfacer
	subjectStatement := &SubjectStatement{}
	for _, o := range options {
		o(subjectStatement)
	}
	return subjectStatement
}

func withSubjectName() func(subjectStatement *SubjectStatement) {
	return func(subjectStatement *SubjectStatement) {
		subjectStatement.Name = "example.com/my-artifact"
	}
}

func withSubjectDigest() func(subjectStatement *SubjectStatement) {
	return func(subjectStatement *SubjectStatement) {
		subjectStatement.Digest = map[string]string{"sha256": "abcd1234"}
	}
}

func NewProvenanceStatement(options ...func(_ *ProvenanceStatement)) *ProvenanceStatement {
	provenanceStatement := &ProvenanceStatement{}
	for _, o := range options {
		o(provenanceStatement)
	}
	withInTotoStatementType()(provenanceStatement)
	withPredicateType()(provenanceStatement)
	return provenanceStatement
}

func withInTotoStatementType() func(statement *ProvenanceStatement) {
	return func(statement *ProvenanceStatement) {
		statement.Type = inTotoStatementType
	}
}

func withPredicateType() func(statement *ProvenanceStatement) {
	return func(statement *ProvenanceStatement) {
		statement.PredicateType = predicateType
	}
}

func withPredicate(predicate SLSAProvenancePredicate) func(statement *ProvenanceStatement) {
	return func(statement *ProvenanceStatement) {
		statement.Predicate = predicate
	}
}

func withSubject(subject SubjectStatement) func(statement *ProvenanceStatement) {
	return func(statement *ProvenanceStatement) {
		statement.Subject = subject
	}
}
