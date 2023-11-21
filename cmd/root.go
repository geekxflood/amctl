// cmd/root.go

package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/spf13/cobra"
	"github.com/geekxflood/alertmanager-cli/global"
)

var alertmanagerURL string
var list bool
var foundAlert bool
var insecure bool
var silenceParams string
var endTime string

var styleSeverityCritical = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#FF0000")).
	Foreground(lipgloss.Color("#FFFFFF"))

var styleHeader = lipgloss.NewStyle().
	PaddingTop(1).
	Underline(true)

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
		postSilenceParams := silence.NewPostSilencesParams()
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
			postSilenceParams.HTTPClient = &http.Client{
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
					fmt.Println(styleHeader.Render("Alert Name:", alert.Labels["alertname"]))
					fmt.Printf("Starts At: %s\n", ConvertDate(alert.StartsAt.String()))
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
			// If the user does not want to list the alerts
			// We will just print the Alertmanager status
			if silenceParams != "" {
				fmt.Println("Silencing alerts")

				var endsAt time.Time
				startsAt := time.Now()
				if endTime == "" {
					fmt.Println("No end time provided, silencing for 2 hours")
					endsAt = startsAt.Add(2 * time.Hour)
				} else {
					endsAt, _ = time.Parse("2006-01-02T15:04:05", endTime)
				}
				if silenceParams == "" {
					panic("No silence parameters provided, quitting")
				}

				// Parse silenceParams into key-value pairs
				paramsSets := strings.Split(silenceParams, ";")
				silenceMatcher := make([]*models.Matcher, 0)
				for _, paramSet := range paramsSets {
					pairs := strings.Split(paramSet, ",")
					for _, pair := range pairs {
						kv := strings.Split(pair, "=")
						if len(kv) != 2 {
							panic("Invalid silence parameters format, expecting key1=value1,key2=value2,...")
						}
						silenceMatcher = append(silenceMatcher, &models.Matcher{
							Name:    swag.String(kv[0]),
							Value:   swag.String(kv[1]),
							IsRegex: swag.Bool(false),
						})
					}
				}

				// Convert time.Time to strfmt.DateTime
				startsAtStrFmt, err := strfmt.ParseDateTime(startsAt.Format(time.RFC3339))
				if err != nil {
					panic("Failed to convert startsAt to strfmt.DateTime")
				}
				endsAtStrFmt, err := strfmt.ParseDateTime(endsAt.Format(time.RFC3339))
				if err != nil {
					panic("Failed to convert endsAt to strfmt.DateTime")
				}

				// Create a new silence
				newSilence := &models.PostableSilence{}
				newSilence.Comment = swag.String("Silenced by amctl")
				newSilence.Matchers = silenceMatcher
				newSilence.StartsAt = &startsAtStrFmt
				newSilence.EndsAt = &endsAtStrFmt

				fmt.Println("Silence parameters:")
				fmt.Printf("  Starts At: %s\n", startsAtStrFmt)
				fmt.Printf("  Ends At: %s\n", endsAtStrFmt)
				fmt.Printf("  Matchers: %v\n", silenceMatcher)

				resp, err := apiClient.Silence.PostSilences(postSilenceParams.WithSilence(newSilence))
				if err != nil {
					if apiErr, ok := err.(*runtime.APIError); ok {
						if httpResponse, ok := apiErr.Response.(*http.Response); ok {
							responseBody, readErr := io.ReadAll(httpResponse.Body)
							if readErr != nil {
								fmt.Printf("Error reading response body: %s\n", readErr)
							} else {
								fmt.Printf("Error silencing alerts: %s\n", string(responseBody))
							}
						} else {
							fmt.Printf("Error silencing alerts: %v\n", apiErr.Response)
						}
					} else {
						fmt.Println("Error silencing alerts: ", err)
					}
					return
				}

				fmt.Printf("Silence ID: %s\n", *resp.GetPayload())

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
	rootCmd.Flags().BoolVarP(&insecure, "insecure", "i", false, "Skip TLS verification")
	rootCmd.Flags().StringVarP(&endTime, "end-time", "e", "", "End Time for Silence")
}
