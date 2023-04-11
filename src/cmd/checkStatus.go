package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var alertmanagerURL string

func init() {
	rootCmd.AddCommand(checkStatusCmd)
	checkStatusCmd.Flags().StringVarP(
		&alertmanagerURL,
		"alertmanagerURL",
		"u",
		"http://localhost:9093",
		"Alertmanager URL",
	)

}

var checkStatusCmd = &cobra.Command{
	Use:   "chs",
	Short: "Check the status of the Alertmanager API",
	Long:  `Check the status of the Alertmanager API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Alertmanager URL: %s\n", alertmanagerURL)
		//
	},
}
