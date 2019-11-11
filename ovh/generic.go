package ovh

import (
	"sort"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Resource OVH
type Resource interface {
	ToArrow() []string
}

type fn func(wg *sync.WaitGroup, index int, applicationID int)

// FetchResource returne resource from OVH API
func FetchResource(client *Client, resourceURI string, resourceID int, resource Resource) error {
	err := (*client).Get(resourceURI+strconv.Itoa(resourceID), resource)
	if err != nil {
		log.Errorln("Unable to get resource details:", err)
	}
	return err
}

func FetchResourcesID(client *Client, resourceURI string) ([]int, error) {
	ResourcesID := []int{}

	err := (*client).Get(resourceURI, &ResourcesID)
	if err != nil {
		log.Errorln("Unable to list resources:", err)
		return nil, err
	}
	sort.Ints(ResourcesID)
	return ResourcesID, err
}

func DeleteResources(client *Client, resourceURI string, resourcesID []string) {
	for _, resourceID := range resourcesID {
		err := (*client).Delete(resourceURI+resourceID, nil)
		if err != nil {
			log.Errorln("Unable to delete resource with id", resourceID)
		}
	}
}
