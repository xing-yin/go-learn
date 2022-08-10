package main

import "C"
import (
	"unsafe"
)

/*
#include <stdio.h>
*/

func main() {
	cstr := C.Cstring("hello c")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
