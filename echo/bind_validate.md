## Bind

Echo 提供一個簡便的方式，將 Query String, Form, JSON 等資料，透過 Go Structure 取得，一方面便利讀取資料，也可以設定檢查絛件，加速開發。切記，任何由 Client 端傳來的資料，都需要經過檢查後，才能使用(如:寫入資料庫)。

### Structure Bind

透過 Structure Tag 的方式，設定 Structure 每個 Field 的資料來源。如下：

```go {.line-numbers}
type User struct {
	ID    int    `json:"id" param:"id"`
	Name  string `json:"name" query:"name" form:"name"`
	Email string `json:"email" query:"email" form:"email"`
}
```

Structure Tag 說明:

- json: 使用 json 格式上傳的資料
- param: URL PATH 的參數值
- query: Query string 參數值
- form: form submit 參數值 

使用方式如下：

```go {.line-numbers}
g.PUT("/user/:id", func(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	data[user.ID] = user

	return c.JSONPretty(http.StatusOK, echo.Map{
		"cdoe": "0000",
		"data": user,
	}, "\t")
})
```

- User.ID: 有設定 `param:id`，因此可以取得 `/user/:id` 上的 **id** 值。

如果 request 同時有相同的參數值傳入時，比如: `curl --include -X PUT -d "name=abc" -d "email=abc@gmail" http://127.0.0.1:8080/api/v1/user/1\?name\="aaa"`，同時利用 form 與 query string 傳入 name。 Echo Bind 的順序會是先 Body (json, xml, form)，然後是 Query String (如果 method 是 GET/DELETE)，最後是 URL PATH。因此上述的例子，`name` 最後的值會是 `abc` 而不是 `aaa`。


以下是 Echo 的 Bind source code.

```go{.line-numbers}
// Bind implements the `Binder#Bind` function.
// Binding is done in following order: 1) path params; 2) query params; 3) request body. Each step COULD override previous
// step binded values. For single source binding use their own methods BindBody, BindQueryParams, BindPathParams.
func (b *DefaultBinder) Bind(i interface{}, c Context) (err error) {
	if err := b.BindPathParams(c, i); err != nil {
		return err
	}
	// Issue #1670 - Query params are binded only for GET/DELETE and NOT for usual request with body (POST/PUT/PATCH)
	// Reasoning here is that parameters in query and bind destination struct could have UNEXPECTED matches and results due that.
	// i.e. is `&id=1&lang=en` from URL same as `{"id":100,"lang":"de"}` request body and which one should have priority when binding.
	// This HTTP method check restores pre v4.1.11 behavior and avoids different problems when query is mixed with body
	if c.Request().Method == http.MethodGet || c.Request().Method == http.MethodDelete {
		if err = b.BindQueryParams(c, i); err != nil {
			return err
		}
	}
	return b.BindBody(c, i)
}
```

### Binder





### Validation
