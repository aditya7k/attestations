package provenance

type PredicateStatementDTO struct { //nolint:maligned
	ConfigURI          string
	ConfigDigest       map[string]string
	ConfigEntryPoint   string
	PredicateBuilderId string
	PredicateBuildType string
	SubjectName        string
	SubjectDigest      map[string]string
}

func BuildProvenanceStatement(provenanceStatementDTO PredicateStatementDTO) *Statement {

	invocationConfigSource := NewInvocationConfigSource(
		withConfigSourceUri(provenanceStatementDTO.ConfigURI),
		withConfigSourceDigest(provenanceStatementDTO.ConfigDigest),
		withConfigSourceEntryPoint(provenanceStatementDTO.ConfigEntryPoint))

	predicateBuilder := PredicateBuilder{ID: provenanceStatementDTO.PredicateBuilderId}

	predicateInvocation := PredicateInvocation{ConfigSource: *invocationConfigSource}

	predicate := NewSLSAProvenancePredicate(
		withBuildType(provenanceStatementDTO.PredicateBuildType),
		withPredicateBuilder(predicateBuilder),
		withPredicateInvocation(predicateInvocation))

	subjectStatement := NewSubjectStatement(
		withSubjectName(provenanceStatementDTO.SubjectName),
		withSubjectDigest(provenanceStatementDTO.SubjectDigest))

	statement := NewProvenanceStatement(
		withPredicate(*predicate),
		withSubject(*subjectStatement))

	return statement
}
