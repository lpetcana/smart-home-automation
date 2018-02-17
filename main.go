package main

import (

	"time"

	"os"
	"fmt"
	. "smart-home-automation/settings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	var env string
	env = os.Getenv("APP_ENV")

	if env == "" {
		panic("Exit! No environment specified")
	} else {
		log.Println("starting app with environment: ", env)
	}

	settings:=LoadSettings(env)






	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", settings.Db.Username, settings.Db.Password, settings.Db.Host, settings.Db.DBName)
	_, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic("something bad happend",err)
	} else {
		log.Println("successfully connected to sql")
	}

	router := InitRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))

}
