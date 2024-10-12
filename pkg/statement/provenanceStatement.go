package statement

const inTotoStatementType = "https://in-toto.io/Statement/v0.1"
const predicateType = "https://slsa.dev/provenance/v0.1"

type ProvenanceStatement struct {
	Type          string                  `json:"_type"`
	PredicateType string                  `json:"predicateType"`
	Subject       SubjectStatement        `json:"subject"`
	Predicate     SLSAProvenancePredicate `json:"predicate"`
}

type SubjectStatement struct {
	Name   string            `json:"name"`
	Digest map[string]string `json:"digest"`
}

type PredicateBuilder struct {
	ID string `json:"id"`
}

type PredicateInvocation struct {
	ConfigSource InvocationConfigSource `json:"configSource"`
}

type InvocationConfigSource struct {
	URI        string            `json:"uri"`
	Digest     map[string]string `json:"digest"`
	EntryPoint string            `json:"entryPoint"`
}

type SLSAProvenancePredicate struct {
	BuildType  string           `json:"buildType"`
	Builder    PredicateBuilder `json:"builder"`
	Invocation struct {
		ConfigSource InvocationConfigSource `json:"configSource"`
	} `json:"invocation"`
}
