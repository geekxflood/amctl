// cli/cli.go

package cli

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/prometheus/alertmanager/api/v2/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// TLS verify flag
var TLSverify bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.AlertmanagerAPI, error) {
	hostname := viper.GetString("hostname")
	viper.SetDefault("base_path", client.DefaultBasePath)
	basePath := viper.GetString("base_path")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, basePath, []string{scheme})

	// call DisableTLSVerification if TLSverify is set to true
	if TLSverify {
		r.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/json"] = runtime.JSONConsumer()

	r.Producers["application/json"] = runtime.JSONProducer()

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))
	rootCmd.PersistentFlags().String("base-path", client.DefaultBasePath, fmt.Sprintf("For example: %v", client.DefaultBasePath))
	viper.BindPFlag("base_path", rootCmd.PersistentFlags().Lookup("base-path"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")
	// configure TLS verification flag
	rootCmd.PersistentFlags().BoolVar(&TLSverify, "insecure", false, "disable TLS verification")

	// register security flags
	// add all operation groups
	operationGroupAlertCmd, err := makeOperationGroupAlertCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAlertCmd)

	operationGroupAlertgroupCmd, err := makeOperationGroupAlertgroupCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAlertgroupCmd)

	operationGroupGeneralCmd, err := makeOperationGroupGeneralCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupGeneralCmd)

	operationGroupReceiverCmd, err := makeOperationGroupReceiverCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupReceiverCmd)

	operationGroupSilenceCmd, err := makeOperationGroupSilenceCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupSilenceCmd)

	rootCmd.AddCommand()

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

func makeOperationGroupAlertCmd() (*cobra.Command, error) {
	operationGroupAlertCmd := &cobra.Command{
		Use: "alert",
	}

	operationGetAlertsCmd, err := makeOperationAlertGetAlertsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAlertCmd.AddCommand(operationGetAlertsCmd)

	operationPostAlertsCmd, err := makeOperationAlertPostAlertsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAlertCmd.AddCommand(operationPostAlertsCmd)

	return operationGroupAlertCmd, nil
}
func makeOperationGroupAlertgroupCmd() (*cobra.Command, error) {
	operationGroupAlertgroupCmd := &cobra.Command{
		Use: "alertgroup",
	}

	operationGetAlertGroupsCmd, err := makeOperationAlertgroupGetAlertGroupsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupAlertgroupCmd.AddCommand(operationGetAlertGroupsCmd)

	return operationGroupAlertgroupCmd, nil
}
func makeOperationGroupGeneralCmd() (*cobra.Command, error) {
	operationGroupGeneralCmd := &cobra.Command{
		Use: "general",
	}

	operationGetStatusCmd, err := makeOperationGeneralGetStatusCmd()
	if err != nil {
		return nil, err
	}
	operationGroupGeneralCmd.AddCommand(operationGetStatusCmd)

	return operationGroupGeneralCmd, nil
}
func makeOperationGroupReceiverCmd() (*cobra.Command, error) {
	operationGroupReceiverCmd := &cobra.Command{
		Use: "receiver",
	}

	operationGetReceiversCmd, err := makeOperationReceiverGetReceiversCmd()
	if err != nil {
		return nil, err
	}
	operationGroupReceiverCmd.AddCommand(operationGetReceiversCmd)

	return operationGroupReceiverCmd, nil
}
func makeOperationGroupSilenceCmd() (*cobra.Command, error) {
	operationGroupSilenceCmd := &cobra.Command{
		Use: "silence",
	}

	operationDeleteSilenceCmd, err := makeOperationSilenceDeleteSilenceCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSilenceCmd.AddCommand(operationDeleteSilenceCmd)

	operationGetSilenceCmd, err := makeOperationSilenceGetSilenceCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSilenceCmd.AddCommand(operationGetSilenceCmd)

	operationGetSilencesCmd, err := makeOperationSilenceGetSilencesCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSilenceCmd.AddCommand(operationGetSilencesCmd)

	operationPostSilencesCmd, err := makeOperationSilencePostSilencesCmd()
	if err != nil {
		return nil, err
	}
	operationGroupSilenceCmd.AddCommand(operationPostSilencesCmd)

	return operationGroupSilenceCmd, nil
}
