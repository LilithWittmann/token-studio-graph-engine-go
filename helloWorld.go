package main

import "C"

func main() {
	// main() won't be called, but it is required for compilation
}

//export sayHello
func sayHello() *C.char {
	return C.CString("Hello")

}
