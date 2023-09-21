package generics

// Contains determines if a slice contains any value passed
func Contains[T comparable](in []T, h T) bool {
	for _, v := range in {
		if h == v {
			return true
		}
	}
	return false
}

// Reduce takes a slice of objects and a selector function and will reduce it
// to an array of member properties.
func Reduce[T interface{}, E any](in []T, f func(T) E) []E {
	out := make([]E, len(in))
	for x, v := range in {
		out[x] = f(v)
	}
	return out
}

// Filter takes a slice and a selector function and will return a slice
// of objects that match the selector function.
func Filter[T any](in []T, f func(T) bool) []T {
	out := make([]T, 0)
	for x, v := range in {
		if f(v) {
			out = append(out, in[x])
		}
	}
	return out
}

// FilterMap takes a map and reduces it to a map where the objects are keyed by
// the return value of a keySelector function and the values are the return value
// of a valueSelector function.
func FilterMap[T, D any, F, E comparable](in map[E]T, keySelector func(E, T) F, valueSelector func(E, T) D) map[F]D {
	out := make(map[F]D, len(in))
	for k, v := range in {
		out[keySelector(k, v)] = valueSelector(k, v)
	}
	return out
}

// Map takes a slice and puts them in a map where the objects are keyed by
// the return value of the selector function.
func Map[T interface{}, E comparable](in []T, f func(T) E) map[E]T {
	out := make(map[E]T, len(in))
	for _, v := range in {
		out[f(v)] = v
	}
	return out
}

// Unique takes a slice of comparables and will return a slice
// of unique elements.
func Unique[T comparable](in []T) []T {
	u := make(map[T]struct{}, 0)
	out := make([]T, 0)
	for _, v := range in {
		if _, ok := u[v]; ok {
			continue
		}
		u[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// Must takes and interface and an error and will panic if the error is not nil
// this is a helper function for open telemetry to create meters globally
func Must[T interface{}](c T, e error) T {
	if e != nil {
		panic(e)
	}
	return c
}
