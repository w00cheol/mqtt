package dto

type ClientConfig struct {
	PublisherID  string `json: "publisherID"`
	SubscriberID string `json: "subscriberID"`
	Username     string `json: "username"`
	Password     string `json: "password"`
}
