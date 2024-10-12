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

	predicateBuilder := PredicateBuilder{
		ID: "example.com/builder",
	}

	invocationConfigSource := InvocationConfigSource{
		URI:        "git+https://github.com/example/repo",
		Digest:     map[string]string{"sha1": "abc123"},
		EntryPoint: "build script",
	}

	predicateInvocation := PredicateInvocation{
		ConfigSource: invocationConfigSource,
	}

	predicate := SLSAProvenancePredicate{
		BuildType:  "https://example.com/build/type",
		Builder:    predicateBuilder,
		Invocation: predicateInvocation,
	}

	subjectStatement := SubjectStatement{
		Name:   "example.com/my-artifact",
		Digest: map[string]string{"sha256": "abcd1234"},
	}

	statement := ProvenanceStatement{
		Type:          inTotoStatementType,
		PredicateType: predicateType,
		Subject:       subjectStatement,
		Predicate:     predicate,
	}

	// Act

	// Step 3: Marshal the statement into JSON
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

	assert.Equal(t, "example.com/builder", j.Get("predicate.builder.id").AsString(), "for path predicate.builder.id expected 'example.com/builder'")
	assert.Equal(t, "git+https://github.com/example/repo", j.Get("predicate.invocation.configSource.uri").AsString(), "expected 'git+https://github.com/example/repo'")
	assert.Equal(t, "build script", j.Get("predicate.invocation.configSource.entryPoint").AsString(), "expected 'build script'")
	assert.Equal(t, "example.com/my-artifact", j.Get("subject.name").AsString(), "expected 'example.com/my-artifact'")
	assert.Equal(t, "abcd1234", j.Get("subject.digest.sha256").AsString(), "expected 'abcd1234'")
	assert.Equal(t, "https://example.com/build/type", j.Get("predicate.buildType").AsString(), "expected 'https://example.com/build/type'")
	assert.Equal(t, "https://slsa.dev/provenance/v0.1", j.Get("predicateType").AsString(), "expected 'https://slsa.dev/provenance/v0.1'")
	assert.Equal(t, "https://in-toto.io/Statement/v0.1", j.Get("_type").AsString(), "expected 'https://in-toto.io/Statement/v0.1'")
}
