package settings

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path"
)

type Settings struct {
	Name				string
	PrivateKeyPath    	string 	`json:"privateKeyPath"`
	PublicKeyPath      	string 	`json:"publicKeyPath"`
	JWTExpirationDelta 	int 	`json:"JWTExpirationDelta"`
	WeatherUrl 			string 	`json:"weatherUrl"`
	WeatherSecret 		string 	`json:"weatherSecret"`
	Db 					DbConfig `json:"database"`
}


type DbConfig struct {
	Host string  `json:"host"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	DBName string  `json:"dbName"`
}



func LoadSettings(env string) *Settings {
	envFilePath := path.Join(".", "settings", "environments", env+".json")
	envFileContents, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		panic(fmt.Sprintf("Could not load %s config file: %s", envFilePath, err.Error()))
	}

	settings := &Settings{Name: env}

	err = json.Unmarshal(envFileContents, &settings)
	if err != nil {
		panic(fmt.Sprintf("Found config but could not load JSON: %s", err.Error()))
	}

	return settings
}




