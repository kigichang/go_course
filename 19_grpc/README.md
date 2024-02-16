# 19 ProtoBuf and gRPC


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [19 ProtoBuf and gRPC](#19-protobuf-and-grpc)
  - [0. 前言](#0-前言)
  - [1. 安裝 protoc](#1-安裝-protoc)
  - [2. Package 與目錄結構](#2-package-與目錄結構)
  - [3. 撰寫 .proto](#3-撰寫-proto)
    - [3.1 protos/test.proto](#31-protostestproto)
    - [3.2 組成元素](#32-組成元素)
    - [3.3 Protocol Buffers Coding Style](#33-protocol-buffers-coding-style)
    - [3.4 資料型別](#34-資料型別)
    - [3.5 message Hello](#35-message-hello)
  - [4 轉成 Go 程式](#4-轉成-go-程式)
    - [4.1 grpc_test/protos/test.pb.go](#41-grpc_testprotostestpbgo)
    - [4.2 為 Protobuf Message 新增 Method](#42-為-protobuf-message-新增-method)
    - [4.3 Marshal / Unmarshal](#43-marshal-unmarshal)
  - [5. gRPC (Google Remote Procedure Call)](#5-grpc-google-remote-procedure-call)
    - [5.1 grpc_test/service/service.proto](#51-grpc_testserviceserviceproto)
    - [5.2 Serivce Definition](#52-serivce-definition)
    - [5.3 Server and Client Interface](#53-server-and-client-interface)
    - [5.4 Server 實作 (grpc_test/server/main.go)](#54-server-實作-grpc_testservermaingo)
    - [5.5 Client 實作 (grpc_test/client/main.go)](#55-client-實作-grpc_testclientmaingo)

<!-- /code_chunk_output -->

## 0. 前言

ProtoBuf 是 Google 開發的工具，主要來取代 JSON, 與 XML，通常會用在 RPC (Remote Procedure Call) 上，也因此 ProtoBuf 會撘配 Google 開發的 gRPC 使用。ProtoBuf 本身支援多種常用的程式語言，也因此可以利用 ProtoBuf 當作中介的橋樑，在不同的程式語言間，交換資料。

相關資料：

- [Protocol Buffers (ProtoBuf) 官網](https://developers.google.com/protocol-buffers/)
- [Developer Guide](https://developers.google.com/protocol-buffers/docs/overview)
- [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)

## 1. 安裝 protoc

**protoc** 是 Protobuf 的工具，主要是將 protobuf 的定義檔 (.proto) 轉成對應的程式語言。使用 protoc 時，要留意專案的目錄結構。以本例來說，專案 package name 是 `grpc_test`，專案的目錄名稱也設定成 `grpc_test`。

1. 到 [protoc release](https://github.com/google/protobuf/releases) 下載對應作業系統 (Linux, OSX, Win32) 的執行檔。如：__protoc-3.17.3-osx-x86_64.zip__
1. 解壓縮上述檔案。將目錄內的 __bin/protoc__ copy 至 __$GOPATH/bin__
1. 將目錄內的 __include__ copy 至 __$GOPATH__。
1. 執行 `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` 下載 protoc 的 go plugin。
1. 執行 `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`下載 protoc 的 go grpc plugin。

## 2. Package 與目錄結構

由於 Go package 會和目錄結構相關，因此在撰寫專案時，要留意 package name 與目錄。本例測試專案 package name 為 `grpc_test` (見 go.mod)，因此目錄名稱也命名為 `grpc_test`。

```text
19_grpc/grpc_test
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

## 3. 撰寫 .proto

使用 protobuf 前，我們需要先定義資料格式，寫起來有點像在寫 struct。首先在專案目錄下，開一個目錄，如: `protos`，在 `protos` 下還可以依功能再細分。

### 3.1 protos/test.proto

@import "grpc_test/protos/test.proto" {as="protobuf" class=line-numbers}

### 3.2 組成元素

1. syntax: `syntax = "proto3";` 指定 protobuf 的版本，目前有 proto2 與 proto3。建議用 proto3.
1. package: 定義 protobuf 的 package, eg: `package protos;`
1. import: 如果有用到其他的 protobuf 資料型別，一樣需要 import, eg: `google/protobuf/timestamp.proto`
1. `option go_package = "grpc_test/protos";` 可以指定對應的程式語言，package name 為何。像 GO, Java 的 package name 也會對應到目錄結構，在產出 `*.pb.go` 時，會使用此設定，建立相對應的目錄。
1. message: 定義資料結構 `message 資料名稱`。
    - `google.protobuf.Timestamp` 是定義在 __google/protobuf/timestamp.proto__。這個檔案是在 __$GOPATH/google/protobuf/timestamp.proto__ 。檔案內的定義 `package google.protobuf;`，因此使用 google timestamp 的資料型別會是 `google.protobuf.Timestamp`。

### 3.3 Protocol Buffers Coding Style

[Style Guide](https://developers.google.com/protocol-buffers/docs/style)

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

### 3.4 資料型別

[proto3](https://developers.google.com/protocol-buffers/docs/proto3)

### 3.5 message Hello

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

## 4 轉成 Go 程式

1. 目錄切到 `19_grpc`
1. protoc 需要加入 `-I` 來指定搜尋 *.proto 的 路徑。
    - 由於使用 `google.protobuf.Timestamp`，必須指定 __google/protobuf/timestamp.proto__ 檔案在那。
1. 執行 `protoc -I grpc_test/protos -I $GOPATH/include --go_out=. grpc_test/protos/*.proto`
    - `-I` 類似 C 的 include，指定路徑，讓 protoc 去找尋相依的 protobuf 檔案。
    - 放 proto 檔案的目錄，也必須加到 `-I`。eg: `-I grpc_test/protos`
    - 因為使用 timestamp，因此也需要加 `-I $GOPATH/include`.
    - `--go_out` 是指要輸出 Go 的程式，並指定目標目錄。
    - 由於在 test.proto 有指定 Go 的 package name `grpc_test/protos`，因此輸出的檔案，就放在 `./grpc_test/protos/test.pb.go`。一來 package name 和目錄結構就相符合。

### 4.1 grpc_test/protos/test.pb.go

__test.pb.go__ 由 __test.proto__，產生的 Go 程式檔。不要去修改這個檔案。

@import "grpc_test/protos/test.pb.go" {class=line-numbers}

### 4.2 為 Protobuf Message 新增 Method

如果要需要新增功能，要另外用新檔案來處理，如: `test.go`。否則更新 protobuf 定義時，會覆蓋原先的修改的程式。

也可以使用 `go generate` 的方式，在 test.go 內加 `//go:generate protoc -I ../../grpc_test/protos -I $GOPATH/include --go_out=../../ test.proto`。請留意路徑設定。

@import "grpc_test/protos/test.go" {class=line-numbers}

### 4.3 Marshal / Unmarshal

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

## 5. gRPC (Google Remote Procedure Call)

也是撰寫 .proto ，建議定義 gRPC service 要與資料 message 分開, 只放 service 會用到的 message，一來程式管理比較方便，二來也避免互相干擾。

詳細說明：[gRPC](https://www.grpc.io/docs/)

### 5.1 grpc_test/service/service.proto

@import "grpc_test/service/service.proto" {as="protobuf" class="line-numbers"}

### 5.2 Serivce Definition

主要 gRPC 的定義是這一段：

```go { .line-numbers }
service HelloService {
    rpc Hello(Request) returns (protos.Hello) {}
}
```

1. 用 `rpc` 與 `returns` 這兩個關鍵字來定義 service.
1. 與上述動作一樣，切換到 `19_grpc`，執行 `protoc -I grpc_test/service -I . -I $GOPATH/include --go-grpc_out=. --go_out=. grpc_test/service/*.proto`。
    1. 與上述不一樣的地方，多了 `--go-grpc_out`。
    1. `-I .` 主要是要讓 protoc 來尋找 `grpc_test/protos/test.proto`
    1. 在 `grpc_test/service` 的目錄下，會產生 `service_grpc.pb.go` 與 `service.pb.go`，一樣不建議修改。

#### service.pb.go

@import "grpc_test/service/service.pb.go" {class=line-numbers}

#### service_grpc.pb.go

@import "grpc_test/service/service_grpc.pb.go" {class=line-numbers}

### 5.3 Server and Client Interface

gRPC 主要會定義 server 與 client 的 interface。

```go { .line-numbers }
// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
    Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*protos.Hello, error)
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
    Hello(context.Context, *Request) (*protos.Hello, error)
    mustEmbedUnimplementedHelloServiceServer()
}
```

### 5.4 Server 實作 (grpc_test/server/main.go)

#### grpc_test/server/main.go

@import "grpc_test/server/main.go" {class="line-numbers"}

1. 在 `helloService` 加入 `service.UnimplementedHelloServiceServer`。
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

### 5.5 Client 實作 (grpc_test/client/main.go)

#### grpc_test/client/main.go

@import "grpc_test/client/main.go" {class=line-numbers}

說明：

1. connect to service: `conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())`, 因為沒有設定加密，因此要多一個 `grpc.WithInsecure()` 選項。(gRPC 預設是要用加密的，但我們沒有加密的相關設定，因此請用 *Insecure*)
1. 透過 connection 產生 client: `client := service.NewHelloServiceClient(conn)`
1. 呼叫 service 的 function: `resp, err := client.Hello(context.Background(), &service.Request{Name: "Bob"})`
