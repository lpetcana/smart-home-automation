package main

import (
	"log"

	"net/http"
	"time"
	"fmt"
)

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

var mySigningKey = []byte("secret")

func main() {
	router := InitRoutes()
	fmt.Println("there we go!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
