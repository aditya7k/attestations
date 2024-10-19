package provenance

type PredicateBuilder struct {
	ID string `json:"id"`
}

type SLSAProvenancePredicate struct {
	BuildType  string           `json:"buildType"`
	Builder    PredicateBuilder `json:"builder"`
	Invocation struct {
		ConfigSource InvocationConfigSource `json:"configSource"`
	} `json:"invocation"`
}

func NewSLSAProvenancePredicate(options ...func(_ *SLSAProvenancePredicate)) *SLSAProvenancePredicate {
	predicate := &SLSAProvenancePredicate{}
	for _, o := range options {
		o(predicate)
	}
	return predicate
}

func withBuildType(buildType string) func(predicate *SLSAProvenancePredicate) {
	return func(predicate *SLSAProvenancePredicate) {
		predicate.BuildType = buildType
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
