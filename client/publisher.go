package client

import (
	"fmt"
	"log"
	"main/config"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish() error {
	topic := "woo/test"

	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	options := mqtt.NewClientOptions()

	brokerConfig, err := config.LoadBrokerConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientConfig, err := config.LoadClientConfig()
	if err != nil {
		log.Fatal(err)
	}

	options.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerConfig.Url, brokerConfig.Port))
	options.SetClientID(clientConfig.PublisherID)
	options.SetUsername(clientConfig.Username)
	options.SetPassword(clientConfig.Password)

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Print()
		return token.Error()
	}

	for i := 0; i < 10; i++ {
		if token := client.Publish(topic, 0, false, fmt.Sprint(time.Now())); token.Wait() && token.Error() != nil {
			log.Print(token.Error())
			return token.Error()
		}
		time.Sleep(1 * time.Second)
	}

	// if token := client.Unsubscribe(); token.Wait() && token.Error() != nil {
	// 	log.Print("??")
	// 	log.Print(token.Error())
	// 	return token.Error()
	// }

	client.Disconnect(300)

	return nil
}
