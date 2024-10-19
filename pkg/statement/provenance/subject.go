package provenance

type SubjectStatement struct {
	Name   string            `json:"name"`
	Digest map[string]string `json:"digest"`
}

func NewSubjectStatement(options ...func(_ *SubjectStatement)) *SubjectStatement { //nolint:interfacer
	subjectStatement := &SubjectStatement{}
	for _, o := range options {
		o(subjectStatement)
	}
	return subjectStatement
}

func withSubjectName(name string) func(subjectStatement *SubjectStatement) {
	return func(subjectStatement *SubjectStatement) {
		subjectStatement.Name = name
	}
}

func withSubjectDigest(digest map[string]string) func(subjectStatement *SubjectStatement) {
	return func(subjectStatement *SubjectStatement) {
		subjectStatement.Digest = digest
	}
}
