/*
*
此範例程式，需要 MySQL。目前利過 go generate 與 Docker，在本機端啟動 MySQL。
使用方式如下：

1. 先安裝 docker
2. 執行範例程式前，先執行 `go generate`，會開始建立程式需要的 MySQL image 並啟動。
3. 執行 `docker logs -f go_course_db` 確認 MySQL 已經啟動
4. 執行範例程式 `go run .`
*
*/
package main

//go:generate docker rm -f go_course_db
//go:generate docker build -t go_course_ex16/db:latest .
//go:generate docker run -d --name=go_course_db -p 3306:3306 go_course_ex16/db:latest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Member ...
type Member struct {
	ID       int
	Name     string
	Info     string
	Birthday time.Time // MySQL min Date: 1000-01-01
	Register time.Time // MySQL min DateTime: 1000-01-01 00:00:00
	Login    time.Time // MySQL min DateTime: 1000-01-01 00:00:00
	VIP      string
	Created  time.Time
	Updated  time.Time
}

func (m *Member) String() string {
	memberBytes, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err.Error()
	}

	return string(memberBytes)
}

// InitDB ...
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "abc:1234test@tcp(localhost)/mytest?charset=utf8mb4,utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetMember ...
func GetMember(db *sql.DB, id int64) (*Member, error) {
	mem := &Member{}
	err := db.QueryRow("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?", id).Scan(&mem.ID, &mem.Name, &mem.Info, &mem.Birthday, &mem.Register, &mem.Login, &mem.VIP, &mem.Created, &mem.Updated)

	if err != nil {
		return nil, err
	}
	return mem, nil
}

func main() {

	db, err := InitDB()
	if err != nil {
		log.Fatal("initial db:", err)
	}
	defer db.Close()

	birthday := time.Now()
	register := time.Now()
	login := time.Now()

	m1 := &Member{
		Name:     "小明",
		Info:     "小明",
		Birthday: birthday,
		Register: register,
		Login:    login,
		VIP:      "A",
	}
	log.Println("member:", m1)

	/* Insert start */
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

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("last id:", err)
		return
	}
	/* Insert end */

	/* Update start */
	upt, err := db.Prepare("update member set name = ? where id = ?")
	if err != nil {
		log.Println(err)
		return
	}
	defer upt.Close()

	result, err = upt.Exec("小小明", id)
	if err != nil {
		log.Println(err)
		return
	}
	rowAffected, _ := result.RowsAffected()
	log.Println("update record", rowAffected)
	/* Update start */

	/* Select start */
	sel, err := db.Prepare("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?")
	if err != nil {
		log.Println("prepare select:", err)
		return
	}
	defer sel.Close()

	rows, err := sel.Query(id)
	if err != nil {
		log.Println("query:", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		m2 := &Member{}
		err = rows.Scan(&m2.ID, &m2.Name, &m2.Info, &m2.Birthday, &m2.Register, &m2.Login, &m2.VIP, &m2.Created, &m2.Updated)
		if err != nil {
			log.Println("scan:", err)
			return
		}
		fmt.Println("get member:", m2)
	} else {
		log.Printf("data (%d) not found\n", id)
	}

	/* Select end */

	other, err := GetMember(db, 100)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("id 100 not found")
		} else {
			log.Println("other error:", err)
		}
	} else {
		log.Println("other member:", other)
	}

	/* Transaction start */
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	result, err = tx.Exec(
		"insert into member(name, info, birthday, register, login, vip) values(?, ?, ?, ?, ?, ?)",
		"小華", "小華", birthday, register, login, "B")

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		log.Println("last id:", err)
		tx.Rollback()
		return
	}

	result, err = tx.Exec("update member set name= ? where id = ?", "小小華", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}
	rowAffected, _ = result.RowsAffected()
	log.Println("update record", rowAffected)
	tx.Commit()

	other, err = GetMember(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("id %d not found\n", id)
		} else {
			log.Println("other error:", err)
		}
	} else {
		log.Printf("member(%d): %s\n", id, other)
	}
	/* Transaction start */

	log.Println("end")
}
