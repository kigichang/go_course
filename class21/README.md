# 21 Go WebAssembly

## WebAssembly (WASM) Introduction

### History

- Announced in 2015
- More than 80% browsers support (2018/08):
  - Desktop: 88.82%
  - Mobile: 86.64%

- MVP in 2017
- Develop Group:
  - Mozilla (Firefox)
  - Google (Chrome)
  - Microsoft (Edge)
  - Apple (Safari)

### Characteristics

- **Binary Instruction** Format
- Compilation Of **High-Level Languages**
  - C/C++, Rust
  - C#, Go
- Develop on Web Client and Server.
- Security
  - Memory-safe Sandbox
  - Same-Origin and Permissions Security Policies of Browser
- Support Non-Web Embeddings:
  - IoT devices
  - Desktop/Mobile App.

### WASM and Javascript

- WebAssembly does NOT Replace Javascript.
- WebAssembly enable High Performance application:
  - Image Processing (HTML5 Canvas/2D)
  - 3D (WebGL)
  - Crypto
- Javascript applied on UI/UX as usual
  - DOM/CSS
  - RWD

### WebAssembly in Go

- **Experimental** in Go 1.11
- Pure Go code
- Garbage Collection
- Go Runtime Engine & Package
- package: **syscall/js**

### Data Types from Go to Javascript

| Go                     | JavaScript             |
| ---------------------- | ---------------------- |
| js.Value               | [its value]            |
| js.Func                | function               |
| nil                    | null                   |
| bool                   | boolean                |
| integers and floats    | number                 |
| string                 | string                 |
| []interface{}          | new array              |
| map[string]interface{} | new object             |

### js.Value API Summary

- Get
  - get javascript global object, object constructor, object property and user global variables.
- Set
  - sets object property
- Call
  - call object methods or global functions.
- Truthy
  - returns javascript truthiness values.
  - javascript values are false, 0, “”, null, undefined and NaN return false.
- Index/SetIndex:
  - manipulate iterable or array types.

### js.Global API Summary

- Represents javascript Global or window
- js.Global().Get(“window”): javascript **window** object
- js.Global().Get(“document”): javascript **document** object
- js.Global().Get(“object_constructor”).New: new javascript build-in object
  - js.Global().Get(“websocket”).New(url)

### Others API Summary

- Wrapper interface
  - Implements Wrapper interface when wrapping a javascript object
- FuncOf:
  - Event callback
  - Exporting function for javascript

### Go Concurrency in WASM

Followings are fine.

- Go Routine
  - Multiple Thread
- Channel
  - Communication between Go Routine
- Lock
  - sync.Mutex

### Go WASM works with Javascript

- Javascript TypedArray and Go Slice
  - Share memory in Go 1.12, but remove from Go 1.13 (for some bug issues)
- Javascript UInt8Array and Go []byte
  - js.CopyBytesToGo / js.CopyBytesToJS in Go1.13

### Go WASM Difficulty

- Target WASM file (main.wasm) size: 2.3M ~ 16M
  - encoding/json: ~ 200K
  - net/http: 5 ~ 6M
  - html/template: 2 ~ 3M
  - reflect: ~ 200K
  - Google ProtoBuf: ~ 2.2M
- Downsize Package:
  - Don’t add many features in one package
  - More features, more package imported
- Choose Necessary Go Package
  - fmt
  - Replace with Javascript: using XHTTPRequest, instead of net/http
  - Use some functions only in large package: DIY
- Downsize Target WASM file
  - GZip (Most browsers support): 660K ~ 3.4M
  - Zopfli: 640K ~ 3.3M
  - Brotli: 496K ~ 2.4M
    - Mozilla Firefox 44
    - Google Chrome 50
    - Opera 38
    - Microsoft Edge 15
    - Apple Safari 11
    - cURL 7.57

## Hello World and DOM Manipulation

### 目錄結構

```text
wasm_demo
├── Makefile
├── index.html(`$GOROOT/misc/wasm/wasm_exec.html)
├── wasm_exec.js
└── main.go
```

1. copy `$GOROOT/misc/wasm/wasm_exec.html` and rename to `index.html' into prject folder.
1. copy `$GOROOT/misc/wasm/wasm_exec.js` to project folder.
1. execute `go get -u github.com/shurcooL/goexec` to get **goexec** tool.
    1. add `$GOPATH/bin` to `$PATH`


### WASM Demo

1. execute `make clean; make` in wasm_demo.
1. open `http://127.0.0.1:8080/` in browser (==Google Chrome== preferred)
1. Open **Console** tab in **Developer Tools** to trace log.
1. click **My Button**
1. click `choose file` to open a file and get ==Base64== string in console.

### wasm_demo/index.html

```html
<!doctype html>
<!--
Copyright 2018 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->
<html>

<head>
    <meta charset="utf-8">
    <title>Go wasm</title>
</head>

<body>
    <!--
    Add the following polyfill for Microsoft Edge 17/18 support:
    <script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script>
    (see https://caniuse.com/#feat=textencoder)
    -->
    <script src="wasm_exec.js"></script>
    <script>
        if (!WebAssembly.instantiateStreaming) { // polyfill
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go();
        let mod, inst;
        WebAssembly.instantiateStreaming(fetch("test.wasm"), go.importObject).then((result) => {
            mod = result.module;
            inst = result.instance;
            /*document.getElementById("runButton").disabled = false;*/
            run();
        }).catch((err) => {
            console.error(err);
        });

        async function run() {
            console.clear();
            await go.run(inst);
            inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
        }
    </script>

    <!--<button onClick="run();" id="runButton" disabled>Run</button>-->
    <button id='mybtn'>My Button</button>
    <input type='file' id='myfile' disabled>
</body>

</html>
```

### wasm_demo/main.go

```go {.line-numbers}
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
```

1. `select{}` to block main procedure.
1. `js.Value.Get` to get DOM element property.
1. `js.Value.Set` to set DOM element property.
1. `js.Value.Index` to get value in Array.
1. `js.Value.Call` to invoke a Javascript object method.
1. `js.Value.Get(construtor).New(arguments)` to create a Javascript Object.
1. `js.FuncOf` to make a function, or event handler.
1. convert Javascript Array Buffer with `Uint8Array` view and `js.CopyBytesToGo`.
    - `srcBuf := js.Global().Get("Uint8Array").New(result)`, result is an ==Array Buffer==.
    - `js.CopyBytesToGo(dest, srcBuf)` copy Javascript `Uint8Array` to Go Byte Slice.
1. `js.Global().Set("func_name", js.FuncOf(func))` to export a function for Javascript.
    - `<button onclick='window.alert(hello(this.innerText))'>Click Me</button>` invokes `hello` defined in Go.
