package message

/*
#include "../c/hello.c"
*/
import "C"

func SayHello(n string) { 
	name := C.CString(n)
	C.SayHello(name)
}

