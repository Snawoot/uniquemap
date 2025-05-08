package uniquemap

import (
	"slices"

	"github.com/Snawoot/unicmp"
	"github.com/Snawoot/uniqueslice"
)

type pair[K, V comparable] struct {
	key K
	val V
}

// Make returns a globally unique handle for a map[K]V. Handles
// are equal if and only if the values used to produce them are equal.
// Make is safe for concurrent use by multiple goroutines.
func Make[Map ~map[K]V, K, V comparable](m Map) Handle[Map, K, V] {
	pairs := make([]pair[K, V], len(m))
	idx := 0
	for k, v := range m {
		pairs[idx] = pair[K, V]{k, v}
		idx++
	}
	o := unicmp.ForType[pair[K, V]]()
	slices.SortFunc(pairs, o.Cmp)
	return Handle[Map, K, V]{uniqueslice.Make[[]pair[K, V], pair[K, V]](pairs)}
}

// Handle is a globally unique identity for some map[K]V.
//
// Two handles compare equal exactly if the two values used to create the handles
// would have also compared equal. The comparison of two handles is trivial and
// typically much more efficient than comparing the values used to create them.
type Handle[Map ~map[K]V, K, V comparable] struct {
	h uniqueslice.Handle[[]pair[K, V], pair[K, V]]
}

// Value returns a map of shallow copies of the key-value pairs that produced the Handle.
// Value is safe for concurrent use by multiple goroutines.
func (h Handle[Map, K, V]) Value() Map {
	pairs := h.h.Value()
	m := make(Map, len(pairs))
	for _, pair := range pairs {
		m[pair.key] = pair.val
	}
	return m
}
