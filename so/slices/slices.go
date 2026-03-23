// Package slices provides various functions useful with slices of any type.
package slices

import (
	_ "embed"

	"solod.dev/so/mem"
)

//so:embed slices.h
var slices_h string

// Make allocates a slice of type T with given length using allocator a.
// If the allocator is nil, uses the system allocator.
// The returned slice is allocated; the caller owns it.
//
//so:extern
func Make[T any](a mem.Allocator, len int) []T {
	return mem.AllocSlice[T](a, len, len)
}

// MakeCap allocates a slice of type T with given length and capacity using allocator a.
// If the allocator is nil, uses the system allocator.
// The returned slice is allocated; the caller owns it.
//
//so:extern
func MakeCap[T any](a mem.Allocator, len int, cap int) []T {
	return mem.AllocSlice[T](a, len, cap)
}

// Free frees a previously allocated slice.
// If the allocator is nil, uses the system allocator.
//
//so:extern
func Free[T any](a mem.Allocator, s []T) {
	mem.FreeSlice(a, s)
}

// Append appends elements to a heap-allocated slice, growing it if needed.
// If the allocator is nil, uses the system allocator.
// Returns an updated allocated slice; the caller owns it.
//
//so:extern
func Append[T any](a mem.Allocator, s []T, elems ...T) []T {
	return append(s, elems...)
}

// Extend appends all elements from another heap-allocated slice, growing if needed.
// If the allocator is nil, uses the system allocator.
// Returns an updated allocated slice; the caller owns it.
//
//so:extern
func Extend[T any](a mem.Allocator, s []T, other []T) []T {
	return append(s, other...)
}

// Clone returns a shallow copy of the slice.
// If the allocator is nil, uses the system allocator.
// The returned slice is allocated; the caller owns it.
//
//so:extern
func Clone[T any](a mem.Allocator, s []T) []T {
	return append([]T{}, s...)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. Empty and nil slices are considered equal.
//
//so:extern
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
