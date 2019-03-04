# Cowork With C/C++

在 Golang 有提供 [cgo](https://golang.org/cmd/cgo/#hdr-Go_references_to_C) 與 C 的程式互動。

Go 使用 C:

1. 在 Go 的程式碼檔案 (*.go) 撰寫 C 程式。
1. C 提供 libary 或程式碼，讓 Go 使用。

C 使用 Go:

1. Go 編繹成 C 的 library，讓 C 使用。

因為 Golang 本身沒有 OOP 的設計，Go 無法直接使用 C++ 有關 OOP 程式，如果需要使用，則必須在 C++ 的 OOP 程式，再封裝一層 C 程式來使用。目前 Go 有支援 swig 套件，可以協助封裝 C/C++ 程式。

Go 預設使用的編繹器 **gc**，如果遇到官方不支援 OS 或者 CPU，則需要下載 Go 原始碼，使用 **gccgo** 來編繹。

## unsafe

如果使用到有關 C 的 pointer 時，就會用到 Go 的 **unsafe** package。因為 Go 的 unsafe 套件使用到系統底層的特性，所以會失去相容與移植。如非必要，儘可能不要使用。
  
unsafe package 主要有三個 function 與一個 data type
  
- functions:
  - func Alignof(x ArbitraryType) uintptr: memory alignment
    > The unsafe.Alignof function reports the required alignment of its argument’s type. Like Sizeof, it may be applied to an expression of any type, and it yields a constant. Typically, boolean and numeric types are aligned to their size (up to a maximum of 8 bytes) and all other types are word-aligned. (where a **word** is **4** bytes on a **32-bit** platform and **8** bytes on a **64-bit** platform)
  
  - func Offsetof(x ArbitraryType) uintptr: offset from start for struct and array
  - func Sizeof(x ArbitraryType) uintptr: size of for type
- data type:
  - type Pointer: 等同 void* in c. [Go reference to C](https://golang.org/cmd/cgo/#hdr-Go_references_to_C )
     > The C type void* is represented by Go's unsafe.Pointer
  
### Sizeof and Alignof

```go
package main
  
import (
    "fmt"
    "unsafe"
)
  
func main() {
    var x struct {
        a bool
        b float64
        c int16
    }
  
    fmt.Println("Sizeof x:", unsafe.Sizeof(x))
    fmt.Println("Alignof x:", unsafe.Alignof(x))
  
    fmt.Println("Sizeof x.a:", unsafe.Sizeof(x.a), "AlignOf x.a:", unsafe.Alignof(x.a), "Offsetof x.a:", unsafe.Offsetof(x.a))
    fmt.Println("Sizeof x.b:", unsafe.Sizeof(x.b), "AlignOf x.b:", unsafe.Alignof(x.b), "Offsetof x.b:", unsafe.Offsetof(x.b))
    fmt.Println("Sizeof x.c:", unsafe.Sizeof(x.c), "AlignOf x.c:", unsafe.Alignof(x.c), "Offsetof x.c:", unsafe.Offsetof(x.c))
  
    var y struct {
        a float64
        b int16
        c bool
    }
  
    fmt.Println("Sizeof y:", unsafe.Sizeof(y))
    fmt.Println("Alignof y:", unsafe.Alignof(y))
  
    fmt.Println("Sizeof y.a:", unsafe.Sizeof(y.a), "AlignOf y.a:", unsafe.Alignof(y.a), "Offsetof y.a:", unsafe.Offsetof(y.a))
    fmt.Println("Sizeof y.b:", unsafe.Sizeof(y.b), "AlignOf y.b:", unsafe.Alignof(y.b), "Offsetof y.b:", unsafe.Offsetof(y.b))
    fmt.Println("Sizeof y.c:", unsafe.Sizeof(y.c), "AlignOf y.c:", unsafe.Alignof(y.c), "Offsetof y.c:", unsafe.Offsetof(y.c))
  
    var z struct {
        a bool
        b int16
        c float64
    }
  
    fmt.Println("Sizeof z:", unsafe.Sizeof(z))
    fmt.Println("Alignof z:", unsafe.Alignof(z))
  
    fmt.Println("Sizeof z.a:", unsafe.Sizeof(z.a), "AlignOf y.a:", unsafe.Alignof(z.a), "Offsetof z.a:", unsafe.Offsetof(z.a))
    fmt.Println("Sizeof z.b:", unsafe.Sizeof(z.b), "AlignOf y.b:", unsafe.Alignof(z.b), "Offsetof z.b:", unsafe.Offsetof(z.b))
    fmt.Println("Sizeof z.c:", unsafe.Sizeof(z.c), "AlignOf y.c:", unsafe.Alignof(z.c), "Offsetof z.c:", unsafe.Offsetof(z.c))
}
```

Output:

```text
Sizeof x: 24
Alignof x: 8
Sizeof x.a: 1 AlignOf x.a: 1 Offsetof x.a: 0
Sizeof x.b: 8 AlignOf x.b: 8 Offsetof x.b: 8
Sizeof x.c: 2 AlignOf x.c: 2 Offsetof x.c: 16
Sizeof y: 16
Alignof y: 8
Sizeof y.a: 8 AlignOf y.a: 8 Offsetof y.a: 0
Sizeof y.b: 2 AlignOf y.b: 2 Offsetof y.b: 8
Sizeof y.c: 1 AlignOf y.c: 1 Offsetof y.c: 10
Sizeof z: 16
Alignof z: 8
Sizeof z.a: 1 AlignOf y.a: 1 Offsetof z.a: 0
Sizeof z.b: 2 AlignOf y.b: 2 Offsetof z.b: 2
Sizeof z.c: 8 AlignOf y.c: 8 Offsetof z.c: 8
```

x, y, z 在 64-bit 系統下， alignment 都是 8 bytes. 但 x 的 size 是 24 bytes (3 words[^word])，而其他 y, z 都是 16 bytes (2 words)。
主要因為 x 的 bool (x.a) 與 int16 (x.c)中間是 float64 (x.b)，佔了 8 bytes (1 word)，bool 雖只佔 1 byte，但要補足成 8 bytes (1 word), 同理 x.c 只佔 2 bytes，也要補足成 8 bytes (1 word)。
而 y, z 因為 bool, int16 是相連，因此在 bool 後面補 1 bytes, int16 補 4 bytes，補足成 8 bytes(1 word)。所以 x 是 24 bytes，而 y, z 是 16 bytes。[^src_sample]
  
[^word]: 在 32 bit 系統下，1 word = 4 bytes (32bit), 64 bit 是 8 bytes (64bit)
[^src_sample]: [unsafe.Sizeof, Alignof 和 Offsetof](https://wizardforcel.gitbooks.io/gopl-zh/ch13/ch13-01.html )
  
### unsafe.Pointer

unsafe.Pointer 可以是任意型別的指標。在 Golang 的 strong type 安全機制下，不同的資料型別與指標都不可以直接轉換，如：
  
- 不同指標的值，即使是相同 bit 數，如 int64 和 float64。
- 指標 與 uintptr 的值。
  
unsafe.Pointer 可以破壞 Go 的安全機制；unsafe.Pointer 的功能有：
  
1. 任何資料型別的指標，都可以轉成 unsafe.Pointer。
1. unsafe.Pointer 也可轉成任何資料型別的指標。
1. unsafe.Pointer 可以轉成 uintptr，之後可以利用 uintptr 做指標位移。
1. uintptr 可以轉成 unsafe.Pointer，但有風險。
  
unsafe.Pointer 一定要遵守[官網](https://golang.org/pkg/unsafe/#Pointer )提到的使用模式，官網有提到，如果沒有遵守的話，現在雖然沒問題，但在未來不一定正確。
  
1. Conversion of a *T1 to Pointer to *T2，前提 T2 的所需的記憶體空間，要小於或等於 T1。如果 T2 > T1 的話，有可能在操作 T2 時，會造成溢位。

    ```go
    package main
  
    import (
        "fmt"
        "unsafe"
    )
  
    func main() {
  
        f := 3.1415926
        t1 := *(*uint64)(unsafe.Pointer(&f))
        fmt.Printf("%v, %T\n", t1, t1)
  
        t2 := *(*uint32)(unsafe.Pointer(&f))
        fmt.Printf("%v, %T\n", t2, t2)
  
        var f1 float32 = 3.1415926
        t3 := (*uint64)(unsafe.Pointer(&f1))
        fmt.Printf("%v, %v, %T\n", t3, *t3, t3)
        fmt.Printf("%v, %v, %T\n", &f1, f1, f1)
  
        *t3 = 1<<63 - 1
        fmt.Printf("%v, %T\n", *t3, t3)
        fmt.Printf("%v, %T\n", f1, f1)
    }
    ```

    (以上 t3 是 *uint64，大於 float32，在最後造成 float32 溢位)
  
1. Conversion of a Pointer to a uintptr (but not back to Pointer)
  
    >Even if a uintptr holds the address of some object, the garbage collector will not update that uintptr's value if the object moves, nor will that uintptr keep the object from being reclaimed.
  
    官網文件提到，Pointer 可以轉 uintptr，但反之 uintptr 轉 Pointer 不行。因為 uintptr 只是存當時的記憶體位址，但實體 (instance) 有可能已經被 GC (Garbage collection)，因此轉成 Pointer 操作後，會有問題。下面的例子中，Go 的工具會幫忙檢查。(go vet)

    ```go
    package main
  
    import (
        "fmt"
        "unsafe"
    )
  
    func main() {
        a := 10
  
        ap1 := uintptr(unsafe.Pointer(&a))
  
        fmt.Printf("%x, %T\n", ap1, ap1)
  
        ap2 := (*int)(unsafe.Pointer(ap1)) // warning: possible misuse of unsafe.Pointer
  
        fmt.Printf("%x, %v, %T\n", ap2, *ap2, ap2)
    }
    ```

1. Conversion of a Pointer to a uintptr and back, with arithmetic.
  
    指標位移

    ```go
    package main
  
    import (
        "fmt"
        "unsafe"
    )
  
    type test struct {
        name string
        age  int
        addr string
    }
  
    func main() {
  
        slice := []int{1, 2, 3, 4, 5}
  
        snd := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + 1*unsafe.Sizeof(slice[0])))
        fmt.Printf("%v, %T\n", snd, snd)
  
        t := test{"A", 10, "ABC"}
  
        addr := *(*string)(unsafe.Pointer((uintptr(unsafe.Pointer(&t)) + unsafe.Offsetof(t.addr))))
        fmt.Printf("%v, %T\n", addr, addr)
    }
    ```

    以下的操作方式是**不正確**的，雖然是可以執行。
  
    1. 取資料的結束位址
    1. 先用變數保留 uintptr 再轉回 Pointer

        ```go
        package main

        import (
            "fmt"
            "unsafe"
        )

        type test struct {
            name string
            age  int
            addr string
        }

        func main() {
            slice := []int{1, 2, 3, 4, 5}

            // 取結束位址
            out := unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + uintptr(len(slice)))
            fmt.Println(out)

            // 暫存 uintptr, 再轉成 unsafe.Pointer()
            firstp := uintptr(unsafe.Pointer(&slice[0]))
            sndp := firstp + unsafe.Sizeof(slice[0])

            snd := *(*int)(unsafe.Pointer(sndp)) // warning: possible misuse of unsafe.Pointer
            fmt.Printf("%v, %T\n", snd, snd)

            t := test{"A", 10, "ABC"}

            // 取結束位址
            end := unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + unsafe.Sizeof(t))
            fmt.Println(end)
        }
        ```

1. Conversion of a Pointer to a uintptr when calling syscall.Syscall.

    > If a pointer argument must be converted to uintptr for use as an argument, that conversion must appear in the call expression itself:

    ex:

    ```go
    syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
    ```

    invalid:

    ```go
    // INVALID: uintptr cannot be stored in variable
    // before implicit conversion back to Pointer during system call.
    u := uintptr(unsafe.Pointer(p))
    syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n)
    ```

1. Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr from uintptr to Pointer.

    ```go
    package main

    import (
        "fmt"
        "reflect"
        "unsafe"
    )

    type test struct {
        name string
        age  int
        addr string
    }

    func main() {
        u := reflect.ValueOf(new(int)).Pointer()
        p := (*int)(unsafe.Pointer(u))
        fmt.Printf("%v, %v, %T\n", p, *p, p)
        *p = 200
        fmt.Printf("%v, %v, %T\n", p, *p, p)
    }
    ```

1. Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer.

    1. 不可以 assign 值至 Data

        ```go
        var s string
        hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
        hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
        hdr.Len = n
        ```

    1. 使用 SliceHeader or StringHeader 時，請用 pointer 形式

        ```go
        // INVALID: a directly-declared header will not hold Data as a reference.
        var hdr reflect.StringHeader
        hdr.Data = uintptr(unsafe.Pointer(p))
        hdr.Len = n
        s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost
        ```

## cgo

使用 **cgo** 時，需要在 go 的程式碼中加 `import "C"` 這個 **pseudo-package**，即可使用 C 的型別與函式，如 `C.size_t` 或 `C.free`。在 `import "C"` **緊接**這一行上面(不能留空行)，可再加入 C 的程式碼或者 **cgo** 的編繹參數。

### 基本用法

```go
package main

/*
#include "stdio.h"
#include "stdlib.h"

void hello(char* name) {
    printf("hello %s\r\n", name);
}
*/
import "C"

import (
    "unsafe"
)

func main() {
    cs := C.CString("cyberon")
    C.hello(cs)
    C.free(unsafe.Pointer(cs))
}
```

### 基本型別互轉

C                  | CGo             | Go
:----------------: | :-------------: | :-:
char               | C.char          | int8
unsigned char      | C.uchar         | uint8
short int          | C.short         | int16
unsigned short     | C.ushort        | uint16
int                | C.int           | int (os dependent)
unsigned int       | C.uint          | uint (os dependent)
long               | C.long          | int32
unsigned long      | C.ulong         | uint32
long long          | C.longlong      | int64
unsigned long long | C.ulonglong     | uint64
float              | C.float         | float32
double             | C.double        | float64
float complex      | C.complexfloat  | complex64
double complex     | C.complexdouble | complex128
void*              | &nbsp;          | unsafe.Pointer

### C 與 CGo 的記憶體管理

在 CGo 的 pseudo-package **"C"** 有[內建幾個函式](https://golang.org/cmd/cgo/#hdr-Go_references_to_C)，可以互轉 Go 與 C 的資料型別, 在 Go 轉成 C 的型別時，將 Go 的資料 clone 一份給 C 用，因此在使用完畢後，要記得 call C 的 `free` 將記憶體釋放。

- func C.CString(string) *C.char
    Go string 轉 C 的 char*，**需要用 free**。

- func C.CBytes([]byte) unsafe.Pointer
    Go byte slice 轉 C 的 void*, **需要用 free**。

- func C.GoString(*C.char) string
    C 的 char* 轉 Go string

- func C.GoStringN(*C.char, C.int) string
    C 的 char*，並指定長度， 轉 Go string

- func C.GoBytes(unsafe.Pointer, C.int) []byte
    C 的 void*，並指定長度，轉成 Go byte slice.

### Swig

Swig 可以將 C/C++ 與其他高階語言(eg: PHP, Java, C#，Go) 等結合的工具，其原理是將 C/C++ 再用對應的程式語言封裝。Go 的 build 工具已內建整合 Swig。

#### Swig Sample

專案目錄 **$GOPATH/src/go_course/classxy/swig**

```text
.
├── foo
│   ├── foo.cpp
│   ├── foo.go
│   ├── foo.hpp
│   └── foo.swigcxx
└── main.go
```

**compile**:
  
在專案目錄 `$GOPATH/src/go_test/class16/swig` 下，執行 `go build -x -work`, `-x` 是列印所有編譯的指令，`-work` 是列印出中繼的產檔目錄。在所有指令中，`go build` 會自動去執行 `swig` 指令: `swig -go -cgo -intgosize 64 -module foo -o $WORK/b023/foo_wrap.cxx -outdir $WORK/b023/ -C++ foo.swigcxx`。
  
想了解 swig 怎麼封裝，可以到 `$WORK` 目錄下，看產生出來的中斷檔 `_foo_swig.go`
  
**說明**:
  
- `foo/foo.hpp` and `foo/foo.cpp`: C++ 程式，有一組 class, 一個 global 變數，二個 global functions
- `foo/foo.go`: 空白的 go 程式，避免在 build 程式時，`go build` 發生 `foo` package 沒有 go 的程式
- `foo/foo.swigcxx`: swig 的 interface 檔案。
    1. 如果是要連 C++ 程式，副檔名要用 `.swigcxx`，如果是 C 程式，副檔名 `.swig`
    1. interface 的檔名，要與 package 名稱相同
    1. 因為是 go build 會自定 module 名稱，因此檔案中的 `module myfoo` 就無用
    1. interface 的寫法，就把 C++ 的 header 放入
    1. 如果想看中繼檔，可到 `$WORK` 目錄。eg: `/var/folders/5t/jkthvjgn0gxc7k98wczh5k1c0000gn/T/go-build519318596/b023/`
- swig 封裝原則
    1. C++ 的 class 會轉成 Go interface
    1. class 的 public member data 會產生對應的 Get 及 Set Functions.
    1. C++ 的 new/delete class, 會有 NewClass 及 DeleteClass 的對應
    1. Global 變數，也會有 Get/Set 的對應。
    1. Global function 會改成第一個字母大寫的 function.

### 自行封裝

專案目錄 **$GOPATH/src/go_course/classxy/foo2**

```text
.
├── foo
│   ├── Makefile
│   ├── cfoo.cpp
│   ├── cfoo.h
│   ├── foo.cpp
│   └── foo.hpp
└── main.go
```

**說明**:

1. foo.cpp, foo.h C/C++ 程式碼: 有兩組 Class 定義，接下來就是要封裝這兩個 Class，讓 Go 使用。
1. cfoo.h, cfoo.cpp: 封裝 foo 裏面的兩個 Class。
1. main.go: 測試程式。

### Generate C library from Go

以 **cicd.icu/cyberon/mrcp/stt** 為例。

1. 請將 Go 的專案做成是應用程式的模式，也就是說需要一個 main.go 有 `func main()`，function body 可以留空白；以及 package 是 `package main`。
1. 在 Go 的源碼中，將要匯出給 C 用的 function，緊接在宣告行的上方，加入 `//export FUNC_NAME` (`//` 與 `export` 間，**不要留空白**)，eg:

    ```go
    //export Init
    func Init() {
        ...
    }
    ```

1. 在 compile 時，請加入 `-o` 來指定輸出的檔案
1. 在 compile 時，請如入 `-buildmode=c-xxxx`，來指定是要輸入動態或靜態連結 library.
    - c-archive: 靜態連結
    - c-shared: 動態連結
1. compile 完成後，會自動產生 **.h** 檔案，再拿產出的 **.a** 或 **.so**，給 C 程式使用。

### CGo compile and link 參數

環境變數：`CGO_CFLAGS`, `CGO_CPPFLAGS`, `CGO_CXXFLAGS`, `CGO_FFLAGS` and `CGO_LDFLAGS`。可以在 Go 源碼中的 `import "C"` 緊接上方，以註解的方式，修改 CGo 的環境變數，甚至可以指定平台。

eg:

```go
#cgo darwin LDFLAGS:  -L. -lCybSttDnn -lCybServerDnnRecognizer
#cgo linux LDFLAGS:  -L. -lCybSttDnn -lCybServerDnnRecognizer /xxx/libfst.a -ldl /usr/lib/libatlas.so.3 /usr/lib/libf77blas.so.3 /usr/lib/libcblas.so.3 /usr/lib/liblapack_atlas.so.3 -lm -lpthread -ldl -lglib-2.0 -lgmodule-2.0
```

因為 security 因素，請在 compile 時，修改 CGo 的環境變數 `CGO_LDFLAGS_ALLOW`, eg: `env CGO_LDFLAGS_ALLOW='.*\.so.[0-9]$'`