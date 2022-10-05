# 15 flag and spf13 Cobra/Viper

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [15 flag and spf13 Cobra/Viper](#15-flag-and-spf13-cobraviper)
  - [0. 說明](#0-說明)
  - [1. Go flag (ex15_01)](#1-go-flag-ex15_01)
  - [2. Viper (ex15_02)](#2-viper-ex15_02)
    - [2.1 取值](#21-取值)
    - [2.2 自定 Viper, 不使用預設的 viper](#22-自定-viper-不使用預設的-viper)
  - [3 Cobra](#3-cobra)
    - [3.1 Cobra 參數管理方式](#31-cobra-參數管理方式)
    - [3.2 定義 Sub commands](#32-定義-sub-commands)
    - [3.3 定義 flag，參數與工作](#33-定義-flag參數與工作)
    - [3.3 測試](#33-測試)
    - [3.4 與 Viper 結合](#34-與-viper-結合)

<!-- /code_chunk_output -->

## 0. 說明

- Go flag: 內建程式參數管理。
- spf13/Cobra: 進階程式參數管理，可實作類似 `go` 或 `docker` 程式效果。
- spf13/Viper: 設定檔套件，支援多種設定檔格式(.json, .yaml 等)。

## 1. Go flag (ex15_01)

@import "ex15_01/main.go" {class="line-numbers"}

說明：

1. 先設定會傳入的參數名稱與資料型別：
    - `flag.StringVar(&account, "account", account, "account to login")`
    - `flag.StringVar(&password, "password", password, "password for account")`
    - `flag.BoolVar(&debug, "debug", debug, "dump account and password or not")`
1. 呼叫 `flag.Parse()` 處理傳入的參數
1. 驗証資料是否正確，否則可以呼叫 `flag.PrintDefaults()` 輸出說明
1. 執行時，可用 `--` or `-` 接參數名稱，參數值前可加 `=` 或空白, eg:
    - `-a=1`
    - `-a 1`
    - `--a=1`
    - `-a=1`

執行：

- `go run .`:

    ```console {.line-numbers}
    -account string
            account to login
    -debug
            dump account and password or not
    -password string
            password for account
    ```

- `go run . -account 123 -password 321`:

    ```console {.line-numbers}
    2020/01/16 15:22:13 end
    ```

- `go run . --account 123 --password 321 --debug`:

    ```console {.line-numbers}
    2020/01/16 15:22:44 account: 123 password: 321
    2020/01/16 15:22:44 end
    ```

## 2. Viper (ex15_02)

- [Viper](https://github.com/spf13/viper) 設定檔套件。
- 支援 JSON, TOML, YAML, HCL, and Java properties 等格式
- 可設定預設值

在 import viper 後，會有一個 global 變數 **viper**，可以直接用。在使用前，需先設定要載入的檔名以及設定檔放的路徑，最後再執行讀取。

eg: 在專案的目錄下，放置一個 **config.json** 的設定檔，Viper 設定好目錄與設定檔名(**不含副檔名**)，呼叫 `ReadInConfig`，來載入設定檔。

eg:

@import "ex15_02/main.go" {class="line-numbers"}

設定檔 **config.json**:

@import "ex15_02/config.json" {class="line-numbers"}

結果：

```console {.line-numbers}
abc:  def
aaa:  true
def:
```

### 2.1 取值

```go { .line-numbers }
Get(key string) : interface{}
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetString(key string) : string
GetStringMap(key string) : map[string]interface{}
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool
```

注意：Viper 取值時，如果 key 不存在，會回傳 zero value。所以要特別小心。

[詳細說明](https://github.com/spf13/viper#getting-values-from-viper)

### 2.2 自定 Viper, 不使用預設的 viper

Viper 也允許自己產生一個全新的 viper，方便管理不同的設定檔。

@import "ex15_03/main.go" {class="line-numbers"}

## 3 Cobra

[Cobra](https://github.com/spf13/cobra)

管理 Command line 程式參數的套件，雖然 Go 已有內建 Flag 套件，但很多第三方套件還是用 Cobra。

### 3.1 Cobra 參數管理方式

以 docker 為例，當執行 `docker` 時：

```text
Usage:  docker COMMAND

A self-sufficient runtime for containers

Options:
      --config string      Location of client config files (default "/Users/kigi/.docker")
  -D, --debug              Enable debug mode
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
      --tls                Use TLS; implied by --tlsverify
      --tlscacert string   Trust certs signed only by this CA (default "/Users/kigi/.docker/ca.pem")
      --tlscert string     Path to TLS certificate file (default "/Users/kigi/.docker/cert.pem")
      --tlskey string      Path to TLS key file (default "/Users/kigi/.docker/key.pem")
      --tlsverify          Use TLS and verify the remote
  -v, --version            Print version information and quit

Management Commands:
  checkpoint  Manage checkpoints
  config      Manage Docker configs
  container   Manage containers
  ...

Commands:
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
  commit      Create a new image from a container's changes
  ...

Run 'docker COMMAND --help' for more information on a command
```

`docker` 是主程式，之後會再接 command，在 cobra 的設計中，`docker` 是 root command，以下的 command 稱做 sub command。

`docker run --help` 為例:

```text
Usage:  docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

Run a command in a new container

Options:
      --add-host list                  Add a custom host-to-IP mapping (host:ip)
  -a, --attach list                    Attach to STDIN, STDOUT or STDERR
      --blkio-weight uint16            Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
      --blkio-weight-device list       Block IO weight (relative device weight) (default [])
  ...
```

`--xxx` 與 `-x` 是指傳入的 flag，一個 flag 可以有 longterm (`--xxx` 表示)，與 shortterm (`-x` 表示)，變數可以是 string, numbers, boolean 等資料型別。

以下，是模擬以上的效果。

### 3.2 定義 Sub commands

先產生 root 及 sub commands

eg:

```go { .line-numbers }
package main

import (
    "github.com/spf13/cobra"
)

func main() {
    rootCmd := &cobra.Command{Use: "myapp"}

    createCmd := &cobra.Command{Use: "create"}

    updateCmd := &cobra.Command{Use: "update"}

    rootCmd.AddCommand(createCmd, updateCmd)

    rootCmd.Execute()
}
```

在專案目錄下，執行 `go run .` 結果是：

```text
Usage:

Flags:
  -h, --help   help for myapp

Additional help topics:
  myapp create
  myapp update
```

執行 `go run . create` or `go run . update` 會立即執行完畢，因為我們還沒定義 sub command 要做什麼事情。

### 3.3 定義 flag，參數與工作

接下來定義每個 sub command 需要的 flag, 參數與工作。

@import "ex15_04/main.go" {class=line-numbers}

1. 定義兩個 flag，`name` 及 `proxy`

    ```go { .line-numbers }
    createCmd.Flags().StringVarP(&name, "name", "n", "myname", "assign a name")
    createCmd.Flags().BoolVarP(&proxy, "proxy", "p", false, "use proxy to connect")
    ```

1. 設定只能有一個參數。詳細設定，請見：[cobra#Positional and Custom Arguments](https://github.com/spf13/cobra#positional-and-custom-arguments)

    ```go { .line-numbers }
    createCmd.Args = cobra.ExactArgs(1)
    ```

1. 設定執行動作

    ```go { .line-numbers }
    createCmd.Run = func(cmd *cobra.Command, args []string) {
        fmt.Println("creating")
        fmt.Println("name:", name)
        fmt.Println("proxy:", proxy)
        fmt.Println("args:", args)
    }
    ```

### 3.3 測試

1. `go run . create`

    ```console {.line-numbers}
    Error: accepts 1 arg(s), received 0
    Usage:
        myapp create [flags]

    Flags:
        -h, --help          help for create
        -n, --name string   assign a name (default "myname")
        -p, --proxy         use proxy to connect
    ```

    會發現回傳錯誤，並沒有執行動作。因為我們並沒有傳入任何參數。

1. `go run . create abc`

    ```console {.line-numbers}
    creating
    name: myname
    proxy: false
    args: [abc]
    ```

    執行成功，並使用預設值

1. `go run . create abc def`

    ```console {.line-numbers}
    Error: accepts 1 arg(s), received 2
    Usage:
        myapp create [flags]

    Flags:
        -h, --help          help for create
        -n, --name string   assign a name (default "myname")
        -p, --proxy         use proxy to connect
    ```

    執行失敗，因為多傳了一個參數。

1. `go run . create --name=bob --proxy abc` or `go run . create -n bob -p abc`

    ```console {.line-numbers}
    creating
    name: bob
    proxy: true
    args: [abc]
    ```

    boolean 型別的 flag，後面可以不用接值。

### 3.4 與 Viper 結合

可以將 flag 當作設定檔的資料。如此一來，在大型的程式中，就可以統一都使用 Viper 來當共用設定，而這些設定可以是來自設定檔或者是 command line 的 flag。

請使用 `PersistentFlags` 撘配 `Viper` 使用。

eg:

#### main.go

@import "ex15_04/main.go" {class="line-numbers"}

#### config.json

@import "ex15_04/config.json" {class="line-numbers"}

#### 執行

執行: `go run . update -t 123`

結果:

```text
viper test: 123
[]
```

測試時，可以試著將設定檔的中的 `test` 移除，與執行時，不帶 `-t`, 看看執行的結果。

| - | config 有 test  | config 沒有 test |
|:-:|:-:|:-:|
| 有 -t (123) | test: 123 | test: 123 |
| 沒有 -t | test: xyz | test: my test |

#### Summary

1. 如果參數有值，則用參數值。
1. 如果參數沒有值，則看設定檔有沒有該設定。如果有，則用設定檔的值，如果沒有，則用程式預設的值 (mytest)。

值的讀取：參數值 > 設定檔設定 > 程式預設值
