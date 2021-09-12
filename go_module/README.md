# Go Module


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Go Module](#go-module)
  - [0. 前言](#0-前言)
  - [1. 命名規則](#1-命名規則)
  - [2. 私有庫設定](#2-私有庫設定)
  - [3. Go Module 常用指令](#3-go-module-常用指令)
    - [3.1 Replace](#31-replace)

<!-- /code_chunk_output -->

## 0. 前言
Go 在 1.11 版後，加了 `go mod` 工具來管理 Package Dependency。使用 Go Module 管理，專案可以不用放在 `$GOPATH`下。

## 1. 命名規則

mod 的命名規則與原本 `go get` 相同。主要是由 `SOURCE_CONTROL_SERVICE/AUTHOR_OR_GROUP_NAME/PROJECT_NAME` 組成。

1. SOURCE_CONTROL_SERVICE: 像 github.com, gitlab.com，或我們自建的私有庫。
1. AUTHOR_OR_GROUP_NAME: 在版控服務上的作者或群組名稱，像 `spf13` 或 `mygroup`。
1. PROJECT_NAME: 專案名稱，像：`viper` 或 `config`。

像我們常用的 `github.com/spf13/viper` 或 `PRIVATE_GIT_SERVICE/mygroup/myproject`。在 GO 程式 import 時，需撰寫完整的 url。eg: `import "github.com/spf13/viper"` 或 `"PRIVATE_GIT_SERVICE/mygroup/myproject"`。

## 2. 私有庫設定

自建的 GitLab 需要帳號/密碼登入，無法直接用 http，需使用 SSH Key 方式存取程式。在 git 設定檔 **YOUR_HOME/.gitconfig** 下，加入以下設定。

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

## 3. Go Module 常用指令

在新的專案目錄下，執行 `go mod init package_name`，如: `go mod init abc.xyz/hello`。此時在專案的目錄下，會產生 **go.mod** 的檔案。內容會是：

```go
module abc.xyz/hello

go 1.17
```

如果專案會使用到其他 package 時，通常 IDE 工具會自動編輯 **go.mod**，會編譯過程，會產生 **go.sum** (不用理會)。如果 IDE 工具沒有編輯時，或者需要指定使用 package 版本時，則需要手動編輯。

### 3.1 Replace

有時我們會同時間編輯多個專案，且專案有相依性；又或者使用 package 要導向另一個網址時，可以使用 `replace`，來指定專案路徑，或者是新的 URL。eg:

```go
module myhello

go 1.17

require abc.xyz/hello v0.0.0

replace abc.xyz/hello => ../hello
```

ps: 為了讓 go mod 不到網路上查詢 `abc.xyz` (會編譯失數)，在編輯 **go.mod** 時，請先自定 package 版本，eg: `v0.0.0`

步驟整理：

1. 編輯 **go.mod**
1. 如果要整理使用的 package，並移除沒再使用的 package，可以用 `go mod tidy`