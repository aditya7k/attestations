package provenance

type PredicateStatementDTO struct { //nolint:maligned
	configURI          string
	configDigest       map[string]string
	configEntryPoint   string
	predicateBuilderId string
	predicateBuildType string
	subjectName        string
	subjectDigest      map[string]string
}

func buildProvenanceStatement(provenanceStatementDTO PredicateStatementDTO) *Statement {

	invocationConfigSource := NewInvocationConfigSource(
		withConfigSourceUri(provenanceStatementDTO.configURI),
		withConfigSourceDigest(provenanceStatementDTO.configDigest),
		withConfigSourceEntryPoint(provenanceStatementDTO.configEntryPoint))

	predicateBuilder := PredicateBuilder{ID: provenanceStatementDTO.predicateBuilderId}

	predicateInvocation := PredicateInvocation{ConfigSource: *invocationConfigSource}

	predicate := NewSLSAProvenancePredicate(
		withBuildType(provenanceStatementDTO.predicateBuildType),
		withPredicateBuilder(predicateBuilder),
		withPredicateInvocation(predicateInvocation))

	subjectStatement := NewSubjectStatement(
		withSubjectName(provenanceStatementDTO.subjectName),
		withSubjectDigest(provenanceStatementDTO.subjectDigest))

	statement := NewProvenanceStatement(
		withPredicate(*predicate),
		withSubject(*subjectStatement))

	return statement
}
