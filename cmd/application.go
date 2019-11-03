package cmd

import (
	"sort"
	"strconv"
	"sync"

	ovhClient "github.com/emmanuelCarre/ovh-cli/ovh"
	"github.com/emmanuelCarre/ovh-cli/ovh/model"
	"github.com/emmanuelCarre/ovh-cli/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const applicationURI string = "/me/api/application/"

func getHeaderToDisplay() []string {
	return []string{"Id", "Key", "Name", "Description", "Status"}
}

func fetchApplication(client *ovhClient.Client, applicationID int, application *model.Application) error {
	err := (*client).Get(applicationURI+strconv.Itoa(applicationID), &application)
	if err != nil {
		log.Errorln("Unable to get application details:", err)
	}
	return err
}

var listApplicationCmd = &cobra.Command{
	Use:   "list",
	Short: "list applications",
	Run: func(cmd *cobra.Command, args []string) {
		applicationsID := []int{}
		client := ovhClient.GetOvhClient()
		isQuiet, _ := cmd.Flags().GetBool("quiet")

		err := (*client).Get(applicationURI, &applicationsID)
		if err != nil {
			log.Errorln("Unable to list applications:", err)
			return
		}
		table := utils.GetTable()

		sort.Ints(applicationsID)
		if isQuiet {
			for _, applicationID := range applicationsID {
				table.Append([]string{strconv.Itoa(applicationID)})
			}
			table.SetHeaderLine(false)
		} else {
			applications := make([]model.Application, len(applicationsID))
			var wg sync.WaitGroup
			for index, applicationID := range applicationsID {
				wg.Add(1)
				go func(index int, applicationID int) {
					defer wg.Done()
					client := ovhClient.GetOvhClient()
					var application model.Application
					fetchApplication(client, applicationID, &application)
					applications[index] = application
				}(index, applicationID)
			}
			wg.Wait()
			table.SetHeader(getHeaderToDisplay())
			for _, application := range applications {
				table.Append(application.ToArrow())
			}
		}
		table.Render()
	},
}

var deleteApplicationCmd = &cobra.Command{
	Use:   "delete [application id]...",
	Short: "delete applications",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := ovhClient.GetOvhClient()
		for _, applicationID := range args {
			err := (*client).Delete(applicationURI+applicationID, nil)
			if err != nil {
				log.Errorln("Unable to delete application with id", applicationID)
			}
		}
	},
}

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manipulate applications",
}

func init() {
	rootCmd.AddCommand(applicationCmd)
	applicationCmd.AddCommand(listApplicationCmd)
	applicationCmd.AddCommand(deleteApplicationCmd)
	listApplicationCmd.Flags().BoolP("quiet", "q", false, "Only display application id")
}
