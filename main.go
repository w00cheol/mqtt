package main

import (
	"fmt"
	"os"
	"time"
)

const (
	qos byte = 0
)

func main() {
	SetLogger()
	logger := NewLogger()

	config, err := LoadConfig(os.Getenv("MQTT_CONFIG_PATH"))
	if err != nil {
		logger.Error(err)
		return
	}

	client := NewClient(NewClientOption(config))

	broker := NewBroker(
		config.BrokerIP,
		config.BrokerPort,
		qos,
		client,
	)
	broker.Connect()
	defer broker.Disconnect(300)

	broker.Topics = append(broker.Topics, "woo/test")
	topicIdx := 0
	if err := broker.Publish(topicIdx, 0, false, fmt.Sprint(time.Now())); err != nil {
		logger.Error(err)
		return
	}

	in := make(chan *Datum)
	if err := broker.Subscribe(0, 0, in); err != nil {
		logger.Error(err)
		return
	}

	for payload := range in {
		logger.Debug(fmt.Sprintf("%+v", payload))
	}
}
