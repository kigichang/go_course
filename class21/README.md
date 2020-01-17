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
