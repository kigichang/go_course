# 07 Methods

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [07 Methods](#07-methods)
  - [0. 前言](#0-前言)
  - [1. Struct Receiver (ex07_01)](#1-struct-receiver-ex07_01)
  - [2. Struct Pointer Receiver (ex07_02)](#2-struct-pointer-receiver-ex07_02)
  - [3. Methods in Value and Pointer](#3-methods-in-value-and-pointer)
    - [3.1 Struct (ex07_03)](#31-struct-ex07_03)
    - [3.2 Pointer](#32-pointer)
    - [3.3 注意](#33-注意)
  - [4. Method Signature (ex07_05)](#4-method-signature-ex07_05)
  - [5. Side Effects in Functional Language](#5-side-effects-in-functional-language)

<!-- /code_chunk_output -->

## 0. 前言

在 OOP 中，定義在 Class 內的 Function，稱為 Methods。Class 的 Method 來處理資料，達到封裝的功能。在 Go 也有一樣的功能，主要是針對 struct 來定義 method。Go 定義 Method 時，(Method 的 receiver) 可以使用 Struct 或者是 Struct Pointer。

## 1. Struct Receiver (ex07_01)

@import "ex07_01/main.go" {class=line-numbers}

## 2. Struct Pointer Receiver (ex07_02)

在定義 method 也可以使用 struct pointer。

@import "ex07_02/main.go" {class=line-numbers}

## 3. Methods in Value and Pointer

由上面的例子來看，不論 methods 是被宣告在 struct 或者是 struct pointer，只要是該 struct 或者是 struct pointer，都可以呼叫。而這兩者差別是，用 struct pointer 定義 method 要特別注意會修改到原本的值。

### 3.1 Struct (ex07_03)

使用 Method 時，不會修改到原本 Struct 值。

@import "ex07_03/main.go" {class=line-numbers}

### 3.2 Pointer

使用 Method 時，會修改到原本 Struct 值。(原因：傳 Pointer 至 Method)

@import "ex07_04/main.go"

### 3.3 注意

與 slice 類似，但因為是 method 很難查覺是否有修改原本的資料。因此在實作上，儘量 method 都用 pointer 的方式。

1. 避免 pass by value 的記憶體浪費
1. 避免 golang 在 struct pointer 語法上的 puzzle (因為 struct 與 struct pointer 在 call method 的語法都一樣，不像 C 有分 `.` 與 `->`).

## 4. Method Signature (ex07_05)

Method 本身就是 funcation，因此也有 signature.

@import "ex07_05/main.go" {class=line-numbers}

## 5. Side Effects in Functional Language

程式的函式有以下的行為時，就會稱該函式有 **Side Effects**。

- Reassigning a variable (val v.s. var)
- Modify a data structure in place (mutable v.s. immutable)
- Setting a field on an object (change object state)
  - 這裏指的 object 是 OOP 的 Object 不是 Scala 的 object (singleton)
  - OOP 修改物件的值，在程式語言的術語：改變物件的狀態，如上說的 **changing-state**
- Throwing an exception or halting with error
- Printing to the console or reading user input (I/O)
- Reading from or write to a file (I/O)
- Drawing on the screen (I/O)

截自 [Functional Language in Scala](http://www.amazon.com/Functional-Programming-Scala-Paul-Chiusano/dp/1617290653)

~~就實際狀況來說，我們寫程式不可能不去碰 I/O。~~
