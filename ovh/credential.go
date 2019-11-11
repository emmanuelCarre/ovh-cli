package ovh

import (
	"bytes"
	"strconv"
	"sync"
	"time"

	"github.com/AdFabConnect/ovh-cli/ovh"
	log "github.com/sirupsen/logrus"
)

type Rule struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type PartialCredentials struct {
	ExpirationDate time.Time `json:"expiration"`
	ApplicationID  int       `json:"applicationId"`
	CredentialID   int       `json:"credentialId"`
	Rules          []Rule    `json:"rules"`
}

type PartialCredentialApplication struct {
	Name string `json:"name"`
}

const credentialURI string = "/me/api/credential/"

// ToArrow return string slice from Application struct
func (c PartialCredentials) ToArrow() []string {
	var credentialRight bytes.Buffer
	for index, item := range c.Rules {
		credentialRight.WriteString(item.Method)
		credentialRight.WriteString(" ")
		credentialRight.WriteString(item.Path)
		if index < len(c.Rules)-1 {
			credentialRight.WriteString(", ")
		}
	}
	client := ovh.GetOvhClient()
	var credentialApplication PartialCredentialApplication
	err := client.Get(credentialURI+strconv.Itoa(c.CredentialID)+"/application", &credentialApplication)
	if err != nil {
		log.Errorln("Unable to get credential details:", err)
	}
	return []string{strconv.Itoa(c.CredentialID), c.ExpirationDate.Format(time.ANSIC), strconv.Itoa(c.ApplicationID), credentialApplication.Name, credentialRight.String()}
}

// FetchCredential return OVH resource instance
func FetchCredential(client *Client, applicationID int, resource *PartialCredentials) error {
	return FetchResource(client, credentialURI, applicationID, resource)
}

func FetchCredentials(client *Client) []PartialCredentials {
	credentialsID, _ := FetchResourcesID(client, credentialURI)
	credentials := make([]PartialCredentials, len(credentialsID))
	var wg sync.WaitGroup
	for index, credentialID := range credentialsID {
		wg.Add(1)
		go func(index int, credentialID int) {
			defer wg.Done()
			client := GetOvhClient()
			var credential PartialCredentials
			FetchCredential(client, credentialID, &credential)
			credentials[index] = credential
		}(index, credentialID)
	}
	wg.Wait()
	return credentials
}

func DeleteCredentials(client *Client, credentialsID []string) {
	DeleteResources(client, credentialURI, credentialsID)
}
