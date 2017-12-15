package services

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type test_struct struct {
	Test string
}

func GetWeather(updateParam string) []byte {
	updateNow, err := strconv.ParseBool(updateParam)
	GetWeatherFromAPI()
	fmt.Println("Update now value is:", updateNow)
	if err != nil {
		return []byte("update is false/wrong")
	} else {
		return []byte("update is true")
	}

}

func GetWeatherFromAPI() {
	//baseUrl := settings.Get().WeatherUrl
	//secret :=settings.Get().WeatherSecret

	weatherUrl := "https://api.darksky.net/forecast/e8c395ab6c6b8d46717b0afd44a74d1c/44.446373,26.059751?units=si"
	// Build the request
	req, err := http.NewRequest("GET", weatherUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}
