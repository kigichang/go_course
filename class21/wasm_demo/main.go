package main

import (
	"encoding/base64"
	"fmt"
	"syscall/js"
)

var window = js.Global().Get("window")
var doc = js.Global().Get("document")

func alert(f string, a ...interface{}) {
	window.Call("alert", fmt.Sprintf(f, a...))
}

func hello(_ js.Value, args []js.Value) interface{} {
	return fmt.Sprintf("hello, %s", args[0].String())
}

func main() {
	myfile := doc.Call("querySelector", "#myfile")
	myfile.Call("addEventListener", "change", js.FuncOf(func(_this js.Value, _args []js.Value) interface{} {
		reader := js.Global().Get("FileReader").New()

		reader.Call("addEventListener", "load", js.FuncOf(func(this js.Value, _args []js.Value) interface{} {
			result := this.Get("result")
			srcBuf := js.Global().Get("Uint8Array").New(result)
			size := srcBuf.Length()
			dest := make([]byte, size)
			js.CopyBytesToGo(dest, srcBuf)

			fmt.Println(base64.StdEncoding.EncodeToString(dest))

			return nil
		}))
		reader.Call("readAsArrayBuffer", _this.Get("files").Index(0))
		return nil
	}))

	mybtn := doc.Call("querySelector", "#mybtn")
	mybtn.Call("addEventListener", "click", js.FuncOf(func(_this js.Value, _args []js.Value) interface{} {
		alert("hello, world")
		myfile.Set("disabled", false)
		return nil
	}))

	js.Global().Set("hello", js.FuncOf(hello))
	fmt.Println("hello world!")
	select {}
}
