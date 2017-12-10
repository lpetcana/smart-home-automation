package controllers

import (
	"net/http"
	"fmt"
	"os"
	"smart-home-automation/settings"
)

func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Test Settings: check in console"))

	fmt.Println("Checking settings...")

	_, err := os.Stat("settings/environments/dev.json")
	if err == nil {
		fmt.Println("settings file exists!")
	} else {
		fmt.Println("Settings file does not exists!")
	}

	_, err2 := os.Stat("settings/keys/private_key.pem")
	if err2 == nil {
		fmt.Println("private key file exists!")
	} else {
		fmt.Println("private key  does not exists!")
	}



	fmt.Println("Environmnent from Settings:" + settings.GetEnvironment())

}

func Index(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))

}
