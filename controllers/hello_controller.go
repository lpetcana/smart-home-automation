package controllers

import (
	"net/http"
	"os"
	"log"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Test Settings: check in console"))

	log.Println("Checking settings...")

	_, err := os.Stat("settings/environments/dev.json")
	if err == nil {
		log.Println("settings file exists!")
	} else {
		log.Println("Settings file does not exists!")
	}

	_, err2 := os.Stat("settings/keys/private_key.pem")
	if err2 == nil {
		log.Println("private key file exists!")
	} else {
		log.Println("private key  does not exists!")
	}






}

func Index(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))

}
