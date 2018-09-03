package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config is the data structure for the mapping of the "config.json" properties
type Config struct {
	Organisation string `json:"organisation"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	AccessToken  string `json:"access_token"`
}

// LoadConfiguration loads all configuration that exists in
// config.json and assigns it to a Config struct object
func LoadConfiguration() Config {
	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		log.Println("error when opening config.json")
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)

	var config Config
	jsonParser.Decode(&config)
	return config
}
