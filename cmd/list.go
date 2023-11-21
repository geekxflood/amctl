// cmd/list.go

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all current alerts from Alertmanager",
	Long: `List all current alerts from Alertmanager.
This command retrieves and displays the list of all current alerts
being managed by Alertmanager.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
