package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"html"
	"encoding/json"
)

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("hello")
	log.Println("Index called")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars:=mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Implemented"))
})

