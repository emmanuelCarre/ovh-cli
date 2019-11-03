package model

import "strconv"

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
func (a *Application) ToArrow() []string {
	return []string{strconv.Itoa(a.ApplicationID), a.ApplicationKey, a.Name, a.Description, a.Status}
}
