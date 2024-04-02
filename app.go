package ptr

import "unsafe"

func Atop[T any](addr *T) uintptr {
	return uintptr(unsafe.Pointer(addr))
}

func AtopOffset[T any](addr *T, offset int) uintptr {
	if offset < 0 {
		return uintptr(unsafe.Pointer(addr)) - uintptr(offset)
	}
	return uintptr(unsafe.Pointer(addr)) + uintptr(offset)
}

func Atov[E, T any](addr *T, offset int) *E {
	return Ptoa[E](AtopOffset(addr, offset))
}

func Ptoa[T any](iptr uintptr) *T {
	return (*T)(unsafe.Pointer(iptr))
}

func Ptoas[T []E, E any](iptr uintptr) *T {
	return (*T)(unsafe.Pointer(iptr))
}
