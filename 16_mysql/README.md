# 16 MySQL

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [16 MySQL](#16-mysql)
  - [0. 前言](#0-前言)
    - [Test table schema (MySQL):](#test-table-schema-mysql)
    - [Sample Code](#sample-code)
  - [1. Initial Connection Pool](#1-initial-connection-pool)
  - [2. Insert](#2-insert)
  - [3. Select](#3-select)
  - [4. Fast Query/Insert](#4-fast-queryinsert)
  - [5. Transaction](#5-transaction)
  - [6. Connect, Prepared Statement, Rows, Cursor 關係](#6-connect-prepared-statement-rows-cursor-關係)

<!-- /code_chunk_output -->

## 0. 前言

與 Java JDBC 類似，Go 有定義一套 interface，所有要連 DB 的 driver，都需要實作這些 interface (["database/sql/driver"](https://golang.org/pkg/database/sql/driver/))。以下我是用 [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

### Test table schema (MySQL):

@import "ex16/init.sql" {class="line-numbers"}

### Sample Code

@import "ex16/main.go" {class="line-numbers"}

## 1. Initial Connection Pool

1. import package.

    ```go { .line-numbers }
    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
    )
    ```

    1. `"database/sql"` 是 go 定義 sql interface 的 package
    1. `_ "github.com/go-sql-driver/mysql"` mysql driver package

1. 定義資料的 struct，類似要做 ORM 的動作，當然也可以不要這個定義，都用變數來存資料。

    ```go { .line-numbers }
    // Member ...
    type Member struct {
        ...
    }
    ```

1. 連線資料庫，並取得 Connection Pool 物件。

    ```go { .line-numbers }
    // InitDB ...
    func InitDB() (*sql.DB, error) {
        db, err := sql.Open("mysql", "abc:1234test@tcp(localhost)/mytest?charset=utf8mb4,utf8&parseTime=true")
        if err != nil {
            return nil, err
        }
        return db, nil
    }
    ```

    與 JDBC 連線類似，指定 driver 的種類，並傳入一組 url 的設定, 格式是：`[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]`。詳細的說明，請見：[DSN (Data Source Name)](https://github.com/go-sql-driver/mysql#dsn-data-source-name)。我在連線後，可以多做了 Ping 的動作，如下：

    ```go { .line-numbers }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    ```

    記得取的 db 連線後，立即下 `defer db.Close()`，確保主程式在結束後，會關閉連線。如下：

    ```go { .line-numbers }
    db, err := InitDB()
	if err != nil {
		log.Fatal("initial db:", err)
	}
	defer db.Close()
    ```

## 2. Insert

```go { .line-numbers }
birthday := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
register := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
login := time.Now()

m1 := Member{
    Name:     "小明",
    Info:     "小明",
    Birthday: birthday,
    Register: register,
    Login:    login,
    VIP:      "A",
}
log.Println(m1)

ins, err := db.Prepare("insert into member(name, info, birthday, register, login, vip) values(?, ?, ?, ?, ?, ?)")
if err != nil {
    log.Println("prepare insert:", err)
    return
}
defer ins.Close()

result, err := ins.Exec(m1.Name, m1.Info, m1.Birthday, m1.Register, m1.Login, m1.VIP)
if err != nil {
    log.Println("insert:", err)
    return
}
defer ins.Close()

id, err := result.LastInsertId()
if err != nil {
    log.Println("last id:", err)
    return
}
```

說明：

1. 利用 DB.Pepare 建立一個 PreparedStatement 連線，記得下 `defer ins.Close()`

    ```go { .line-numbers }
    ins, err := db.Prepare("insert into member(name, info, birthday, register, login, vip) values(?, ?, ?, ?, ?, ?)")
    if err != nil {
        log.Println("prepare insert:", err)
        return
    }
    defer ins.Close()
    ```

1. 使用 Stmt.Exec 執行指 SQL 指令。

    ```go { .line-numbers }
    result, err := ins.Exec(m1.Name, m1.Info, m1.Birthday, m1.Register, m1.Login, m1.VIP)
    if err != nil {
        log.Println("insert:", err)
        return
    }
    defer ins.Close()

    id, err := result.LastInsertId()
    if err != nil {
        log.Println("last id:", err)
        return
    }
    ```

    Result 功能：

    1. LastInsertId(): 取得 **auto_increment** 的 id
    1. RowsAffected(): 取得異動的資料筆數，注意，[文件](https://golang.org/pkg/database/sql/#Result)上說並非所有的 driver 都會實作。

**Update/Delete** 與 Insert 類似。

## 3. Select

1. 下 SQL，與 Java 的 PreparedStatement 類似。

    ```go { .line-numbers }
    sel, err := db.Prepare("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?")
    if err != nil {
        log.Println("prepare select:", err)
        return
    }
    defer sel.Close()
    ```

    與上述 db 類似，取得連線後，記得下 `defer sel.Close()` 確保程式結束後，會關閉 statement 連線。

1. 使用 Stmt.Query 方式，取得 Rows

    ```go { .line-numbers }
    rows, err := sel.Query(id)
    if err != nil {
        log.Println("query:", err)
        return
    }

    defer rows.Close()
    ```

    與上述取得連線一樣，立即下 `defer rows.Close()` 確保程式結束後，會關閉 rows。(說明文件說，會[自動關閉](https://golang.org/pkg/database/sql/#Rows.Close)。這部分就看自己的習慣了。但 DB 與 Stmt 一定要記得關。)

1. 一定要先執行 **Next** 才能取資料。

    ```go { .line-numbers }
    if rows.Next() {
        m2 := Member{}
        err = rows.Scan(&m2.ID, &m2.Name, &m2.Info, &m2.Birthday, &m2.Register, &m2.Login, &m2.VIP, &m2.Created, &m2.Updated)
        if err != nil {
            log.Println("scan:", err)
            return
        }
        bytes, _ := json.Marshal(m2)
        fmt.Println(string(bytes))
    } else {
        log.Printf("data (%d) not found\n", id)
    }
    ```

1. 透過 Rows.Scan 取得資料。

    ```go { .line-numbers }
    err = rows.Scan(&m2.ID, &m2.Name, &m2.Info, &m2.Birthday, &m2.Register, &m2.Login, &m2.VIP, &m2.Created, &m2.Updated)
    if err != nil {
        log.Println("scan:", err)
        return
    }
    ```

## 4. Fast Query/Insert

上述的例子，是使用 `Stmt` 執行 SQL，並透過 `Rows` 來取得資料。如果不是需要重覆使用 `Stmt`時，可以直接用 `DB` 的 `DB.Query` 取得 `Rows`，或用 `DB.Exec` 執行 SQL。

如果只要取得單一筆資料，可以用 `DB.QueryRow` 取的 `Row` 物件，再直接用 `Row.Scan` 取得資料。

```go { .line-numbers }
func GetMember(db *sql.DB, id int64) (*Member, error) {
    mem := &Member{}
    err := db.QueryRow("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?", id).Scan(&mem.ID, &mem.Name, &mem.Info, &mem.Birthday, &mem.Register, &mem.Login, &mem.VIP, &mem.Created, &mem.Updated)

    if err != nil {
        return nil, err
    }
    return mem, nil
}
```

可利用 `errors.Is(err, sql.ErrNoRows)` 來判斷資料是否存在。

## 5. Transaction

要進行 Transction，可以先用 `DB.Begin` 取得 transaction 物件，如：`tx, err := db.Begin()`。之後的資料操作都是透過此物件進行，步驟與方法，與上面雷同。

1. 當有錯誤時，請用 `Rollback()`，回復先前修改的資料，如：`tx.Rollback()`。
1. 最後成功時，請用 `Commit()`，寫入資料庫，如：`tx.Commit()`。

## 6. Connect, Prepared Statement, Rows, Cursor 關係

![DB](db.png)
