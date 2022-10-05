# Echo

[Echo](https://echo.labstack.com/) 是 GO Webframework 之一，其他還有像 Gin, Beego 等。

- [官網](https://echo.labstack.com/)

## 第一支程式 Hello World

```go{.line-numbers}
package main

import (
 "net/http"

 "github.com/labstack/echo/v4"
)

func main() {
 e := echo.New()
 e.GET("/", func(c echo.Context) error {
  return c.String(http.StatusOK, "hello world")
 })
 e.Logger.Fatal(e.Start(":8080"))
}

```

- 首先產生 echo 物件: `e := echo.New()`
- 設定首頁的 handler function。與 Go http package 觀念相同，每個網址對應的 HTTP Method 都需要設定 handler。
  - 在 Echo 是透過 `GET`, `POST`, `PUT`, `DELETE` 等 function 的方式，指定 HTTP Method.
  - Echo 有自定義 handler function `func(c echo.Context) error`。
- 回應字串: `c.String(http.StatusOK, "hello world")`。第一個參數為 http status code。
- 啟動 Web server，監聽 `8080` port: `e.Logger.Fatal(e.Start(":8080"))`

@import "handle_request_and_response.md"
@import "bind_validate.md"
@import "cookie.md"
@import "middleware_and_custom_context.md"
