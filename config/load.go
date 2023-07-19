package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/dto"
	"os"
)

func LoadBrokerConfig() (*dto.BrokerConfig, error) {
	var readBrokerData []byte
	readBrokerData, err := ioutil.ReadFile(os.Getenv("BROKER_CONFIG_PATH"))
	if err != nil {
		log.Print(err)
		fmt.Println("broker")
		return nil, err
	}

	brokerConfig := &dto.BrokerConfig{}
	if err := json.Unmarshal(readBrokerData, brokerConfig); err != nil {
		log.Print(err)
		fmt.Println("broker")
		return nil, err
	}

	return brokerConfig, nil
}

func LoadClientConfig() (*dto.ClientConfig, error) {
	var readClientData []byte
	readClientData, err := ioutil.ReadFile(os.Getenv("CLIENT_CONFIG_PATH"))
	if err != nil {
		log.Print(err)
		fmt.Println("LoadClientConfig")
		return nil, err
	}

	clientConfig := &dto.ClientConfig{}
	if err := json.Unmarshal(readClientData, clientConfig); err != nil {
		log.Print(err)
		fmt.Println("LoadClientConfig")
		return nil, err
	}

	return clientConfig, nil
}
