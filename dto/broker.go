package dto

type BrokerConfig struct {
	Url  string `json:"brokerUrl"`
	Port uint16 `json:"brokerPort"`
}
