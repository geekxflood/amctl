/*
Copyright © 2023 GeekxFlood

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/spf13/cobra"
)

var alertmanagerURL string
var list bool
var foundAlert bool
var insecure bool

var styleSeverityCritical = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#FF0000")).
	Foreground(lipgloss.Color("#FFFFFF"))

var rootCmd = &cobra.Command{
	Use:   "amctl",
	Short: "Alertmanager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if alertmanagerURL == "" {
			fmt.Println("Please provide an Alertmanager URL")
			return
		}

		// Create a new client for the Alertmanager API
		// The client will be used to make API calls
		apiClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
			Host:     alertmanagerURL,
			BasePath: client.DefaultBasePath,
			Schemes:  []string{"http"},
		})

		// Create a new parameter for the API call
		// The parameter will be used to filter the alerts
		statusParams := general.NewGetStatusParams()
		alertParams := alert.NewGetAlertsParams()
		if insecure {
			statusParams.HTTPClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
			alertParams.HTTPClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			}
		}

		// Check if the user wants to list the alerts
		// If not, we will just print the Alertmanager status
		if list {
			fmt.Println("Listing alerts")

			// Make the API call
			// The API call will return a list of alerts
			alerts, err := apiClient.Alert.GetAlerts(alertParams)
			if err != nil {
				fmt.Println("Error fetching alerts: ", err)
				return
			}

			// Loop through the list of alerts
			// Print the alert name, start and end time, labels and annotations
			for _, alert := range alerts.GetPayload() {
				if *alert.Status.State == "active" {
					fmt.Printf("Alert Name: %s\n", alert.Labels["alertname"])
					fmt.Printf("Starts At: %s\n", convertDate(alert.StartsAt.String()))
					fmt.Printf("Labels:\n")

					for labelKey, labelValue := range alert.Labels {
						if labelKey == "severity" {
							if labelValue == "critical" {
								fmt.Printf("  %s: %s\n", labelKey, styleSeverityCritical.Render(labelValue))
							}
						} else {
							fmt.Printf("  %s: %s\n", labelKey, labelValue)
						}
					}

					fmt.Printf("Annotations:\n")

					for annotationKey, annotationValue := range alert.Annotations {
						fmt.Printf("  %s: %s\n", annotationKey, annotationValue)
					}

					fmt.Println("----------------------------------------------------")
					foundAlert = true
				}
			}

			// If no alerts are found, print a message
			if !foundAlert {
				fmt.Println("No alert found")
			}

			// If the user does not want to list the alerts
			// We will just print the Alertmanager status
		} else {
			fmt.Println("Fetching Alertmanager status:")

			status, err := apiClient.General.GetStatus(statusParams)
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&alertmanagerURL, "alertmanager", "a", "", "Alertmanager URL")
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "List alerts")
	rootCmd.Flags().BoolVarP(&insecure, "insecure", "i", false, "Skip TLS verification")
}

func convertDate(date string) string {
	dateParse, err := time.Parse("2006-01-02T15:04:05.999999999Z", date)
	if err != nil {
		fmt.Println("Error converting date: ", err)
		dateParse = time.Now()
	}

	dateConv := dateParse.Format("Mon 2 Jan 2006 15:04:05")

	return dateConv
}
