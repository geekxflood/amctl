// cmd/root.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var alertmanagerURL string
var ignoreTLS bool

var rootCmd = &cobra.Command{
	Use:   "alertmanager-cli",
	Short: "CLI tool to interact with Alertmanager",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&alertmanagerURL, "url", "u", "", "URL of the Alertmanager instance")
	rootCmd.PersistentFlags().BoolVarP(&ignoreTLS, "insecure", "i", false, "Ignore TLS for self-signed certificates")
	rootCmd.MarkPersistentFlagRequired("url")
}
