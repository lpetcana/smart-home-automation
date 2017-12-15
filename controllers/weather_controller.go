package controllers

import (
	"net/http"
	"fmt"
	"smart-home-automation/services"
)

func Weather(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("GET params were:", r.URL.Query());

	// if only one expected
	updateNow := r.URL.Query().Get("update-now")

	x:=services.GetWeather(updateNow)



	//w.Header().Set("Content-Type", "application/json")
	w.Write(x)
}