package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(insecureCmd)
}

// insecureCmd represents insecure TLS flag
var insecureCmd = &cobra.Command{}
