package client

import (
	"fmt"
	"log"
	"main/config"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe() error {
	topic := "woo/test"

	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	// mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	subscriberOptions := mqtt.NewClientOptions()

	brokerConfig, err := config.LoadBrokerConfig()
	if err != nil {
		log.Fatal(err)
	}

	subcriberConfig, err := config.LoadClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	subscriberOptions.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerConfig.Url, brokerConfig.Port))
	subscriberOptions.SetClientID(subcriberConfig.SubscriberID)
	subscriberOptions.SetUsername(subcriberConfig.Username)
	subscriberOptions.SetPassword(subcriberConfig.Password)

	subcriber := mqtt.NewClient(subscriberOptions)
	if token := subcriber.Connect(); token.Wait() && token.Error() != nil {
		log.Print()
		return token.Error()
	}

	if token := subcriber.Subscribe(topic, 0, func(c mqtt.Client, msg mqtt.Message) {
		fmt.Printf("c: %+v\n", c)
		fmt.Printf("msg: %+v\n", string(msg.Payload()))
	}); token.Wait() && token.Error() != nil {
		log.Print(token.Error())
		return token.Error()
	}

	return nil
}
