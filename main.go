package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("hello", js.FuncOf(hello))

	js.Global().Get("console").Call("log", "wasm loaded")
	select {}
}

func hello(this js.Value, args []js.Value) interface{} {
	js.Global().Get("console").Call("log", "hello world")
	return nil
}
