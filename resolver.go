package deps

import "errors"

var ErrMissingDependency = errors.New("missing dependency in dependency map")

type Resolver[K comparable, V any] interface {
	Resolve(identifier K) (V, error)
}

type GenericResolver[K comparable, V any] struct {
	deps map[K]V
}

func NewGenericResolver[K comparable, V any](deps map[K]V) *GenericResolver[K, V] {
	return &GenericResolver[K, V]{deps: deps}
}

func (g GenericResolver[K, V]) Resolve(identifier K) (V, error) {
	resolved, ok := g.deps[identifier]
	if !ok {
		return *new(V), ErrMissingDependency
	}

	return resolved, nil
}
