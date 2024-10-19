package provenance

type PredicateInvocation struct {
	ConfigSource InvocationConfigSource `json:"configSource"`
}

type InvocationConfigSource struct {
	URI        string            `json:"uri"`
	Digest     map[string]string `json:"digest"`
	EntryPoint string            `json:"entryPoint"`
}

func NewInvocationConfigSource(options ...func(_ *InvocationConfigSource)) *InvocationConfigSource {
	invocationConfigSource := &InvocationConfigSource{}
	for _, o := range options {
		o(invocationConfigSource)
	}
	return invocationConfigSource
}

func withConfigSourceUri(uri string) func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.URI = uri
	}
}

func withConfigSourceDigest(digest map[string]string) func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.Digest = digest
	}
}

func withConfigSourceEntryPoint(entryPoint string) func(invocationConfigSource *InvocationConfigSource) {
	return func(invocationConfigSource *InvocationConfigSource) {
		invocationConfigSource.EntryPoint = entryPoint
	}
}
