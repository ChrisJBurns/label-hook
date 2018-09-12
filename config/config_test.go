package config

import (
	"testing"
)

func TestLoadConfigurationFailure(t *testing.T) {
	_, err := LoadConfiguration("")
	if err == nil {
		t.Errorf("Test failed, expected an error when loading a invalid config file")
	}
}

func TestLoadConfigurationSuccess(t *testing.T) {
	_, err := LoadConfiguration("config.json")
	if err != nil {
		t.Errorf("Test failed, expected no error when loading a valid config file")
	}
}
