## Handle Request and Response

### Handler

Echo 處理 request 上，直接用 `GET`, `POST`, `PUT`, `DELETE` 等 function ，對應 HTTP Method，也可以使用 `Any` 方法，不限 http method。設定 HTTP Handler。如下：

```go {.line-numbers}
e.GET("/users", func(c echo.Context) error {
	...
})

e.POST("/users", func(c echo.Context) error {
	...
})
```

Echo 的 Handler 與 Go 的 http Handler 不同，是 `func (c echo.Context) error`。


### Request
Echo 提拱以下方式，處理 request：

- QueryParam: 用在 Query String 上。
- FromValue: 用在 Form Submit 上，如 `x-www-form-urlencoded` or `multipart/form-data`。
- Param: 用在 URL PATH 上。

#### QueryParam

用在 URL 的 Query String, Query String 是指 URL 上 `?` 之後的參數值。如: `https://www.google.com.tw/search?q=NuEiP&source=hp`， Query String 部分是 `?q=NuEiP&source=hp`。一般常見 Query String 會出現在 **GET** method 上，但也可用在其他方法上。如下：

```go{.line-numbers}
e.Any("/test", func(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParam("a"))
})
```

使用 CURL 測試: `curl --include -X POST http://127.0.0.1:8080/test?a=100`。透過 `QueryParam` 方法，取得參數 `a` 的值。

#### FromValue

與 Go 相同，處理 `x-www-form-urlencoded` 或 `multipart/form-data` 資料。如下：

```go{.line-numbers}
e.POST("/users", func(c echo.Context) error {
	name := strings.TrimSpace(c.FormValue("name"))
	email := strings.TrimSpace(c.FormValue("email"))
	...
})
```

使用 CURL 測試: `curl -v -d "name=Joe" -d "email=joe@labstack.com" http://127.0.0.1:8080/users`。

#### Param

用在 URL Path 上，URL Path 是指 Domain 之後，Query String 之前的字串，如檔案路徑字串。如: `http://127.0.0.1:8080/user/1?a=100` 中的 `/user/1`。在 Echo 可以使用 `:` 來為 Path 命名變數。如上的 `/user/1`，可以設計成 `/user/:id`，透過 `Param` 方法，取得 `id` 值。如下：

```go{.line-numbers}
e.GET("/user/:id", func(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	...
})
```

使用 CURL 測試: `curl -v http://127.0.0.1:8080/user/1`

### Response

- Context.String: 回傳一般字串，Content-Type 為 `text/plain; charset=UTF-8`。
- Context.Json/Context.JSONPretty: 回傳 JSON 結構資料，Content-Type 為 `application/json; charset=UTF-8`。
- Context.HTML: 回傳 html 網頁，Content-Type 為 `text/html; charset=UTF-8`。
- Context.Render: 自定義版型 Engine，回傳 html 網頁。 Content-Type 為 `text/html; charset=UTF-8`。

每種回傳資料的方法，第一個參數都是指定 [HTTP Status Code](https://golang.org/src/net/http/status.go)。

[HTTP Status Code 說明](https://developer.mozilla.org/zh-TW/docs/Web/HTTP/Status)

#### String

回傳一般字串。如下：

```go {.line-numbers}
e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
})
```

#### JSON

回傳 JSON 格式。如下：

```go {line-numbers}
e.GET("/render", func(c echo.Context) error {
	...
	return c.JSON(http.StatusOK, data[1])
})
```

或使用 `JSONPretty` 整理成比較 friendly 呈現方式。如下：

```go{.line-numbers}
e.GET("/user/:id", func(c echo.Context) error {
	...
	rreturn c.JSONPretty(
		http.StatusOK,
		echo.Map{
			"code": "0000",
			"data": user,
	}, "\t")
})
```

#### Render

自定義版型 Engine，需實作 `echo.Renderer` interface。 `echo.Renderer` interface 如下：

```go{.line-numbers}
// Renderer is the interface that wraps the Render function.
Renderer interface {
	Render(io.Writer, string, interface{}, Context) error
}
```

實作步驟：

1. 實作 `echo.Renderer`，如下 (see render_engine.go)：

	```go {.line-numbers}
	//go:embed templates
	var content embed.FS

	type RenderEngine struct {
		tmpl *template.Template
	}

	func (e *RenderEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
		return e.tmpl.ExecuteTemplate(w, name, data)
	}

	func newRenderEngine() *RenderEngine {
		return &RenderEngine{
			tmpl: template.Must(template.ParseFS(content, "templates/*.html")),
		}
	}
	```

1. 設定 echo `Renderer`，如下：

	```go {.line-numbers}
	e := echo.New()
	e.Renderer = newRenderEngine()
	```

1. 使用 `Context.Render` 回傳 html。

```go {.line-numbers}
e.GET("/render", func(c echo.Context) error {
	switch c.QueryParam("tmpl") {
	case "simple":
		return c.Render(http.StatusOK, "simple", data[1])
	case "table":
		return c.Render(http.StatusOK, "table", data[1])
	default:
		return c.JSON(http.StatusOK, data[1])
	}
})
```

測試：

- `curl -v http://127.0.0.1:8080/render`
- `curl -v http://127.0.0.1:8080/render?tmpl=table`
- `curl -v http://127.0.0.1:8080/render?tmpl=simple`
