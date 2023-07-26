package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(path string) (*MQTTConfig, error) {
	readBrokerData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	brokerConfig := new(MQTTConfig)
	if err := json.Unmarshal(readBrokerData, brokerConfig); err != nil {
		return nil, err
	}

	return brokerConfig, nil
}
