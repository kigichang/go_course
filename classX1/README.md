# X1 Build and Dependency Management

## Build

執行指令：`env GOOS=target-OS GOARCH=target-architecture go build package-import-path`。 eg: 編繹成 linux 64bit 版本，`env GOOS=linux GOARCH=amd64 go build .`

在 Go 1.11 版本，支援 Web Assembly, 可以將 Go 的程式，編繹了 WebAssembly，讓瀏覽器執行。eg: `env GOOS=js GOARCH=wasm go build .`

要查支援的環境，可以執行 `go tool dist list` 查詢。

| GOOS | GOARCH |
|:----:|:------:|
| android | 386 |
| android | amd64 |
| android | arm |
| android | arm64 |
| darwin | 386 |
| darwin | amd64 |
| darwin | arm |
| darwin | arm64 |
| dragonfly | amd64 |
| freebsd | 386 |
| freebsd | amd64 |
| freebsd | arm |
| js | wasm |
| linux | 386 |
| linux | amd64 |
| linux | arm |
| linux | arm64 |
| linux | mips |
| linux | mips64 |
| linux | mips64le |
| linux | mipsle |
| linux | ppc64 |
| linux | ppc64le |
| linux | riscv64 |
| linux | s390x |
| nacl | 386 |
| nacl | amd64p32 |
| nacl | arm |
| netbsd | 386 |
| netbsd | amd64 |
| netbsd | arm |
| openbsd | 386 |
| openbsd | amd64 |
| openbsd | arm |
| plan9 | 386 |
| plan9 | amd64 |
| plan9 | arm |
| solaris | amd64 |
| windows | 386 |
| windows | amd64 |

**Warning: Cross-compiling executables for Android requires the Android NDK, and some additional setup which is beyond the scope of this tutorial.**

## Dependency Management

Go 在 1.11 版後，加了 `go mod` 工具來管理 library dependency。原本專案目錄要設定在 GOPATH 下，如果要使用 mod 工具來進行開發，專案目錄不能放在 GOPATH 下，否則 Go 不會啟用 go module 式。

### 命名規則

mod 的命名規則與原本 `go get` 相同。主要是由 `source_controll_service/editor_or_group_name/project_name` 組成。

1. source_controll_service: 像 github.com, gitlab.com，或我們自建的 cicd.icu。
1. editor_or_group_name: 在版控服務上的作者或群組名稱，像 `spf13` 或 `cyberon`。
1. project_name: 專案名稱，像：`viper` 或 `config`。

像我們常用的 `github.com/spf13/viper` 或 `cicd.icu/cyberon/config`。在 GO 程式 import 時，需撰寫完整的 url。eg: `import github.com/spf13/viper` 或 `cicd.icu/cyberon/config`。

原本 Go 下載 package 的方式，是使用 `go get` 的工具，如 `go get -u github.com/spf13/viper`，`go get` 工具會去下載並存在 `$GOPATH/src/`，如: `$GOPATH/src/github.com/spf13/viper`。如此撰寫 import 時，就會找到相對應的 package。

因為我們自建的 gitlab 沒有使用 ssl, 因此需要在 git 設定檔 **YOUR_HOME/.gitconfig** 下，加入以下設定。

```text
[url "[git@cicd.icu:10022]:"]
    insteadOf = http://cicd.icu/
```

並在下載時，多加 `--insecure` 參數。eg: `go get -u --insecure cicd.icu/root/hello`。

### Mod

在新的專案目錄下，執行 `go mod init project_url`，如: `go mod init cicd.icu/root/hello`。此時在專案的目錄下，會產生 **go.mod** 的檔案。內容會是：

```go
module cicd.icu/root/hello
```

如果專案會用到其他的 package。在做 `go build` 時，會自動去下載需要的 package，放在 **$GOPATH/pkg/mod** 下。如果需要指定 package 版本。建議編輯 **go.mod**，加入 `require package_url version`，如：`require cicd.icu/cyberon/config v0.0.1`。然後再執行 `go mod download` 來下載相對應的 package。

步驟整理：

1. 編輯 **go.mod**
1. 執行 `go mod download`
1. 如果要整理使用的 package，並移除沒再使用的 package，可以用 `go mod tidy`
