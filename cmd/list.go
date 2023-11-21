// cmd/list.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List active alerts",
	Long: `List all active alerts from Alertmanager. 
           This command retrieves and displays currently active alerts, 
           including details such as alert names, start times, labels, and annotations.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation to list alerts
		fmt.Println("Listing active alerts...")
		// Here, you will call a function to interact with the Alertmanager API
		// and list the active alerts.
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// Initialize any flags or settings specific to the list command
}
