package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config := new(Config)
	err := config.Load(ServiceName)
	if err != nil {
		t.Error("Expected loading of environment vars, got", err)
	}
}
