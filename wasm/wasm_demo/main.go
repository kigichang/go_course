//go:build js && wasm
// +build js,wasm

/*
*
此範例程式，需要 https://github.com/shurcooL/goexec.
請使用 make goexec 或 go generate 安裝
*
*/
package main

//go:generate go install github.com/shurcooL/goexec

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

func isNaN(v js.Value) bool {
	return js.Global().Call("isNaN", v).Bool()
}

func parseInt(val string, radix int) (int, bool) {
	x := js.Global().Call("parseInt", val, radix)
	if isNaN(x) {
		return 0, false
	}
	return x.Int(), true
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
	fmt.Println(parseInt("abc", 10))
	fmt.Println("hello world!")
	select {}
}
