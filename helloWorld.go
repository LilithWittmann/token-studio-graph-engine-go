package token_studio_graph_engine_go

import "C"

func main() {
	// main() won't be called, but it is required for compilation
}

//export sayHello
func sayHello() *C.char {
	return C.CString("Hello")
	
}
