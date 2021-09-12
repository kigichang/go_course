# 18 RESTful

## REST API

### What is REST

Representational State Transfer (REST) 由 Roy Thomas Fielding 在 2000 年在 [Architectural Styles and
the Design of Network-based Software Architectures](https://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm) 這篇論文中提出。

REST 是一種設計風格，並非一種標準。通常使用 HTTP, URI, XML, HTML 來實作，近來 XML 的部分，改用了 JSON，或 XML 與 JSON 兩者並存。

### Architectural Goals

#### Performance

#### Scalability

- scaling up - increasing the capacity of services, consumers, and network devices
- scaling out - distributing load across services and programs

#### Simplicity

#### Modifiability

可以應需求改變，而可輕易更動

#### Visibility

客戶與伺服器端，或伺服器之間，可加入 middleware 來監控或調適彼此間的溝通

#### Portability

服務可以輕易地被佈暑

#### Reliability

### Architectural Constrains

#### Client-Server

- 滿足 RESTful 設計的系統，需要是 Client-Server 的架構 (客戶端與伺服器端)
- RESTful 只規範通訊協定

![Client-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/client_server_style.gif)

#### Layered System

- 客戶端無法指定每次要連線的伺服器，也無法指定伺服器間的路由

![Uniform-Layered-Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/layered_uccss.gif)

#### Cache

- 系統內的所有資料，都需定義是否能使用快取 (cache)
- 客戶端能建立並使用本地的快取，如伺服端沒有更動，客戶端則使用本地的快取資料
- 可以節省伺服器端的資源及網路頻寬
- 通常不變的資料，適用快取，但也不是所有的資料都合適
- Cache 也能應用在伺服器端

![Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/ccss_style.gif)

#### Stateless

- RESTful 規範客戶端與伺服器端的溝通是 Stateless。但伺服器以後的設計則沒有限制，也就是說伺服器端以後的部分可以使用像 memcached, Redis 等來存放狀態(如 session)資訊
- 由於 RESTful 是 Stateless, 在 API 的設計上，需要符合 Atomic 特性。簡單來說，atomic 是指不能用二個以上的 API 來完成一個動作。如轉帳，需要扣 A 帳號加 B 帳號，在設計時，就不要拆成兩個 API 以免發生扣 A 但沒加 B。
- 每次 API 的呼叫，都要傳足夠且完整的資料給伺服器端，不要認為伺服器端會記錄現在的使用者資訊。
- Idempotent 特性，簡單來說，客戶端可以重覆呼叫 API，在伺服器端要可以處理這種情境。一般說來，如果是讀取未被更動的資料，每次的呼叫，回傳結果要為一致；但更新與刪除資料則不一定。因此在設計 RESTful API 時，要將重覆呼叫的因素考量進去。

![Client-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/stateless_cs.gif)

#### Uniform Interface

- 介面上定義應該要基於資源，而不是動作
- 資源通常有四種動作：新增(Create)，讀取(Read)，更新(Update)，刪除(Delete) (CRUD)

![Uniform-Client-Cache-Stateless-Server](https://www.ics.uci.edu/~fielding/pubs/dissertation/uniform_ccss.gif)

#### Code-On-Demand (Optional)

簡單來說，伺服器端可以回傳一個指定或一段程式碼 (如: javascript) 讓客戶端來執行

![Restful](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_style.gif)

### RESTful Interface Example

#### Use HTTP Response Codes to Indicate Status

常用的 10 種 HTTP Status Code

- 200 OK

  最常用回覆成功的狀態碼

- 201 CREATED

  用在新增資料成功 (透過 PUT or POST 方法)。也可以在 Header 的 Location 加入新資料的連結

- 204 NO CONTENT

  回覆處理成功，但不會回覆訊息，常用在 DELETE and PUT

- 400 BAD REQUEST

  表示請求錯誤。大都情形下的錯誤，都可以回覆此狀態碼。

- 401 UNAUTHORIZED

  認証失敗

- 403 FORBIDDEN

  未授權

- 404 NOT FOUND

  資源不存在

- 405 METHOD NOT ALLOWED

  存取資源的方法不支援，比如說： POST _/users/123_，在新增用戶時，不能指定用戶的編號，此時就可以回覆 405

- 409 CONFLICT

  更新資源發生衝突，比如說：在用戶資料中，假設 email 是唯一值，當有兩筆用戶資料，填入同一個 email 時，就可以回覆 409。

- 500 INTERNAL SERVER ERROR

  發生異外錯誤時，通常都是系統發生 Exception 時，回覆 500

#### Using HTTP Methods for RESTful Services

- POST

  用於新增資源；新增成功時，可以回覆 201，並且在 Header 的 Location 回傳新資源的連結 (內含新資源的 ID, 比如說：_/users/1234_)

  Examples:

  - POST [http://www.example.com/customers]([http://www.example.com/customers)
  - POST [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)

- GET

  用於讀取資源，且在資源未更新前，每次讀取，回覆的資料都應一致(Idempotent)。使用 GET 時，不要去新增/更新資源。

  Examples:

  - GET [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - GET [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)
  - GET [http://www.example.com/buckets/sample](http://www.example.com/buckets/sample)

- PUT

  通常用於更新資料，但也可以用在新增。與 POST 不同的是，POST 在新增時，Client 不會指定要新增的 ID，但 PUT 會指定新增的 ID。

  Examples:

  - PUT [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - PUT [http://www.example.com/customers/12345/orders/98765](http://www.example.com/customers/12345/orders/98765)
  - PUT [http://www.example.com/buckets/secret_stuff](http://www.example.com/buckets/secret_stuff)

- DELETE

  用在刪除資料。

  Examples:

  - DELETE [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - DELETE [http://www.example.com/customers/12345/orders](http://www.example.com/customers/12345/orders)
  - DELETE [http://www.example.com/bucket/sample](http://www.example.com/bucket/sample)

- PATCH

  用於修改資料，與 PUT 不同的是， PUT 需要傳入完整的資料，但 PATCH 只要傳入要修改的部分。

  Examples:

  - PATCH [http://www.example.com/customers/12345](http://www.example.com/customers/12345)
  - PATCH [http://www.example.com/customers/12345/orders/98765](http://www.example.com/customers/12345/orders/98765)
  - PATCH [http://www.example.com/buckets/secret_stuff](http://www.example.com/buckets/secret_stuff)

#### Summary

HTTP Verb | CRUD | Entire Collection (e.g. /customers) | Specific Item (e.g. /customers/{id})
:--------:| :---: | ----------------------------------- | ------------------------------
POST | Create | 201 (Created), 'Location' header with link to /customers/{id} containing new ID. | 404 (Not Found), 409 (Conflict) if resource already exists..
GET | Read | 200 (OK), list of customers. Use pagination, sorting and filtering to navigate big lists. | 200 (OK), single customer. 404 (Not Found), if ID not found or invalid.
PUT | Update/Replace | 404 (Not Found), unless you want to update/replace every resource in the entire collection. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.
DELETE | Delete | 404 (Not Found), unless you want to delete the whole collection—not often desirable. | 200 (OK). 404 (Not Found), if ID not found or invalid.
PATCH | Update/Modify | 404 (Not Found), unless you want to modify the collection itself. | 200 (OK) or 204 (No Content). 404 (Not Found), if ID not found or invalid.

### Test

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type category struct {
    ID     uint64 `json:"id,omitempty"`
    Name   string `json:"name,omitempty"`
    Parent uint64 `json:"parent,omitempty"`
}

var categories = make(map[uint64]*category)

func list(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    lst := make([]*category, 0, len(categories))

    for _, v := range categories {
        lst = append(lst, v)
    }

    dataBytes, err := json.Marshal(lst)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Write(dataBytes)
}

func find(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    category, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }

    dataBytes, err := json.Marshal(category)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    w.Write(dataBytes)

}

func add(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    category := new(category)

    dataBytes, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    err = json.Unmarshal(dataBytes, category)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    id := uint64(len(categories) + 1)
    category.ID = id

    categories[id] = category

    w.Header().Add("Location", fmt.Sprintf("/categories/%d", id))
}

func update(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    _, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }

    category := new(category)

    dataBytes, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    err = json.Unmarshal(dataBytes, category)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    category.ID = id
    categories[id] = category

    w.WriteHeader(204)
}

func del(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")

    vars := mux.Vars(r)

    id, err := strconv.ParseUint(vars["id"], 10, 64)
    if err != nil {
        w.WriteHeader(400)
        return
    }

    _, ok := categories[id]
    if !ok {
        w.WriteHeader(404)
        return
    }
    delete(categories, id)
    w.WriteHeader(204)

}

func main() {

    categories[1] = &category{
        ID:     1,
        Name:   "3C",
        Parent: 0,
    }

    r := mux.NewRouter()

    r.HandleFunc("/categories", list).Methods("GET")
    r.HandleFunc("/categories", add).Methods("POST")
    r.HandleFunc("/categories/{id:[0-9]+}", find).Methods("GET")
    r.HandleFunc("/categories/{id:[0-9]+}", update).Methods("PUT")
    r.HandleFunc("/categories/{id:[0-9]+}", del).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
}
```

- List all categories:  `curl --include http://localhost:8080/categories`
- Get one category:
  - success: `curl --include http://localhost:8080/categories/1`
  - failure: `curl --include http://localhost:8080/categories/100`
- Add a category: `curl --include --request POST --header "Content-type: application/json" --data '{"name": "PC", "parent": 1}' http://localhost:8080/categories`
- Update a category: `curl --include --request PUT --header "Content-type: application/json" --data '{"name": "NB-2", "parent": 1}' http://localhost:8080/categories/2`
- Delete a category: `curl --include --request DELETE http://localhost:8080/categories/2`

