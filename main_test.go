package main

import (
	"os"
	"testing"

	"github.com/ChrisJBurns/label-hook/config"
)

func TestGetConfigurationPathWithTooLittleCommandArguments(t *testing.T) {
	os.Args = []string{"first"}

	_, err := GetConfigurationPath()
	if err == nil {
		t.Errorf("Test failed, expected to fail when no arguments are provided to application")
	}
}

func TestGetConfigurationPathWithTooManyCommandArguments(t *testing.T) {
	os.Args = []string{"first", "second", "third"}

	_, err := GetConfigurationPath()
	if err == nil {
		t.Errorf("Test failed, expected to fail when too many arguments are provided to application")
	}
}

func TestGetConfigurationPathWithCorrectNumberOfCommandArguments(t *testing.T) {
	os.Args = []string{"first", "second"}

	config, err := GetConfigurationPath()
	if err != nil {
		t.Errorf("Test failed, error returned from GetConfigurationPath with correct number of arguments")
	}

	if config == "" {
		t.Errorf("Test failed, incorrect config path returned")
	}
}

func TestLoadConfigurationFailure(t *testing.T) {
	_, err := config.LoadConfiguration("")
	if err == nil {
		t.Errorf("Test failed, expected an error when loading a invalid config file")
	}
}

func TestLoadConfigurationSuccess(t *testing.T) {
	_, err := config.LoadConfiguration("config/config.json")
	if err != nil {
		t.Errorf("Test failed, expected no error when loading a valid config file")
	}
}

func TestValidateConfigFailureEmptyConfig(t *testing.T) {
	cfg := config.Config{}
	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("Test failed, expected error when validating blank config struct")
	}
}

func TestValidateConfigFailureNoOrganisation(t *testing.T) {
	cfg := config.Config{Organisation: "",
		Host:        "host",
		Port:        "port",
		AccessToken: "access_token"}

	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("Test failed, expected error when validating blank organisation property")
	}
}

func TestValidateConfigFailureNoHost(t *testing.T) {
	cfg := config.Config{Organisation: "organisation",
		Host:        "",
		Port:        "port",
		AccessToken: "access_token"}

	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("Test failed, expected error when validating blank host property")
	}
}

func TestValidateConfigFailureNoPort(t *testing.T) {
	cfg := config.Config{Organisation: "organisation",
		Host:        "host",
		Port:        "",
		AccessToken: "access_token"}

	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("Test failed, expected error when validating blank port property")
	}
}

func TestValidateConfigFailureNoAccessToken(t *testing.T) {
	cfg := config.Config{Organisation: "organisation",
		Host:        "host",
		Port:        "port",
		AccessToken: ""}

	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("Test failed, expected error when validating blank access token property")
	}
}

func TestCreateClient(t *testing.T) {
	CreateClient()

	if client == nil {
		t.Errorf("Test failed, expected created github client")
	}
}
