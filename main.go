package main

import (
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"html"
	"time"
	"encoding/json"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("hello")
	log.Println("Index called")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
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

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo Show:", todoId)
}