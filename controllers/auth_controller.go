package controllers

import (
	"encoding/json"
	"smart-home-automation/models"
	"net/http"
	"smart-home-automation/services"

	"smart-home-automation/core/authentication"
)

func Login(w http.ResponseWriter, r *http.Request) {

	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func CheckTokenInfo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	token := r.URL.Query()["token"]
	if token ==nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authentication.CheckToken(token[0])


	w.WriteHeader(http.StatusOK)

}
