package config

import (
	"encoding/json"
	"errors"
	"os"
)

// Config is the data structure for the mapping of the "config.json" properties
type Config struct {
	Organisation string `json:"organisation"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	AccessToken  string `json:"access_token"`
}

// LoadConfiguration loads the config file that is specified as a parameter and decodes
// it into a Config struct object
func LoadConfiguration(configPath string) (Config, error) {
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		return Config{}, errors.New("Configuration loading failed: Error when opening config file: " + configPath + "\r\n" +
			"Please check you have specified the correct path to the json file. Example: folder/config/config.json")
	}

	jsonParser := json.NewDecoder(configFile)

	var config Config
	if jsonParser.Decode(&config) != nil {
		return Config{}, errors.New("Configuration loading failed: Error decoding config.json file, please check it is valid json")
	}
	return config, nil
}
