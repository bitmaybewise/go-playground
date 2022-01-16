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

type ImmutableSet[T comparable] map[T]bool

func (s ImmutableSet[T]) String() string {
	items := make([]string, len(s))
	i := 0
	for k := range s {
		items[i] = fmt.Sprint(k)
		i++
	}
	return strings.Join(items, ", ")
}

func (s ImmutableSet[T]) Add(item T) ImmutableSet[T] {
	if _, ok := s[item]; ok {
		return s
	}
	newset := NewImmutable[T]()
	for k := range s {
		newset[k] = true
	}
	newset[item] = true
	return newset
}

func (s ImmutableSet[T]) Del(item T) ImmutableSet[T] {
	newset := NewImmutable[T]()
	for k := range s {
		if k == item {
			continue
		}
		newset[k] = true
	}
	return newset
}

func (s ImmutableSet[T]) Len() int {
	return len(s)
}

func NewImmutable[T comparable]() ImmutableSet[T] {
	return ImmutableSet[T]{}
}
