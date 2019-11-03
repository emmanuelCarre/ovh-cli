package ovh

import (
	"github.com/ovh/go-ovh/ovh"
	log "github.com/sirupsen/logrus"
)

// Client interface
type Client interface {
	Get(string, interface{}) error
	Delete(string, interface{}) error
}

// GetOvhClient return configured OVH client
func GetOvhClient() *Client {
	//func GetOvhClient() *ovh.Client {
	config := GetCurrentAPIConfig()
	client, err := ovh.NewClient(
		config.Endpoint,
		config.ApplicationKey,
		config.ApplicationSecret,
		config.ConsumerKey,
	)
	if err != nil {
		log.Fatalf("Unable to create ovh client: %v", err)
	}
	test := Client(client)
	return &test
}
