package provenance

const predicateType = "https://slsa.dev/provenance/v0.1"
const inTotoStatementType = "https://in-toto.io/Statement/v0.1"

type Statement struct {
	Type          string                  `json:"_type"`
	PredicateType string                  `json:"predicateType"`
	Subject       SubjectStatement        `json:"subject"`
	Predicate     SLSAProvenancePredicate `json:"predicate"`
}

func NewProvenanceStatement(options ...func(_ *Statement)) *Statement {
	provenanceStatement := &Statement{}
	for _, o := range options {
		o(provenanceStatement)
	}
	withInTotoStatementType()(provenanceStatement)
	withPredicateType()(provenanceStatement)
	return provenanceStatement
}

func withInTotoStatementType() func(statement *Statement) {
	return func(statement *Statement) {
		statement.Type = inTotoStatementType
	}
}

func withPredicateType() func(statement *Statement) {
	return func(statement *Statement) {
		statement.PredicateType = predicateType
	}
}

func withPredicate(predicate SLSAProvenancePredicate) func(statement *Statement) {
	return func(statement *Statement) {
		statement.Predicate = predicate
	}
}

func withSubject(subject SubjectStatement) func(statement *Statement) {
	return func(statement *Statement) {
		statement.Subject = subject
	}
}
