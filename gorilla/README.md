
# Build Web with Gorilla Toolkit

- [mux](https://github.com/gorilla/mux): Mux Router，可以定義更彈性的 routing path。
- [securecookie](https://github.com/gorilla/securecookie): 加密 cookie。
- [schema](https://github.com/gorilla/schema): 將 post form 的資料，轉成 struct。
- [csrf](https://github.com/gorilla/csrf): 避免被 CSRF 攻擊[^csrf]。

[^csrf]: [讓我們來談談 CSRF](https://blog.techbridge.cc/2017/02/25/csrf-introduction/)

將綜合以上與 Gorilla Tool Kit，撰寫登入功能。

## gorilla_web/main.go

```go {.line-numbers}
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"

    "github.com/gorilla/csrf"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
    "github.com/gorilla/securecookie"
)

type member struct {
    Email string `json:"email"`
}

func (m *member) String() string {
    memBytes, err := json.Marshal(m)
    if err != nil {
        return err.Error()
    }
    return string(memBytes)
}

func generateHTML(w http.ResponseWriter, data interface{}, files ...string) {
    var tmp []string
    for _, f := range files {
        tmp = append(tmp, fmt.Sprintf("templates/%s.html", f))
    }

    tmpl := template.Must(template.ParseFiles(tmp...))
    tmpl.ExecuteTemplate(w, "layout", data)
}

func redirect(w http.ResponseWriter, target string) {
    w.Header().Set("Location", target)
    w.WriteHeader(http.StatusFound)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "index")
}

func showLogin(w http.ResponseWriter, r *http.Request) {
    generateHTML(w, csrf.TemplateField(r), "layout", "login")
}

func doLogin(w http.ResponseWriter, r *http.Request) {
    form := struct {
        Email    string `schema:"email"`
        Password string `schema:"password"`
        Token    string `schema:"auth_token"`
    }{}

    r.ParseForm()

    if err := schema.NewDecoder().Decode(&form, r.PostForm); err != nil {
        log.Println("schema decode:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    mem := &member{
        Email: form.Email,
    }

    tmp, err := secureC.Encode(cookieName, mem)
    if err != nil {
        log.Println("encode secure cookie:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    cookie := &http.Cookie{
        Name:   cookieName,
        Value:  tmp,
        MaxAge: 0,
        Path:   "/",
    }

    http.SetCookie(w, cookie)
    redirect(w, "/member")
}

func logout(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:   cookieName,
        Value:  "",
        MaxAge: -1,
    }

    http.SetCookie(w, cookie)
    redirect(w, "/")
}

func showRegister(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "show register")
}

func doRegister(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "do resgier")
}

func memberIndex(w http.ResponseWriter, r *http.Request) {
    mem, ok := r.Context().Value(ctxKey(cookieName)).(*member)
    if !ok || mem == nil {
        log.Println(mem, ok)
        redirect(w, "/")
        return
    }

    fmt.Fprintln(w, "member:", mem)
}

func memberShowEdit(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "member show edit")
}

func memberDoEdit(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "member do edit")
}

func memberAuthHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // check cookie
        cookie, err := r.Cookie(cookieName)
        if err != nil {
            w.WriteHeader(http.StatusForbidden)
            return
        }

        value := &member{}
        if err := secureC.Decode(cookieName, cookie.Value, value); err != nil {
            log.Println("decode secure cookie:", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        newRequest := r.WithContext(context.WithValue(r.Context(), ctxKey(cookieName), value))

        next.ServeHTTP(w, newRequest)
    })
}

type ctxKey string

var (
    secureC *securecookie.SecureCookie
)

const (
    hashKey  = "1234567890123456789012345678901234567890123456789012345678901234"
    blockKey = "0123456789abcdef"

    cookieName = "mytest"
)

func main() {

    secureC = securecookie.New([]byte(hashKey), []byte(blockKey))
    r := mux.NewRouter()

    r.HandleFunc("/", index)
    r.HandleFunc("/login", showLogin).Methods("GET")
    r.HandleFunc("/login", doLogin).Methods("POST")
    r.HandleFunc("/logout", logout)
    r.HandleFunc("/register", showRegister).Methods("GET")
    r.HandleFunc("/register", doRegister).Methods("POST")

    s := r.PathPrefix("/member").Subrouter()
    s.HandleFunc("", memberIndex)
    s.HandleFunc("/edit", memberShowEdit).Methods("GET")
    s.HandleFunc("/edit", memberDoEdit).Methods("POST")

    s.Use(memberAuthHandler)

    CSRF := csrf.Protect(
        []byte(`1234567890abcdefghijklmnopqrstuvwsyz!@#$%^&*()_+~<>?:{}|,./;'[]\`),
        csrf.RequestHeader("X-ATUH-Token"),
        csrf.FieldName("auth_token"),
        csrf.Secure(false),
    )

    log.Fatal(http.ListenAndServe(":8080", CSRF(r)))
}
```

## mux

```go {.line-numbers}
r := mux.NewRouter()

r.HandleFunc("/", index)
r.HandleFunc("/login", showLogin).Methods("GET")
r.HandleFunc("/login", doLogin).Methods("POST")
r.HandleFunc("/logout", logout)
r.HandleFunc("/register", showRegister).Methods("GET")
r.HandleFunc("/register", doRegister).Methods("POST")

s := r.PathPrefix("/member").Subrouter()
s.HandleFunc("", memberIndex)
s.HandleFunc("/edit", memberShowEdit).Methods("GET")
s.HandleFunc("/edit", memberDoEdit).Methods("POST")

s.Use(memberAuthHandler)
```

- `Methods(xxx)`: 限定某種 Http Method。
- `r.PathPrefix(xxx).Subrouter()`: 產生子 router, 方便管理子模組。
- `s.Use(xxx)`: 此 router 下的所有 request 都需要先經過某個 handler 處理，類似 Java Servlet Filter 功能。

## csrf

1. 設定:

    ```go {.line-numbers}
    CSRF := csrf.Protect(
        []byte(`1234567890abcdefghijklmnopqrstuvwsyz!@#$%^&*()_+~<>?:{}|,./;'[]\`),
        csrf.RequestHeader("X-ATUH-Token"),
        csrf.FieldName("auth_token"),
        csrf.Secure(false),
    )

    log.Fatal(http.ListenAndServe(":8080", CSRF(r)))
    ```

1. 用 `csrf.TemplateField(r)` 產生 token 並傳給版型

    ```go {.line-numbers}
    func showLogin(w http.ResponseWriter, r *http.Request) {
        generateHTML(w, csrf.TemplateField(r), "layout", "login")
    }
    ```

1. 將 token 放在 html form 內 (login.html)

    ```html
    {{ define "content" }}
    <form method="post" action="/login">
    {{ . }}
        <div class="form-group row">
        <label for="email" class="col-sm-2 col-form-label">Email:</label>
        <div class="col-sm-10">
            <input type="email" class="form-control" id="email" name="email" required>
        </div>
        </div>
        <div class="form-group row">
            <label for="password" class="col-sm-2 col-form-label">Password</label>
            <div class="col-sm-10">
            <input type="password" class="form-control" id="password" name="password" required>
            </div>
        </div>
        <div class="form-group row">
        <div class="col-sm-10">
            <button type="submit" class="btn btn-primary">Submit</button>
        </div>
        </div>
    </form>
    {{ end }}
    ```

    - 輸出的結果，會在 form 加一個 hidden 的 field，name 為 **auth_token**。

## schema

1. 用 Gorilla schema 處理時 crsf，記得要加一個 token 欄位，可以不處理

    ```go {.line-number}
    form := struct {
        Email    string `schema:"email"`
        Password string `schema:"password"`
        Token    string `schema:"auth_token"`
    }{}

    r.ParseForm()

    err := schema.NewDecoder().Decode(&form, r.PostForm)
    ```

## securecookie

1. Initialize

    ```go {.line-number}
    secureC = securecookie.New([]byte(hashKey), []byte(blockKey))
    ```

    - hashKey: 32 or 64 bytes
    - blockKey: 16 (AES-128), 24 (AES-192), 32 (AES-256) bytes

1. Encode and Set Cookie

    ```go {.line-numbers}
    tmp, err := secureC.Encode(cookieName, mem)
    if err != nil {
        log.Println("encode secure cookie:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    cookie := &http.Cookie{
        Name:   cookieName,
        Value:  tmp,
        MaxAge: 0,
        Path:   "/",
    }

    http.SetCookie(w, cookie)
    ```

    - cookieName: cookie name
    - value: 可以是字串，也可以是 struct。

1. Read Cookie and Decode

    ```go {.line-numbers}
    cookie, err := r.Cookie(cookieName)
    if err != nil {
        w.WriteHeader(http.StatusForbidden)
        return
    }

    value := &member{}
    if err := secureC.Decode(cookieName, cookie.Value, value); err != nil {
        log.Println("decode secure cookie:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    ```

    - cookieName: cookie name
    - value: 原先寫在 cookie 的值

## MiddlewareFunc

Gorilla mux 提供 MiddlewareFunc，讓所有請求某個 URL 時，必須經過此 func 處理後，才會進入該 URL 的 handler func。作用類似 Java Servlet Filter。通常會在這裏面，作身分驗證，並把用戶的資料，變成 request-scope 的資料。

處理方式：讀 Cookie -> 解析 Cookie 資料 -> 轉換成 Reqeust-Scope 資料 -> Next Handler。在 Go，可以使用 Context 的方式，儲存 request-scope 資料。

```go {.line-numbers}
func memberAuthHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // check cookie
        cookie, err := r.Cookie(cookieName)
        if err != nil {
            w.WriteHeader(http.StatusForbidden)
            return
        }

        value := &member{}
        if err := secureC.Decode(cookieName, cookie.Value, value); err != nil {
            log.Println("decode secure cookie:", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        newRequest := r.WithContext(context.WithValue(r.Context(), ctxKey(cookieName), value))

        next.ServeHTTP(w, newRequest)
    })
}
```

讀取用戶資料時，就不用再重新讀 Cookie。eg:

```go {.line-numbers}
func memberIndex(w http.ResponseWriter, r *http.Request) {
    mem, ok := r.Context().Value(ctxKey(cookieName)).(*member)
    if !ok || mem == nil {
        log.Println(mem, ok)
        redirect(w, "/")
        return
    }

    fmt.Fprintln(w, "member:", mem)
}
```
