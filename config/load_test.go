package config_test

import (
	"main/config"
	"testing"
)

func TestLoadBroker(t *testing.T) {
	brokerConfig, err := config.LoadBrokerConfig()
	if err != nil {
		t.Fatalf("%+v FAILED: Config Required", t.Name())
	}

	t.Logf("SUCCEED: %v\n", brokerConfig)
}

func TestLoadClient(t *testing.T) {
	clientConfig, err := config.LoadClientConfig()
	if err != nil {
		t.Fatalf("%+v FAILED: Config Required", t.Name())
	}

	t.Logf("SUCCEED: %v\n", clientConfig)
}
