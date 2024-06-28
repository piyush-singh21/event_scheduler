package swag

import "unsafe"

// hackStringBytes returns the (unsafe) underlying bytes slice of a string.

func StringData(s string) *byte {
	// Convert the string to a pointer to its underlying data
	return (*byte)(unsafe.Pointer(&s))
}

func hackStringBytes(str string) []byte {
	return unsafe.Slice(StringData(str), len(str))
}
