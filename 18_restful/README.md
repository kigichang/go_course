# 18 RESTful

## REST API

### What is REST

Representational State Transfer (REST) 由 Roy Thomas Fielding 在 2000 年在 [Architectural Styles and
the Design of Network-based Software Architectures](https://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm) 這篇論文中提出。

REST 是一種設計風格，並非一種標準。通常使用 HTTP, URI, XML, HTML 來實作，近來 XML 的部分，改用了 JSON，或 XML 與 JSON 兩者並存。

### Architectural Goals

#### Performance

#### Scalability

- scaling up - increasing the capacity of services, consumers, and network devices
- scaling out - distributing load across services and programs

#### Simplicity

#### Modifiability

可以應需求改變，而可輕易更動

#### Visibility

客戶與伺服器端，或伺服器之間，可加入 middleware 來監控或調適彼此間的溝通

#### Portability

服務可以輕易地被佈暑

#### Reliability

### Architectural Constrains

#### Client-Server

- 滿足 RESTful 設計的系統，需要是 Client-Server 的架構 (客戶端與伺服器端)
- RESTful 只規範通訊協定

![Client-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/client_server_style.gif)

#### Layered System

- 客戶端無法指定每次要連線的伺服器，也無法指定伺服器間的路由

![Uniform-Layered-Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/layered_uccss.gif)

#### Cache

- 系統內的所有資料，都需定義是否能使用快取 (cache)
- 客戶端能建立並使用本地的快取，如伺服端沒有更動，客戶端則使用本地的快取資料
- 可以節省伺服器端的資源及網路頻寬
- 通常不變的資料，適用快取，但也不是所有的資料都合適
- Cache 也能應用在伺服器端

![Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/ccss_style.gif)

#### Stateless

- RESTful 規範客戶端與伺服器端的溝通是 Stateless。但伺服器以後的設計則沒有限制，也就是說伺服器端以後的部分可以使用像 memcached, Redis 等來存放狀態(如 session)資訊
- 由於 RESTful 是 Stateless, 在 API 的設計上，需要符合 Atomic 特性。簡單來說，atomic 是指不能用二個以上的 API 來完成一個動作。如轉帳，需要扣 A 帳號加 B 帳號，在設計時，就不要拆成兩個 API 以免發生扣 A 但沒加 B。
- 每次 API 的呼叫，都要傳足夠且完整的資料給伺服器端，不要認為伺服器端會記錄現在的使用者資訊。
- Idempotent 特性，簡單來說，客戶端可以重覆呼叫 API，在伺服器端要可以處理這種情境。一般說來，如果是讀取未被更動的資料，每次的呼叫，回傳結果要為一致；但更新與刪除資料則不一定。因此在設計 RESTful API 時，要將重覆呼叫的因素考量進去。

![Client-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/stateless_cs.gif)

#### Uniform Interface

- 介面上定義應該要基於資源，而不是動作
- 資源通常有四種動作：新增(Create)，讀取(Read)，更新(Update)，刪除(Delete) (CRUD)

![Uniform-Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/uniform_ccss.gif)

#### Code-On-Demand (Optional)

簡單來說，伺服器端可以回傳一個指定或一段程式碼 (如: javascript) 讓客戶端來執行

![Restful](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_style.gif)

### RESTful Interface Example

#### Use HTTP Response Codes to Indicate Status

常用的 10 種 HTTP Status Code

- 200 OK

  最常用回覆成功的狀態碼

- 201 CREATED

  用在新增資料成功 (透過 PUT or POST 方法)。也可以在 Header 的 Location 加入新資料的連結

- 204 NO CONTENT

  回覆處理成功，但不會回覆訊息，常用在 DELETE and PUT

- 400 BAD REQUEST

  表示請求錯誤。大都情形下的錯誤，都可以回覆此狀態碼。

- 401 UNAUTHORIZED

  認証失敗

- 403 FORBIDDEN

  未授權

- 404 NOT FOUND

  資源不存在

- 405 METHOD NOT ALLOWED

  存取資源的方法不支援，比如說： POST _/users/123_，在新增用戶時，不能指定用戶的編號，此時就可以回覆 405

- 409 CONFLICT

  更新資源發生衝突，比如說：在用戶資料中，假設 email 是唯一值，當有兩筆用戶資料，填入同一個 email 時，就可以回覆 409。

- 500 INTERNAL SERVER ERROR

  發生異外錯誤時，通常都是系統發生 Exception 時，回覆 500

#### Using HTTP Methods for RESTful Services

- POST

  用於新增資源；新增成功時，可以回覆 201，並且在 Header 的 Location 回傳新資源的連結 (內含新資源的 ID, 比如說：_/users/1234_)

  Examples:

  - POST [http://www.example.com/customers]([http://www.example.com/customers)
  - POST [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)

- GET

  用於讀取資源，且在資源未更新前，每次讀取，回覆的資料都應一致(Idempotent)。使用 GET 時，不要去新增/更新資源。

  Examples:

  - GET [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - GET [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)
  - GET [http://www.example.com/buckets/sample](http://www.example.com/buckets/sample)

- PUT

  通常用於更新資料，但也可以用在新增。與 POST 不同的是，POST 在新增時，Client 不會指定要新增的 ID，但 PUT 會指定新增的 ID。

  Examples:

  - PUT [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - PUT [http://www.example.com/customers/12345/orders/98765](http://www.example.com/customers/12345/orders/98765)
  - PUT [http://www.example.com/buckets/secret_stuff](http://www.example.com/buckets/secret_stuff)

- DELETE

  用在刪除資料。

  Examples:

  - DELETE [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - DELETE [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)
  - DELETE [http://www.example.com/bucket/sample](http://www.example.com/bucket/sample)

- PATCH

  用於修改資料，與 PUT 不同的是， PUT 需要傳入完整的資料，但 PATCH 只要傳入要修改的部分。

  Examples:

  - PATCH [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - PATCH [http://www.example.com/customers/12345/orders/98765](http://www.example.com/customers/12345/orders/98765)
  - PATCH [http://www.example.com/buckets/secret_stuff](http://www.example.com/buckets/secret_stuff)

#### Summary

HTTP Verb | CRUD | Entire Collection (e.g. /customers) | Specific Item (e.g. /customers/{id})
:--------:| :---: | ----------------------------------- | ------------------------------
POST | Create | 201 (Created), 'Location' header with link to /customers/{id} containing new ID. | 404 (Not Found), 409 (Conflict) if resource already exists..
GET | Read | 200 (OK), list of customers. Use pagination, sorting and filtering to navigate big lists. | 200 (OK), single customer. 404 (Not Found), if ID not found or invalid.
PUT | Update/Replace | 404 (Not Found), unless you want to update/replace every resource in the entire collection. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
DELETE | Delete | 404 (Not Found), unless you want to delete the whole collection—not often desirable. | 200 (OK). 404 (Not Found), if ID not found or invalid.
PATCH | Update/Modify | 404 (Not Found), unless you want to modify the collection itself. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.

### Test

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type category struct {
    ID     uint64 `json:"id,omitempty"`
    Name   string `json:"name,omitempty"`
    Parent uint64 `json:"parent,omitempty"`
}

var categories = make(map[uint64]*category)

func list(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    lst := make([]*category, 0, len(categories))

    for _, v := range categories {
        lst = append(lst, v)
    }

    dataBytes, err := json.Marshal(lst)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Write(dataBytes)
}

func find(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    category, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }

    dataBytes, err := json.Marshal(category)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Write(dataBytes)

}

func add(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    category := new(category)

    dataBytes, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    err = json.Unmarshal(dataBytes, category)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    id := uint64(len(categories) + 1)
    category.ID = id

    categories[id] = category

    w.Header().Add("Location", fmt.Sprintf("/categories/%d", id))
}

func update(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    _, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }

    category := new(category)

    dataBytes, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    err = json.Unmarshal(dataBytes, category)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    category.ID = id
    categories[id] = category

    w.WriteHeader(204)
}

func del(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    _, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }
    delete(categories, id)
    w.WriteHeader(204)

}

func main() {

    categories[1] = &category{
        ID:     1,
        Name:   "3C",
        Parent: 0,
    }

    r := mux.NewRouter()

    r.HandleFunc("/categories", list).Methods("GET")
    r.HandleFunc("/categories", add).Methods("POST")
    r.HandleFunc("/categories/{id:[0-9]+}", find).Methods("GET")
    r.HandleFunc("/categories/{id:[0-9]+}", update).Methods("PUT")
    r.HandleFunc("/categories/{id:[0-9]+}", del).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
}
```

- List all categories:  `curl --include http://localhost:8080/categories`
- Get one category:
  - success: `curl --include http://localhost:8080/categories/1`
  - failure: `curl --include http://localhost:8080/categories/100`
- Add a category: `curl --include --request POST --header "Content-type: application/json" --data '{"name": "PC", "parent": 1}' http://localhost:8080/categories`
- Update a category: `curl --include --request PUT --header "Content-type: application/json" --data '{"name": "NB-2", "parent": 1}' http://localhost:8080/categories/2`
- Delete a category: `curl --include --request DELETE http://localhost:8080/categories/2`

## gRPC (need updated)

ProtoBuf 是 Google 開發的工具，主要來取代 JSON, 與 XML，通常會用在 RPC (Remote Procedure Call) 上，也因此 ProtoBuf 會撘配 Google 開發的 gRPC 使用。ProtoBuf 本身支援多種常用的程式語言，也因此可以利用 ProtoBuf 當作中介的橋樑，在不同的程式語言間，交換資料。

相關資料：

- [Protocol Buffers (ProtoBuf) 官網](https://developers.google.com/protocol-buffers/)
- [Developer Guide](https://developers.google.com/protocol-buffers/docs/overview)
- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)

### protoc

**protoc** 是 Protobuf 的工具，主要是將 protobuf 的定義檔 (.proto) 轉成對應的程式語言。使用 protoc 時，要留意專案的目錄結構。以本例來說，專案 package name 是 `grpc_test`，專案的目錄名稱也設定成 `grpc_test`。

1. 到 [protoc release](https://github.com/google/protobuf/releases) 下載對應作業系統 (Linux, OSX, Win32) 的執行檔。
1. 執行 `go get -u github.com/golang/protobuf/protoc-gen-go` 下載 protoc 的 go plugin。
1. 執行 `go get -u google.golang.org/grpc`
1. 執行 `go get -u github.com/golang/protobuf`

### Package 與目錄結構

由於 Go package 會和目錄結構相關，因此在撰寫專案時，要留意 package name 與目錄。本例測試專案 package name 為 `grpc_test` (見 go.mod)，因此目錄名稱也命名為 `grpc_test`。

```text
class18/grpc_test
├── go.mod
├── go.sum
├── client
│   └── main.go
├── server
│   └── main.go
├── protos
│   ├── test.go
│   ├── test.pb.go
│   └── test.proto
└── service
    ├── service.pb.go
    └── service.proto
```

### 撰寫 .proto

使用 protobuf 前，我們需要先定義資料格式，寫起來有點像在寫 struct。首先在專案目錄下，開一個目錄，如: `protos`，在 `protos` 下還可以依功能再細分。

#### protos/test.proto

```protobuf
syntax = "proto3";

package protos;

import "github.com/golang/protobuf/ptypes/timestamp/timestamp.proto";

option go_package = "grpc_test/protos";

message Hello {
  string name = 1;
  google.protobuf.Timestamp time = 99;
}
```

#### 組成元素

1. syntax: `syntax = "proto3";` 指定 protobuf 的版本，目前有 proto2 與 proto3。建議用 proto3.
1. package: 定義程式的 package, eg: `package protos;`
1. import: 如果有用到其他的 protobuf 資料型別，一樣需要 import, eg: `import "github.com/golang/protobuf/ptypes/timestamp/timestamp.proto";`

    - 由於使用 `google.protobuf.Timestamp`，必須指定來源。protoc 需要加入 `-I` 來指定搜尋 *.proto 的 路徑。
    - **google.protobuf** 是定義在 **github.com/golang/protobuf/ptypes/timestamp/timestamp.proto** 內的 `package google.protobuf;`，因此使用 google timestamp 的資料型別會是 `google.protobuf.Timestamp`。
1. message: 定義資料結構 `message 資料名稱`
1. `option go_package = "grpc_test/protos";` 可以指定對應的程式語言，package name 為何。像 GO, Java 的 package name 也會對應到目錄結構，在產出 `*.pb.go` 時，會使用此設定，建立相對應的目錄。

#### Protocol Buffers Coding Style

- message naming: CamelCase with an initial capital, eg: `message Hello`
- field naming: underscore_separated_names, eg: `required string song_name = 1;`
- Enums:
  - enums naming: CamelCase with an initial capital
  - enum value naming: CAPITALS_WITH_UNDERSCORES

  ```protobuf
  enum Foo {
    FIRST_VALUE = 0;
    SECOND_VALUE = 1;
  }
  ```

#### 資料型別

[proto3](https://developers.google.com/protocol-buffers/docs/proto3)

#### message Hello

```protobuf
message Hello {
  string name = 1;
  google.protobuf.Timestamp time = 99;
}
```

1. 每一組欄位定義後面都會有個數字。eg：`string name = 1;`。
1. 這個數字是指這個欄位的流水號，有點像資料庫的 primary key 的流水號。因此定義之後，不能再異動這個欄位的資料型別，否則會有相容性的問題。
1. 但可以移除這個欄位。如果有需要異動時，應該是再往下加流水號。
1. 在相容性上，如果傳來的資料，缺少欄位的資料時，protobuf 會改成帶該欄位的 zero value。

### 轉成 Go 程式

1. 目錄切到 `class18`
1. 執行 `protoc -I grpc_test/protos -I $GOPATH/src --go_out=. grpc_test/protos/*.proto`
    - `-I` 類似 C 的 include，指定路徑，讓 protoc 去找尋相依的 protobuf 檔案。
    - 放 proto 檔案的目錄，也必須加到 `-I`。eg: `-I grpc_test/protos`
    - `--go_out` 是指要輸出 Go 的程式，並指定目標目錄。
    - 由於在 test.proto 有指定 Go 的 package name `grpc_test/protos`，因此輸出的檔案，就放在 `./grpc_test/protos/test.pb.go`。一來 package name 和目錄結構就相符合。

#### grpc_test/protos/test.pb.go

```go { .line-numbers }
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package protos

import (
    fmt "fmt"
    proto "github.com/golang/protobuf/proto"
    timestamp "github.com/golang/protobuf/ptypes/timestamp"
    math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Hello struct {
    Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
    Time                 *timestamp.Timestamp `protobuf:"bytes,99,opt,name=time,proto3" json:"time,omitempty"`
    XXX_NoUnkeyedLiteral struct{}             `json:"-"`
    XXX_unrecognized     []byte               `json:"-"`
    XXX_sizecache        int32                `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
    return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
    xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
    return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
    xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetName() string {
    if m != nil {
        return m.Name
    }
    return ""
}

func (m *Hello) GetTime() *timestamp.Timestamp {
    if m != nil {
        return m.Time
    }
    return nil
}

func init() {
    proto.RegisterType((*Hello)(nil), "protos.Hello")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
    // 162 bytes of a gzipped FileDescriptorProto
    0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
    0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0xd6, 0xe9, 0x99, 0x25,
    0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0xfa, 0x60,
    0x99, 0xa4, 0xd2, 0x34, 0xfd, 0x82, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0x92, 0xcc, 0xdc, 0xd4,
    0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x04, 0x0b, 0x62, 0x88, 0x92, 0x37, 0x17, 0xab, 0x47, 0x6a, 0x4e,
    0x4e, 0xbe, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
    0x10, 0x98, 0x2d, 0xa4, 0xc7, 0xc5, 0x02, 0x52, 0x2f, 0x91, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24,
    0xa5, 0x97, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x07, 0x33, 0x5d, 0x2f, 0x04, 0x66, 0x58, 0x10,
    0x58, 0x9d, 0x93, 0x50, 0x94, 0x40, 0x7a, 0x51, 0x41, 0x72, 0x3c, 0xc8, 0x91, 0x10, 0x37, 0x14,
    0x27, 0x41, 0x5c, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xca, 0x1c, 0xd4, 0xff, 0xba, 0x00,
    0x00, 0x00,
}
```

### 為 Protobuf Message 新增 Method

如果要需要新增功能，要另外用新檔案來處理，如: `test.go`。否則更新 protobuf 定義時，會覆蓋原先的修改的程式。

也可以使用 `go generate` 的方式，在 test.go 內加 `//go:generate protoc -I ../../grpc_test/protos -I $GOPATH/src --go_out=../../ test.proto`。請留意路徑設定。

#### grpc_test/protos/test.go

```go {.line-numbers}
package protos

//go:generate protoc -I ../../grpc_test/protos -I $GOPATH/src --go_out=../../ test.proto

import (
    proto "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/ptypes"
)

// CreateHello ...
func CreateHello(name string) *Hello {
    return &Hello{
        Name: name,
        Time: ptypes.TimestampNow(),
    }
}

// UnmarshalHello ...
func UnmarshalHello(data []byte) (*Hello, error) {
    ret := &Hello{}

    if err := proto.Unmarshal(data, ret); err != nil {
        return nil, err
    }

    return ret, nil
}

// MarshalHello ...
func MarshalHello(data *Hello) ([]byte, error) {
    return proto.Marshal(data)
}
```

### Marshal / Unmarshal

使用 protobuf 與 JSON 類似。

```go { .line-numbers }
import (
    proto "github.com/golang/protobuf/proto"
)
// UnmarshalHello ...
func UnmarshalHello(data []byte) (*Hello, error) {
    ret := &Hello{}

    if err := proto.Unmarshal(data, ret); err != nil {
        return nil, err
    }

    return ret, nil
}

// MarshalHello ...
func MarshalHello(data *Hello) ([]byte, error) {
    return proto.Marshal(data)
}
```

### gRPC (Google Remote Procedure Call)

也是撰寫 .proto ，建議定義 gRPC service 要與資料 message 分開, 只放 service 會用到的 message，一來程式管理比較方便，二來也避免互相干擾。

#### grpc_test/service/service.proto

```protobuf
syntax = "proto3";

package service;

import "grpc_test/protos/test.proto";

option go_package = "grpc_test/service";

message Request {
    string name = 1;
}

service HelloService {
    rpc Hello(Request) returns (protos.Hello) {}
}
```

### Serivce Definition

主要 gRPC 的定義是這一段：

```go { .line-numbers }
service HelloService {
    rpc Hello(Request) returns (protos.Hello) {}
}
```

1. 用 `rpc` 與 `returns` 這兩個關鍵字來定義 service.
1. 與上述動作一樣，切換到 `class18`，執行 `protoc -I grpc_test/service -I . -I $GOPATH/src --go_out=plugins=grpc:. grpc_test/service/*.proto`。
    1. 與上述不一樣的地方，是在 `--go_out` 這個多了 `plugins=grpc` 設定。
    1. `-I .` 主要是要讓 protoc 來尋找 `grpc_test/protos/test.proto`
    1. 在 `class18/grpc_test/service` 的目錄下，會產生 `service.pb.go`，一樣不建議直接修改 `service.pb.go`，有新加功能，都另開檔案來處理，eg: `service.go`

#### grpc_test/service/service.pb.go

```go { .line-numbers }
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
    context "context"
    fmt "fmt"
    proto "github.com/golang/protobuf/proto"
    grpc "google.golang.org/grpc"
    protos "grpc_test/protos"
    math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
    Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
    return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
    return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
    return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
    xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
    return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
    xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
    if m != nil {
        return m.Name
    }
    return ""
}

func init() {
    proto.RegisterType((*Request)(nil), "service.Request")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
    // 138 bytes of a gzipped FileDescriptorProto
    0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
    0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0xa4, 0xd3,
    0x8b, 0x0a, 0x92, 0xe3, 0x4b, 0x52, 0x8b, 0x4b, 0xf4, 0xc1, 0x32, 0xc5, 0xfa, 0x20, 0x36, 0x44,
    0x95, 0x92, 0x2c, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b,
    0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x6d, 0x64, 0xc5, 0xc5,
    0xe3, 0x91, 0x9a, 0x93, 0x93, 0x1f, 0x0c, 0x31, 0x4b, 0x48, 0x8b, 0x8b, 0x15, 0xcc, 0x17, 0x12,
    0xd0, 0x83, 0xd9, 0x06, 0xd5, 0x2e, 0xc5, 0x0b, 0x31, 0xb1, 0x58, 0x0f, 0xac, 0x40, 0x89, 0xc1,
    0x49, 0x38, 0x4a, 0x10, 0x61, 0x33, 0x54, 0x75, 0x12, 0x1b, 0x58, 0x91, 0x31, 0x20, 0x00, 0x00,
    0xff, 0xff, 0xd0, 0xa6, 0x59, 0xf2, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
    Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*protos.Hello, error)
}

type helloServiceClient struct {
    cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
    return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*protos.Hello, error) {
    out := new(protos.Hello)
    err := c.cc.Invoke(ctx, "/service.HelloService/Hello", in, out, opts...)
    if err != nil {
        return nil, err
    }
    return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
    Hello(context.Context, *Request) (*protos.Hello, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
    s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(Request)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil {
        return srv.(HelloServiceServer).Hello(ctx, in)
    }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
        FullMethod: "/service.HelloService/Hello",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(HelloServiceServer).Hello(ctx, req.(*Request))
    }
    return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "service.HelloService",
    HandlerType: (*HelloServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "Hello",
            Handler:    _HelloService_Hello_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "service.proto",
}
```

### Server and Client Interface

gRPC 主要會定義 server 與 client 的 interface。

```go { .line-numbers }
type HelloServiceClient interface {
    Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*protos.Hello, error)
}

type HelloServiceServer interface {
    Hello(context.Context, *Request) (*protos.Hello, error)
}
```

### Server 實作 (grpc_test/server/main.go)

```go { .line-numbers }
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "grpc_test/protos"
    "grpc_test/service"

    "google.golang.org/grpc"
)

type helloService struct{}

func (h *helloService) Hello(ctx context.Context, req *service.Request) (*protos.Hello, error) {
    if req == nil || "" == req.Name {
        return nil, fmt.Errorf("request is not ok: %v", req)
    }

    ret := protos.CreateHello(req.Name)
    log.Println("resp:", ret)

    return ret, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    service.RegisterHelloServiceServer(s, &helloService{})

    log.Println("serving...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }

    log.Println("start....")
}
```

1. listen port: `lis, err := net.Listen("tcp", ":50051")`
1. New gRPC Server: `s := grpc.NewServer()`
1. register:

    ```go { .line-numbers }
    service.RegisterHelloServiceServer(s, &helloService{})
    ```

1. Serv:

    ```go { .line-numbers }
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
    ```

### Client 實作 (grpc_test/client/main.go)

```go { .line-numbers }
package main

import (
    "context"
    "fmt"
    "log"

    "grpc_test/service"

    "google.golang.org/grpc"
)

func main() {

    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        panic(fmt.Sprintf("dial grpc server error: %v", err))
    }
    defer conn.Close()

    client := service.NewHelloServiceClient(conn)

    resp, err := client.Hello(context.TODO(), &service.Request{Name: "Bob"})

    log.Println(resp)

    log.Println("end...")
}
```

說明：

1. connect to service: `conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())`, 因為沒有設定加密，因此要多一個 `grpc.WithInsecure()` 選項。(gRPC 預設是要用加密的，但我們沒有加密的相關設定，因此請用 *Insecure*)
1. 透過 connection 產生 client: `client := service.NewHelloServiceClient(conn)`
1. 呼叫 service 的 function: `resp, err := client.Hello(context.Background(), &service.Request{Name: "Bob"})`
