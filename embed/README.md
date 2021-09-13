# Embed

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Embed](#embed)
  - [0. 前言](#0-前言)
  - [1. 實作](#1-實作)

<!-- /code_chunk_output -->

## 0. 前言

Go 在 1.6 推出 __embed__ 功能，主要是可以將程式會用到的檔案，一併包裝進執行檔。過往在沒這項功能時，在佈署程式，需要將相關的檔案一併放到佈署環境，或包入 Docker 內。有 __embed__ 即可省去這些工作，讓佈署更加方便。

__embed__ 主要是在執行時，虛擬出檔案系統的環境，讓程式在操作時，與一般環境相同。

## 1. 實作

@import "myembed/main.go" {class="line-numbers", highlight="4,9,12,15,16"}

1. import embed package。如：`import "embed"`，如果程式內沒有使用到 `embed` 相關，也需要 import，如：`import _ "embed"`。
1. 在指定的變數，加入 `go:embed`，嵌入特定資源。如：

    ```go {.line-numbers}
    //go:embed hello.txt
    var s string

    //go:embed hello.txt
    var b []byte
    ```
    將 __hello.txt__ 的內容，載入變數 `s` 與 `b`。
1. 嵌入多個資源。

    ```go {.line-numbers}
    //go:embed assets/*
    //go:embed hello.txt
    var content embed.FS
    ```
    
    `context` 在程式內，等同一個虛擬目錄，內含 __hello.txt__ 與 __assets__ 目錄。
1. 編譯成執行檔後，放置到其他目錄去執行，會發現即使沒把 __hello.txt__ 與 __assets__ 一起帶走，程式依然可以讀到檔案。