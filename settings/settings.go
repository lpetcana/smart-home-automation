package settings

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Settings struct{
	PrivateKeyPath string
	PublicKeyPath string
	JWTExpirationDelta int
}


var settings Settings = Settings{}
var env = "dev"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting dev environment due to lack of GO_ENV value")
		env = "dev"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile("settings/environments/dev.json")
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}