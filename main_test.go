package main

import (
	"testing"

	"github.com/ChrisJBurns/label-hook/config"
)

func TestLoadConfig(t *testing.T) {
	cfg = config.LoadConfiguration()

	if cfg == (config.Config{}) {
		t.Error("Loading configuration failed: returned null")
	}
	if cfg.Port == "" {
		t.Error("Configured port is empty")
	}
	if cfg.Host == "" {
		t.Error("Configured host is empty")
	}
}
