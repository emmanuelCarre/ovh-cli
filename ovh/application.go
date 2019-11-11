package ovh

import (
	"strconv"
	"sync"
)

const applicationURI string = "/me/api/application/"

// Application data model
// See https://api.ovh.com/console/#/me/api/application/%7BapplicationId%7D#GET
type Application struct {
	Name           string `json:"name"`
	ApplicationKey string `json:"applicationKey"`
	ApplicationID  int    `json:"applicationId"`
	Description    string `json:"description"`
	Status         string `json:"status"`
}

// ToArrow return string slice from Application struct
func (a Application) ToArrow() []string {
	return []string{strconv.Itoa(a.ApplicationID), a.ApplicationKey, a.Name, a.Description, a.Status}
}

// FetchApplication return OVH resource instance
func FetchApplication(client *Client, applicationID int, resource *Application) error {
	return FetchResource(client, applicationURI, applicationID, resource)
}

func FetchApplications(client *Client) []Application {
	applicationsID, _ := FetchResourcesID(client, applicationURI)
	applications := make([]Application, len(applicationsID))
	var wg sync.WaitGroup
	for index, applicationID := range applicationsID {
		wg.Add(1)
		go func(index int, applicationID int) {
			defer wg.Done()
			client := GetOvhClient()
			var application Application
			FetchApplication(client, applicationID, &application)
			applications[index] = application
		}(index, applicationID)
	}
	wg.Wait()
	return applications
}

func DeleteApplications(client *Client, applicationsID []string) {
	DeleteResources(client, applicationURI, applicationsID)
}
