package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Datum struct {
	value any
}

type MQTT struct {
	IP     string `json:"brokerIP"`
	Port   uint16 `json:"brokerPort"`
	Topics []string
	QoS    byte
	client mqtt.Client
}

type MQTTConfig struct {
	BrokerIP   string `json:"brokerIP"`
	BrokerPort uint16 `json:"brokerPort"`
	ClientID   string `json:"clientID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewBroker(ip string, port uint16, qos byte, client mqtt.Client) *MQTT {
	return &MQTT{
		IP:     ip,
		Port:   port,
		Topics: make([]string, 0),
		QoS:    qos,
		client: client,
	}
}

func NewClient(options *mqtt.ClientOptions) mqtt.Client {
	return mqtt.NewClient(options)
}

func NewClientOption(config *MQTTConfig) *mqtt.ClientOptions {
	return mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", config.BrokerIP, config.BrokerPort)).
		SetClientID(config.ClientID).
		SetUsername(config.Username).
		SetPassword(config.Password)
}

func (broker *MQTT) Connect() error {
	client := broker.client
	token := client.Connect()

	token.Wait()
	// could be err or nil
	return token.Error()
}

func (broker *MQTT) Disconnect(quiesce uint) {
	client := broker.client
	client.Disconnect(quiesce)
}

func (broker *MQTT) Publish(idx int, qos byte, retained bool, msg string) error {
	client := broker.client
	if token := client.Publish(broker.Topics[idx], qos, retained, msg); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (broker *MQTT) Subscribe(idx int, qos byte, out chan *Datum) error {
	client := broker.client
	token := client.Subscribe(broker.Topics[idx], qos, func(c mqtt.Client, m mqtt.Message) {
		out <- &Datum{value: string(m.Payload())}
	})

	token.Wait()
	return token.Error()
}

func (broker *MQTT) Unsubscribe(idx int) error {
	client := broker.client
	token := client.Unsubscribe(broker.Topics[idx])

	token.Wait()
	return token.Error()
}

func (broker *MQTT) UnsubscribeAll() error {
	for i := range broker.Topics {
		if err := broker.Unsubscribe(i); err != nil {
			return err
		}
	}

	return nil
}
