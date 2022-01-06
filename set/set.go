package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]bool

func (s Set[T]) Add(item T) {
	s[item] = true
}

func (s Set[T]) Del(item T) {
	for k := range s {
		if k == item {
			delete(s, k)
		}
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) String() string {
	items := make([]string, len(s))
	i := 0
	for k := range s {
		items[i] = fmt.Sprint(k)
		i++
	}
	return strings.Join(items, ", ")
}

func New[T comparable]() Set[T] {
	return Set[T]{}
}
