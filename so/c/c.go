// Package c provides convenience helpers for C interop.
// It bridges C's null-terminated strings and raw pointers
// with So's string and slice types.
package c

import "unsafe"

//so:embed c.h
var c_h string

// Alignof returns the alignment of type T in bytes.
//
//	alignof(T)
//
//so:extern
func Alignof[T any]() int {
	var v T
	return int(unsafe.Alignof(v))
}

// Assert aborts the program with the given message
// if the condition is not true.
// If assertions are disabled, does nothing.
//
//	assert((cond) && msg)
//
//so:extern
func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// Bytes wraps a raw byte pointer and length into a []byte without copying.
// If ptr is nil, returns nil.
//
//	(so_Slice){ptr, n, n}
//
//so:extern
func Bytes(ptr *byte, n int) []byte { _, _ = ptr, n; return nil }

// CharPtr casts a *byte (uint8_t*) to char* for C functions
// that expect char* instead of uint8_t*.
//
//	(char*)ptr
//
//so:extern
func CharPtr(ptr *byte) *byte { _ = ptr; return nil }

// Sizeof returns the size of type T in bytes.
//
//	sizeof(T)
//
//so:extern
func Sizeof[T any]() int {
	var v T
	return int(unsafe.Sizeof(v))
}

// Slice wraps a raw pointer and length into a []T without copying.
// If ptr is nil, returns nil.
//
//	(so_Slice){ptr, len, cap}
//
//so:extern
func Slice[T any](ptr *T, len int, cap int) []T {
	if ptr == nil {
		return nil
	}
	s := unsafe.Slice(ptr, cap)
	return s[:len]
}

// String converts a null-terminated C string to a So string.
// If ptr is nil, returns "".
//
//	(so_String){s, strlen(s)}
//
//so:extern
func String(ptr *byte) string { _ = ptr; return "" }

// PtrAdd adds offset bytes to a raw pointer (void*).
//
//	(void*)(ptr + offset)
//
//so:extern
func PtrAdd(ptr any, offset int) any {
	raw := ptrVal(ptr)
	p := unsafe.Add(raw, offset)
	return (*byte)(p)
}

// PtrAs casts a raw pointer (void*) to *T.
//
//	(T*)(ptr)
//
//so:extern
func PtrAs[T any](ptr any) *T {
	raw := ptrVal(ptr)
	return (*T)(raw)
}

// ptrVal extracts a raw pointer from an interface containing any pointer type.
// For testing only; in C, any pointers are void*.
//
//so:extern
func ptrVal(v any) unsafe.Pointer {
	type iface struct {
		_    unsafe.Pointer
		data unsafe.Pointer
	}
	return (*iface)(unsafe.Pointer(&v)).data
}
