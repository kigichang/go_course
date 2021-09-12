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

	log.Println("service starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
