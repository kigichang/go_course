package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// MyData ...
type MyData struct {
	Title string
	Nav   string
	Data  interface{}
}

func generateHTML(w http.ResponseWriter, data interface{}, files ...string) {
	var tmp []string
	for _, f := range files {
		tmp = append(tmp, fmt.Sprintf("templates/%s.html", f))
	}

	tmpl := template.Must(template.ParseFiles(tmp...))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func test(w http.ResponseWriter, r *http.Request) {
	data := &MyData{
		Title: "測試",
		Nav:   "test",
	}

	data.Data = struct {
		TestString   string
		SimpleString string
		TestStruct   struct{ A, B string }
		TestArray    []string
		TestMap      map[string]string
		Num1, Num2   int
		EmptyArray   []string
		ZeroInt      int
	}{
		`O'Reilly: How are <i>you</i>?`,
		"中文測試",
		struct{ A, B string }{"foo", "boo"},
		[]string{"Hello", "World", "Test"},
		map[string]string{"A": "B", "abc": "DEF"},
		10,
		101,
		[]string{},
		0,
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/nav.html", "templates/test.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

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

func toGo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://go.dev/")
	w.WriteHeader(http.StatusFound)
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", test)
	mux.HandleFunc("/set_cookie", setCookie)
	mux.HandleFunc("/get_cookie", getCookie)
	mux.HandleFunc("/go", toGo)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
