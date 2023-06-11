package types

import "C"

//export sayHello
func sayHello() *C.char {
	return C.CString("Hello")

}
