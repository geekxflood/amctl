package main

import (
	"fmt"
	"os"

	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/spf13/cobra"
)

func main() {
	var alertmanagerURL string
	var list bool
	var foundAlert bool

	var rootCmd = &cobra.Command{
		Use:   "amctl",
		Short: "Alertmanager CLI",
		Run: func(cmd *cobra.Command, args []string) {
			if alertmanagerURL == "" {
				fmt.Println("Please provide an Alertmanager URL")
				return
			}
			apiClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
				Host:     alertmanagerURL,
				BasePath: client.DefaultBasePath,
				Schemes:  []string{"http"},
			})

			if list {
				fmt.Println("Listing alerts")
				params := alert.NewGetAlertsParams()
				alerts, err := apiClient.Alert.GetAlerts(params)
				if err != nil {
					fmt.Println("Error fetching alerts: ", err)
					return
				}

				for _, alert := range alerts.GetPayload() {
					if *alert.Status.State == "active" {
						fmt.Printf("Alert Name: %s\n", alert.Labels["alertname"])
						fmt.Printf("Starts At: %s\n", alert.StartsAt.String())
						fmt.Printf("Ends At: %s\n", alert.EndsAt.String())
						fmt.Printf("Labels:\n")

						for labelKey, labelValue := range alert.Labels {
							fmt.Printf("  %s: %s\n", labelKey, labelValue)
						}

						fmt.Printf("Annotations:\n")

						for annotationKey, annotationValue := range alert.Annotations {
							fmt.Printf("  %s: %s\n", annotationKey, annotationValue)
						}

						fmt.Println("----------------------------------------------------")
						foundAlert = true
					}
				}

				if !foundAlert {
					fmt.Println("No alert found")
				}
			} else {
				fmt.Println("Fetching Alertmanager status:")
				status, err := apiClient.General.GetStatus(general.NewGetStatusParams())
				if err != nil {
					fmt.Println("Error fetching status: ", err)
					return
				}
				fmt.Printf("Status: %s\n", *status.GetPayload().VersionInfo.Version)
				fmt.Printf("Uptime: %s\n", *status.GetPayload().Uptime)
				fmt.Println("Cluster status:")
				fmt.Printf("  Peers: %d\n", len(status.Payload.Cluster.Peers))
				for _, peer := range status.Payload.Cluster.Peers {
					fmt.Printf("Peer: %s\n", *peer.Name)
				}
			}
		},
	}

	rootCmd.Flags().StringVarP(&alertmanagerURL, "alertmanager", "a", "", "Alertmanager URL")
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "List alerts")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
