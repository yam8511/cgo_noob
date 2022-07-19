package mini

/*
#include <stdio.h>  // <--- puts
#include <stdlib.h> // <--- free
#include <stdint.h> // for uintptr_t

extern char* MyGoPrint(uintptr_t handle);
extern char* Concat(char *a, char *b);
void myprint(uintptr_t handle);
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export PrintCString
func PrintCString(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)

	h := cgo.NewHandle(s)
	C.myprint(C.uintptr_t(h))
	h.Delete()
}

//export PrintChar
func PrintChar(cc *C.char) {
	s := C.GoString((*C.char)(cc))
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)

	h := cgo.NewHandle(s)
	C.myprint(C.uintptr_t(h))
	h.Delete()
}

//export MyGoPrint
func MyGoPrint(handle C.uintptr_t) *C.char {
	h := cgo.Handle(handle)
	val := h.Value().(string)
	return C.CString("Go call -> " + val)
	// println()
	// h.Delete()
}

//export Concat
func Concat(a, b *C.char) *C.char {
	// C *char to Go string
	aa := C.GoString((*C.char)(a))
	bb := C.GoString((*C.char)(b))
	return C.CString(aa + " + " + bb)
}
