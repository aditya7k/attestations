package statement

type Statement interface {
	Type() string
	PredicateType() string
	Predicate() ([]byte, error)
}

type Predicate interface {
	BuildType() string
	Builder() Builder
	Invocation() Invocation
}

type Builder interface {
	ID() string
}

type Invocation interface {
	ConfigSource() ConfigSource
}

type ConfigSource interface {
	URI() string
	Digest() map[string]string
	EntryPoint() string
}
