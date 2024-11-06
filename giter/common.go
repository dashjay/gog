package giter

func FromSlice[T any](in []T) Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < len(in); i++ {
			if !yield(in[i]) {
				break
			}
		}
	}
}
