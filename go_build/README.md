# Go Build

## Build

Go 支授多種 CPU 與作業系統平台，也可以編譯成 Android, iOS, 及 WebAssembly。執行指令：`env GOOS=target-OS GOARCH=target-architecture go build package-import-path`。 eg: 編繹成 linux 64bit 版本，`env GOOS=linux GOARCH=amd64 go build .`

在 Go 1.11 版本，支援 Web Assembly, 可以將 Go 的程式，編繹了 WebAssembly，讓瀏覽器執行。eg: `env GOOS=js GOARCH=wasm go build .`

要查支援的環境，可以執行 `go tool dist list` 查詢。


Go 編譯成 Anodroid/iOS 說明：[Gomobile](https://github.com/golang/go/wiki/Mobile)
Go WebAssembly 說明：[Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

## Dependency Management

Go 在 1.11 版後，加了 `go mod` 工具來管理 library dependency。使用 Go Module 管理，專案可以不用放在 `$GOPATH`下。~~原本專案目錄要設定在 GOPATH 下，如果要使用 mod 工具來進行開發，專案目錄不能放在 GOPATH 下，否則 Go 不會啟用 go module 式。~~

### 命名規則

mod 的命名規則與原本 `go get` 相同。主要是由 `source_controll_service/editor_or_group_name/project_name` 組成。

1. source_controll_service: 像 github.com, gitlab.com，或我們自建的私有庫。
1. editor_or_group_name: 在版控服務上的作者或群組名稱，像 `spf13` 或 `mygroup`。
1. project_name: 專案名稱，像：`viper` 或 `config`。

像我們常用的 `github.com/spf13/viper` 或 `private_git_service/mygroup/myproject`。在 GO 程式 import 時，需撰寫完整的 url。eg: `import github.com/spf13/viper` 或 `private_git_service/mygroup/myproject`。

自建的 gitlab 需要帳號/密碼登入，無法直接用 http , 因此需要在 git 設定檔 **YOUR_HOME/.gitconfig** 下，加入以下設定。

```text
[url "[git@private_git_service:other_port]:"]
    insteadOf = http://private_git_service/
```

eg:

```text
[url "[git@abc.xyz:1234]:"]
    insteadOf = http://abc.xyz/
```

由於 Go Module (Go >= 1.13) 會透過 [goproxy.io](https://goproxy.io/) 取得使用的套件，如果是使用私有庫 (private repository) 的話，需要設定 `$GOPRIVATE` 環境變數。eg: `env GOPRIVATE=“abc.xyz/group_name” go mod tidy`

Go 官方不建議使用 insecure 的方式，因此自建的 Git repository 使用用 [CertBot](https://certbot.eff.org/) 來取得免費的 SSL 憑証。
~~原本 Go 下載 package 的方式，是使用 `go get` 的工具，如 `go get -u github.com/spf13/viper`，`go get` 工具會去下載並存在 `$GOPATH/src/`，如: `$GOPATH/src/github.com/spf13/viper`。如此撰寫 import 時，就會找到相對應的 package，並在下載時，多加 `--insecure` 參數。eg: `go get -u --insecure cicd.icu/root/hello`。~~

### Go Module

在新的專案目錄下，執行 `go mod init project_url`，如: `go mod init abc.xyz/hello`。此時在專案的目錄下，會產生 **go.mod** 的檔案。內容會是：

```go
module abc.xyz/hello

go 1.17
```

如果專案會使用到其他 package 時，通常 IDE 工具會自動編輯 **go.mod**，會編譯過程，會產生 **go.sum** (不用理會)。如果 IDE 工具沒有編輯時，或者需要指定使用 package 版本時，則需要手動編輯。

有時我們會同時間編輯多個專案，且專案有相依性；又或者使用 package 要導向另一個網址時，可以使用 `replace`，來指定專案路徑，或者是新的 URL。eg:

```go
module myhello

go 1.17

require abc.xyz/hello v0.0.0

replace abc.xyz/hello => ../hello
```

ps: 為了讓 go mod 不到網路上查詢 `abc.xyz` (會編譯失數)，在編輯 **go.mod** 時，請先自定 package 版本，eg: `v0.0.0`

~~如果專案會用到其他的 package。在做 `go build` 時，會自動去下載需要的 package，放在 **$GOPATH/pkg/mod** 下。如果需要指定 package 版本。建議編輯 **go.mod**，加入 `require package_url version`，如：`require cicd.icu/cyberon/config v0.0.1`。然後再執行 `go mod download` 來下載相對應的 package。~~

步驟整理：

1. 編輯 **go.mod**
~~1. 執行 `go mod download`~~
1. 如果要整理使用的 package，並移除沒再使用的 package，可以用 `go mod tidy`

// TODO: Build Constrain
// TODO: Build CFlags to modify global variables.
// TODO: with Makefile