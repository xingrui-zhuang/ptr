package ptr

import "unsafe"

func Atop[T any](addr *T) uintptr {
	return uintptr(unsafe.Pointer(addr))
}

func Ptoa[T any](iptr uintptr) *T {
	return (*T)(unsafe.Pointer(iptr))
}

func Ptoas[T []E, E any](iptr uintptr) *T {
	return (*T)(unsafe.Pointer(iptr))
}
