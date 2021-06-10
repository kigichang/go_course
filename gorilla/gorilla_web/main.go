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
