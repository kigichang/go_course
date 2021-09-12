# 17 Web (Gorilla Web Toolkit)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [17 Web (Gorilla Web Toolkit)](#17-web-gorilla-web-toolkit)
  - [0. 前言](#0-前言)
  - [1. 主程式](#1-主程式)
    - [1.1 main.go](#11-maingo)
    - [1.2 import 必要的 package](#12-import-必要的-package)
    - [1.3 實作 routing 機制](#13-實作-routing-機制)
      - [1.3.1 處理靜態資料](#131-處理靜態資料)
      - [1.3.2 動態網頁](#132-動態網頁)
    - [1.4 綁定 port 並啟動 web server](#14-綁定-port-並啟動-web-server)
  - [2. Handler 與 Request Parameter](#2-handler-與-request-parameter)
  - [3. Response Header](#3-response-header)
  - [4. Cookie](#4-cookie)
    - [4.1 Cookie Struct](#41-cookie-struct)
    - [4.2 設定/讀取 Cookie](#42-設定讀取-cookie)
  - [5. Templates](#5-templates)
    - [5.2 Template 使用說明](#52-template-使用說明)
    - [5.3 版型結構](#53-版型結構)
      - [5.3.1 layout.html](#531-layouthtml)
      - [5.3.2 nav.html](#532-navhtml)
      - [5.3.3 test.html](#533-testhtml)
  - [6. Templates 重點語法說明](#6-templates-重點語法說明)
    - [6.1 在 `` 的效果](#61-在-scriptscript-的效果)
    - [6.2 String 自動 Escape 效果](#62-string-自動-escape-效果)
    - [6.3 Compare and if-else](#63-compare-and-if-else)
    - [6.4 讀取 Array 值](#64-讀取-array-值)
    - [6.5 讀取 Map](#65-讀取-map)
    - [6.6 Array Travel (range-else)](#66-array-travel-range-else)
    - [6.7 Map Travel](#67-map-travel)
    - [6.8 Zero Value and With](#68-zero-value-and-with)
    - [6.9 在 Block 內，使用 Block 外的變數](#69-在-block-內使用-block-外的變數)
  - [7. 認識 Web 資安](#7-認識-web-資安)

<!-- /code_chunk_output -->

## 0. 前言

Go 有內建撰寫 Web Server 的套件，可以自己實作一套 AP server。因為 web 還會用到版型與靜態資料，因此在專案目錄的配置建議如下：

```text
web
├── main.go
├── public
│   └── db.png
└── templates
    ├── layout.html
    ├── nav.html
    └── test.html
```

1. public: 放置靜態資料，實際運作上，AP server 可以不用處理靜態資料。
1. templates: 放置版型

## 1. 主程式

Go 與 JAVA, PHP 等不同，不需要額外再使用 AP or Web Server，Go 內建實作 HTTP Server。

### 1.1 main.go

@import "web/main.go" {class="line-numbers"}

### 1.2 import 必要的 package

```go {.line-numbers}
import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)
```

1. `"html/template"`: Go 的 template engine。可以直接修改版型，不用重啟系統。
1. `"net/http"`: Http 模組

### 1.3 實作 routing 機制

```go {.line-numbers}
mux := http.NewServeMux()
files := http.FileServer(http.Dir("./public"))

mux.Handle("/static/", http.StripPrefix("/static/", files))
mux.HandleFunc("/", test)
mux.HandleFunc("/set_cookie", setCookie)
mux.HandleFunc("/get_cookie", getCookie)
mux.HandleFunc("/go", toGo)
```

- http.NewServMux(): 產生 `ServMux` 物件，用來處理 url routing 的工作。

#### 1.3.1 處理靜態資料

```go {.line-numbers}
files := http.FileServer(http.Dir("./public"))
mux.Handle("/static/", http.StripPrefix("/static/", files))
```

1. 測試網址：[http://127.0.0.1:8080/static/db.png](http://127.0.0.1:8080/static/db.png)
1. 指定檔案放的目錄：`http.FileServer(http.Dir("./public"))`
1. 設定 `/static/` 開始的網址，指定到剛剛設定的 `FileServer`。`/static/` 之後的目錄與檔案，都會對應該 `public/`。因此 `public/` 的檔案目錄結構要與 `/static/` 相同。

#### 1.3.2 動態網頁

其他 URL 的 routing: 利用 `HandleFunc` 來設定 URL 與處理 function 的關係。以下的 sample，`/` 會執行 `test`, `/set_cookie` 會執行 `setCookie`

```go {.line-numbers}
mux.HandleFunc("/", test)
mux.HandleFunc("/set_cookie", setCookie)
mux.HandleFunc("/get_cookie", getCookie)
mux.HandleFunc("/go", toGo)
```

### 1.4 綁定 port 並啟動 web server

```go {.line-numbers}
if err := http.ListenAndServe(":8080", mux); err != nil {
    log.Fatal(err)
}
```

1. 目前 web service 都會要求使用 HTTPS，Go 實作時，可用 `http.ListenAndServeTLS` 搭配 [CertBot](https://certbot.eff.org/) 取得憑証。
1. 在 `ListenAndServeTLS` 的 `certFile` 請用 CertBot 的 **fullchain** cert 檔案。
    - 因為在 android 上，會驗証是否是 fullchain cert。

## 2. Handler 與 Request Parameter

每一個 route 都會需要有 Handler 來處理，Handler function 的撰寫如下：

```go { .line-numbers }
func name(w http.ResponseWriter, r *http.Request) {
    body
}
```

eg:

```go { .line-numbers }
func test(w http.ResponseWriter, r *http.Request) {
    ...
}
```

1. 可透過 `r.Method` 來判斷是 GET or POST 等
1. 透過 `r.PostFormValue` 來取得 POST 值。Go 有內建多種取 request 值的方式，整理如下：

| Field | Should call method first | parameters in URL | Form | URL encoded | Multipart (upload file)
| - | - | - | - | - | -
| Form | ParseForm | ✓ | ✓ | ✓ | -
| PostForm | Form | - | ✓ | ✓ | -
| MultipartForm | ParseMultipartForm | - | ✓ | - | ✓ |
| FormValue | NA | ✓ | ✓ | ✓ | -
| PostFormValue | NA | - | ✓ | ✓ | -

from: [Go Web Programming](https://www.manning.com/books/go-web-programming)

## 3. Response Header

預設 response 的 status code 是 **200(OK)**，如果要修改 Status Code，可用 `w.WriteHeader`，eg: `w.WriteHeader(http.StatusAccepted)`。

eg:

```go {.line-numbers}
func toGo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://go.dev/")
	w.WriteHeader(http.StatusFound)
}
```

1. 如果有需要異動 HEADER 時，`w.WriteHeader` 必須要在最後使用，否則修改的 HEADER 會無效。
1. 可以試著將上述的 `toNUEiP` 內的程式順序交換，會發現回傳的 HEADER 沒有 `Location` 值。
1. 設定 Cookie 也是。因為 Cookie 是放在 HTTP 的 HEADER 內。

## 4. Cookie

### 4.1 Cookie Struct

```go {.line-numbers}
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string
    MaxAge     int
    Secure     bool
    HttpOnly   bool
    Raw        string
    Unparsed   []string
}
```

### 4.2 設定/讀取 Cookie

```go
func setCookie(w http.ResponseWriter, r *http.Request) {
    c1 := http.Cookie{
        Name:     "first_cookie",
        Value:    "Go Web Programming",
        HttpOnly: true,
    }
    c2 := http.Cookie{
        Name:     "second_cookie",
        Value:    "Manning Publications Co",
        HttpOnly: true,
    }
    http.SetCookie(w, &c1)
    http.SetCookie(w, &c2)
    w.WriteHeader(http.StatusOK)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusAccepted)
    allCookie := r.Cookies()
    firstCookie, err := r.Cookie("first_cookie")
    if err != nil {
        if errors.Is(err, http.ErrNoCookie) {
            fmt.Fprintln(w, "first cookie not found")
        } else {
            fmt.Fprintln(w, "get first cookie failure")
        }
    } else {
        fmt.Fprintln(w, "first cookie:", firstCookie)
    }

    fmt.Fprintln(w, allCookie)
}
```

1. 如果要修改回覆的 HTTP Status Code 時，`w.WriteHeader` 必須在要 `http.SetCookie` 之後。
1. 因為 Cookie 是寫在 HEADER 內。

## 5. Templates

Go template engine 很好用，會自動依版型的內容，來自動做 escape 動作。使用 template engine 需要再學習它的語法。

Go html template 是利用 text template，因此相關的語法，要看 ["text/template"](https://golang.org/pkg/text/template/#pkg-index)

### 5.2 Template 使用說明

```go { .line-numbers }
func generateHTML(w http.ResponseWriter, data interface{}, files ...string) {
    var tmp []string
    for _, f := range files {
        tmp = append(tmp, fmt.Sprintf("templates/%s.html", f))
    }

    tmpl := template.Must(template.ParseFiles(tmp...))
    tmpl.ExecuteTemplate(w, "layout", data)
}
```

1. `template.ParseFiles(tmp...)`: 選擇會用到的版型檔案，要確認版型的路徑與檔案是否正確，並取得版型物件
1. `tmpl.ExecuteTemplate(w, "layout", data)`: 執行版型，並將版型會用的資料(`data`)帶入。其中 **layout** 是定義在版型內，指定要從那個區塊開始執行。

    - 在同一個版型檔案內，可以定義多個模版區塊 (也就是可以使用多組 `{{ define "NAME" }}`)，在 `ExecuteTemplate` 時，可以指定要使用那個區塊。

### 5.3 版型結構

範例的版型結構如下：

#### 5.3.1 layout.html

__layout.html__  是版型的主框。內含 __nav__ `{{ template "navbar" . }}` 與  __content__ `{{ template "content" . }}` 這個子版型。

```html { .line-numbers}
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=9">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Go Course Web - {{ .Title }}</title>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    </head>
    <body>
        {{ template "navbar" . }}
        <div class="container">
            {{ template "content" .Data }}
        </div>
        <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
        <script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
        <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
    </body>
</html>
{{ end }}
```

1. 在 __layout.html__ 定義了這個版型的名稱 __layout__:`{{ define "layout" }}`，也就是程式碼 `tmpl.ExecuteTemplate(w, "layout", data)` 中的 `"layout"`。

1. 在 include 子版型的語法中，eg: `{{ template "navbar" . }}`，有 **`.`**，是指由 `ExecuteTemplate` 傳入的資料。在 ["text/template"](https://golang.org/pkg/text/template/#pkg-index) 有詳細的說明。

#### 5.3.2 nav.html

__nav.html__: Navigation bar。跟 __layout.html__ 一樣，一開頭定義這個版型的名稱 `{{ define "navbar" }}`，也就是 __layout.html__ 中 `{{ template "navbar" . }}` 的 `"navbar"`。

```html { .line-numbers }
{{ define "navbar" }}
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand" href="#">{{.Title}}</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNavDropdown">
        <ul class="navbar-nav">
            <li class="nav-item {{ if eq .Nav "home" }}active{{ end }}">
                <a class="nav-link" href="/">首頁</a>
            </li>
            <li class="nav-item {{ if eq .Nav "add" }}active{{ end }}">
                <a class="nav-link" href="/add">新增</a>
            </li>
            <li class="nav-item {{ if eq .Nav "all" }}active{{ end }}">
                <a class="nav-link" href="/all">看全部</a>
            </li>
        </ul>
    </div>
</nav>
{{ end }}
```

#### 5.3.3 test.html

__test.html__ 是內容的子版型，開頭 `{{ define "content" }}` 與上述相同。

```html
{{ define "content" }}

<script languate="javascript">
    var pair = {{ .TestStruct }};
    var array = {{ .TestArray }};
    var str1 = "{{ .TestString }}";
    var str2 = "{{ .SimpleString }}";
</script>

<p>有特殊字元<br />
    {{ .TestString }} <br />
    <a title='{{ .TestString }}'>{{ .TestString }}</a> <br />
    <a href="/{{ .TestString }}">{{ .TestString }}</a> <br />
    <a href="?q={{ .TestString }}">{{ .TestString }}</a> <br />
    <a onx='f("{{ .TestString }}")'>{{ .TestString }}</a> <br />
    <a onx='f({{ .TestString }})'>{{ .TestString }}</a> <br />
    <a onx='pattern = /{{ .TestString }}/;'>{{.TestString }}</a> <br />
</p>

<p>無特殊字元<br />
    {{ .SimpleString }} <br />
    <a title='{{ .SimpleString }}'>{{ .SimpleString }}</a> <br />
    <a href="/{{ .SimpleString }}">{{ .SimpleString }}</a> <br />
    <a href="?q={{ .SimpleString }}">{{ .SimpleString }}</a> <br />
    <a onx='f("{{ .SimpleString }}")'>{{ .SimpleString }}</a> <br />
    <a onx='f({{ .SimpleString }})'>{{ .SimpleString }}</a> <br />
    <a onx='pattern = /{{ .SimpleString }}/;'>{{ .SimpleString }}</a> <br />
</p>

<p>Compare<br />
    {{ with . }}
    {{ if eq .Num1 .Num2 }} eq {{ else }} ne {{end}}<br/>
    {{ if ne .Num1 .Num2 }} ne {{ else }} eq {{end}}<br/>
    {{ if lt .Num1 .Num2 }} lt {{ else if gt .Num1 .Num2 }} gt {{ else if le .Num1 .Num2 }} le {{ else }} ge {{end}}
    {{ end }}
</p>

<p>Array Index<br />
    {{ index .TestArray 0 }}<br />
    {{ index .TestArray 1 }}<br />
    {{ index .TestArray 2 }}<br />
    len: {{ len .TestArray }}<br />
</p>

<p>Map<br />
    abc : {{ index .TestMap "abc" }} <br />
    A : {{ index .TestMap "A" }} <br />
    B : {{ index .TestMap "B" }} <br />
</p>

<p>Range Array 1<br />
    {{ range .TestArray}}
        {{.}} <br />
    {{ else }}
        No Data
    {{ end }}
    <br />
</p>

<p>Range Array 2<br />
    {{ range $idx, $elm := .TestArray }}
        {{ $idx }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>

<p>Range Empty<br />
    {{ range .EmptyArray}}
        {{ . }} <br />
    {{ else }}
        No Data
    {{ end }}
    <br />
</p>

<p>Range Map<br />
    {{ range $key, $elm := .TestMap }}
        {{ $key }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>

<p>With Empty<br />
    {{with .EmptyArray}}
        Have Data
    {{else}}
        No Data
    {{end}}
</p>

<p>With Int Zero Value<br />
    {{with .ZeroInt}}
        Have Data
    {{else}}
        No Data
    {{end}}
</p>

<p>Reference Outside Variable<br />
    {{ range $idx, $elm := .TestArray }}
        {{$.SimpleString}} {{ $idx }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>
{{ end }}
```

## 6. Templates 重點語法說明

Go template engine 會依照版型的內容，自動作 Escape 處理，避免 XSS 攻擊。以下整理了常用的案例。

### 6.1 在 `<script></script>` 的效果

語法：

```html
<script languate="javascript">
    var pair = {{ .TestStruct }};
    var array = {{ .TestArray }};
    var str1 = "{{ .TestString }}";
    var str2 = "{{ .SimpleString }}";
</script>
```

結果：

```html
<script languate="javascript">
    var pair = {"A":"foo","B":"boo"};
    var array = ["Hello","World","Test"];
    var str1 = "O\x27Reilly: How are \x3ci\x3eyou\x3c\/i\x3e?";
    var str2 = "中文測試";
</script>
```

1. 如果 string 內容有需要做 escape 時，go template engine 會自動做，eg: `"O\x27Reilly: How are \x3ci\x3eyou\x3c\/i\x3e?"`。
1. 如果資料是 struct 或 array，會自動轉成 javascript 的 data type 型態。eg: `var pair = {"A":"foo","B":"boo"};` 及 `var array = ["Hello","World","Test"];`。

### 6.2 String 自動 Escape 效果

語法：

```html
<p>有特殊字元<br />
    {{ .TestString }} <br />
    <a title='{{ .TestString }}'>{{ .TestString }}</a> <br />
    <a href="/{{ .TestString }}">{{ .TestString }}</a> <br />
    <a href="?q={{ .TestString }}">{{ .TestString }}</a> <br />
    <a onx='f("{{ .TestString }}")'>{{ .TestString }}</a> <br />
    <a onx='f({{ .TestString }})'>{{ .TestString }}</a> <br />
    <a onx='pattern = /{{ .TestString }}/;'>{{.TestString }}</a> <br />
</p>

<p>無特殊字元<br />
    {{ .SimpleString }} <br />
    <a title='{{ .SimpleString }}'>{{ .SimpleString }}</a> <br />
    <a href="/{{ .SimpleString }}">{{ .SimpleString }}</a> <br />
    <a href="?q={{ .SimpleString }}">{{ .SimpleString }}</a> <br />
    <a onx='f("{{ .SimpleString }}")'>{{ .SimpleString }}</a> <br />
    <a onx='f({{ .SimpleString }})'>{{ .SimpleString }}</a> <br />
    <a onx='pattern = /{{ .SimpleString }}/;'>{{ .SimpleString }}</a> <br />
</p>
```

結果：

```html
<p>有特殊字元<br />
    O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;? <br />
    <a title='O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?'>O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
    <a href="/O%27Reilly:%20How%20are%20%3ci%3eyou%3c/i%3e?">O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
    <a href="?q=O%27Reilly%3a%20How%20are%20%3ci%3eyou%3c%2fi%3e%3f">O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
    <a onx='f("O\x27Reilly: How are \x3ci\x3eyou\x3c\/i\x3e?")'>O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
    <a onx='f(&#34;O&#39;Reilly: How are \u003ci\u003eyou\u003c/i\u003e?&#34;)'>O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
    <a onx='pattern = /O\x27Reilly: How are \x3ci\x3eyou\x3c\/i\x3e\?/;'>O&#39;Reilly: How are &lt;i&gt;you&lt;/i&gt;?</a> <br />
</p>

<p>無特殊字元<br />
    中文測試 <br />
    <a title='中文測試'>中文測試</a> <br />
    <a href="/%e4%b8%ad%e6%96%87%e6%b8%ac%e8%a9%a6">中文測試</a> <br />
    <a href="?q=%e4%b8%ad%e6%96%87%e6%b8%ac%e8%a9%a6">中文測試</a> <br />
    <a onx='f("中文測試")'>中文測試</a> <br />
    <a onx='f(&#34;中文測試&#34;)'>中文測試</a> <br />
    <a onx='pattern = /中文測試/;'>中文測試</a> <br />
</p>
```

1. Go template 會依據字串在版型內的型別，做相對應的 escape。如：url escape or html escape。

### 6.3 Compare and if-else

語法：

```html
<p>Compare<br />
    {{ with . }}
    {{ if eq .Num1 .Num2 }} eq {{ else }} ne {{end}}<br/>
    {{ if ne .Num1 .Num2 }} ne {{ else }} eq {{end}}<br/>
    {{ if lt .Num1 .Num2 }} lt {{ else if gt .Num1 .Num2 }} gt {{ else if le .Num1 .Num2 }} le {{ else }} ge {{end}}
    {{ end }}
</p>
```

結果：

```html
<p>Compare<br />

    ne <br/>
    ne <br/>
    lt

</p>
```

### 6.4 讀取 Array 值

語法：

```html
<p>Array Index<br />
    {{ index .TestArray 0 }}<br />
    {{ index .TestArray 1 }}<br />
    {{ index .TestArray 2 }}<br />
    len: {{ len .TestArray }}<br />
</p>
```

結果：

```html
<p>Array Index<br />
    Hello<br />
    World<br />
    Test<br />
    len: 3<br />
</p>
```

### 6.5 讀取 Map

語法：

```html
<p>Map<br />
    abc : {{ index .TestMap "abc" }} <br />
    A : {{ index .TestMap "A" }} <br />
    B : {{ index .TestMap "B" }} <br />
</p>
```

結果：

```html
<p>Map<br />
    abc : DEF <br />
    A : B <br />
    B :  <br />
</p>
```

### 6.6 Array Travel (range-else)

語法：

```html
<p>Range Array 1<br />
    {{ range .TestArray}}
        {{.}} <br />
    {{ else }}
        No Data
    {{ end }}
    <br />
</p>

<p>Range Array 2<br />
    {{ range $idx, $elm := .TestArray }}
        {{ $idx }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>

<p>Range Empty<br />
    {{ range .EmptyArray}}
        {{ . }} <br />
    {{ else }}
        No Data
    {{ end }}
    <br />
</p>
```

結果：

```html
<p>Range Array 1<br />

        Hello <br />

        World <br />

        Test <br />

    <br />
</p>

<p>Range Array 2<br />

        0 : Hello <br />

        1 : World <br />

        2 : Test <br />

    <br />
</p>

<p>Range Empty<br />

        No Data

    <br />
</p>
```

### 6.7 Map Travel

語法：

```html
<p>Range Map<br />
    {{ range $key, $elm := .TestMap }}
        {{ $key }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>
```

結果：

```html
<p>Range Map<br />

        A : B <br />

        abc : DEF <br />

    <br />
</p>
```

### 6.8 Zero Value and With

確認值是否**不是 zero value**，要特別小心當值是 **zero value**，像 `int` 型別，值又是 **"0"** 時，會判定成沒有值，會進到 `else` 的區塊。

語法：

```html
<p>With Empty<br />
    {{with .EmptyArray}}
        Have Data
    {{else}}
        No Data
    {{end}}
</p>

<p>With Int Zero Value<br />
    {{with .ZeroInt}}
        Have Data
    {{else}}
        No Data
    {{end}}
</p>
```

結果：

```html
<p>With Empty<br />

        No Data

</p>

<p>With Int Zero Value<br />

        No Data

</p>
```

### 6.9 在 Block 內，使用 Block 外的變數

在 `with` or `range` 內，需要使用到外部的變數時，可以用 `$` 來讀取外部值。`$` 是指傳入該版型的資料。以 `test.html` 來說，`$` 代表的是 `MyData.Data`。

語法：

```html
<p>Reference Outside Variable<br />
    {{ range $idx, $elm := .TestArray }}
        {{$.SimpleString}} {{ $idx }} : {{ $elm }} <br />
    {{else}}
        No Data
    {{end}}
    <br />
</p>
```

結果：

```html
<p>Reference Outside Variable<br />

        中文測試 0 : Hello <br />

        中文測試 1 : World <br />

        中文測試 2 : Test <br />

    <br />
</p>
```

## 7. 認識 Web 資安

- [身為 Web 工程師，你一定要知道的幾個 Web 資訊安全議題](https://medium.com/starbugs/%E8%BA%AB%E7%82%BA-web-%E5%B7%A5%E7%A8%8B%E5%B8%AB-%E4%BD%A0%E4%B8%80%E5%AE%9A%E8%A6%81%E7%9F%A5%E9%81%93%E7%9A%84%E5%B9%BE%E5%80%8B-web-%E8%B3%87%E8%A8%8A%E5%AE%89%E5%85%A8%E8%AD%B0%E9%A1%8C-29b8a4af6e13)
- [讓我們來談談 CSRF](https://blog.techbridge.cc/2017/02/25/csrf-introduction/)
- [Hacksplaining](https://www.hacksplaining.com/lessons)
