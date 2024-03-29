// cli/get_alerts_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/spf13/cobra"
)

// makeOperationAlertGetAlertsCmd returns a cmd to handle operation getAlerts
func makeOperationAlertGetAlertsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getAlerts",
		Short: `Get a list of alerts`,
		RunE:  runOperationAlertGetAlerts,
	}

	if err := registerOperationAlertGetAlertsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAlertGetAlerts uses cmd flags to call endpoint api
func runOperationAlertGetAlerts(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := alert.NewGetAlertsParams()
	if err, _ := retrieveOperationAlertGetAlertsActiveFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertGetAlertsFilterFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertGetAlertsInhibitedFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertGetAlertsReceiverFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertGetAlertsSilencedFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertGetAlertsUnprocessedFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAlertGetAlertsResult(appCli.Alert.GetAlerts(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAlertGetAlertsParamFlags registers all flags needed to fill params
func registerOperationAlertGetAlertsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAlertGetAlertsActiveParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertGetAlertsFilterParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertGetAlertsInhibitedParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertGetAlertsReceiverParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertGetAlertsSilencedParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertGetAlertsUnprocessedParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAlertGetAlertsActiveParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	activeDescription := `Show active alerts`

	var activeFlagName string
	if cmdPrefix == "" {
		activeFlagName = "active"
	} else {
		activeFlagName = fmt.Sprintf("%v.active", cmdPrefix)
	}

	var activeFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(activeFlagName, activeFlagDefault, activeDescription)

	return nil
}
func registerOperationAlertGetAlertsFilterParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDescription := `A list of matchers to filter alerts by`

	var filterFlagName string
	if cmdPrefix == "" {
		filterFlagName = "filter"
	} else {
		filterFlagName = fmt.Sprintf("%v.filter", cmdPrefix)
	}

	var filterFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(filterFlagName, filterFlagDefault, filterDescription)

	return nil
}
func registerOperationAlertGetAlertsInhibitedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	inhibitedDescription := `Show inhibited alerts`

	var inhibitedFlagName string
	if cmdPrefix == "" {
		inhibitedFlagName = "inhibited"
	} else {
		inhibitedFlagName = fmt.Sprintf("%v.inhibited", cmdPrefix)
	}

	var inhibitedFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(inhibitedFlagName, inhibitedFlagDefault, inhibitedDescription)

	return nil
}
func registerOperationAlertGetAlertsReceiverParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	receiverDescription := `A regex matching receivers to filter alerts by`

	var receiverFlagName string
	if cmdPrefix == "" {
		receiverFlagName = "receiver"
	} else {
		receiverFlagName = fmt.Sprintf("%v.receiver", cmdPrefix)
	}

	var receiverFlagDefault string

	_ = cmd.PersistentFlags().String(receiverFlagName, receiverFlagDefault, receiverDescription)

	return nil
}
func registerOperationAlertGetAlertsSilencedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	silencedDescription := `Show silenced alerts`

	var silencedFlagName string
	if cmdPrefix == "" {
		silencedFlagName = "silenced"
	} else {
		silencedFlagName = fmt.Sprintf("%v.silenced", cmdPrefix)
	}

	var silencedFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(silencedFlagName, silencedFlagDefault, silencedDescription)

	return nil
}
func registerOperationAlertGetAlertsUnprocessedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	unprocessedDescription := `Show unprocessed alerts`

	var unprocessedFlagName string
	if cmdPrefix == "" {
		unprocessedFlagName = "unprocessed"
	} else {
		unprocessedFlagName = fmt.Sprintf("%v.unprocessed", cmdPrefix)
	}

	var unprocessedFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(unprocessedFlagName, unprocessedFlagDefault, unprocessedDescription)

	return nil
}

func retrieveOperationAlertGetAlertsActiveFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("active") {

		var activeFlagName string
		if cmdPrefix == "" {
			activeFlagName = "active"
		} else {
			activeFlagName = fmt.Sprintf("%v.active", cmdPrefix)
		}

		activeFlagValue, err := cmd.Flags().GetBool(activeFlagName)
		if err != nil {
			return err, false
		}
		m.Active = &activeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAlertGetAlertsFilterFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filter") {

		var filterFlagName string
		if cmdPrefix == "" {
			filterFlagName = "filter"
		} else {
			filterFlagName = fmt.Sprintf("%v.filter", cmdPrefix)
		}

		filterFlagValues, err := cmd.Flags().GetStringSlice(filterFlagName)
		if err != nil {
			return err, false
		}
		m.Filter = filterFlagValues

	}
	return nil, retAdded
}
func retrieveOperationAlertGetAlertsInhibitedFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("inhibited") {

		var inhibitedFlagName string
		if cmdPrefix == "" {
			inhibitedFlagName = "inhibited"
		} else {
			inhibitedFlagName = fmt.Sprintf("%v.inhibited", cmdPrefix)
		}

		inhibitedFlagValue, err := cmd.Flags().GetBool(inhibitedFlagName)
		if err != nil {
			return err, false
		}
		m.Inhibited = &inhibitedFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAlertGetAlertsReceiverFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("receiver") {

		var receiverFlagName string
		if cmdPrefix == "" {
			receiverFlagName = "receiver"
		} else {
			receiverFlagName = fmt.Sprintf("%v.receiver", cmdPrefix)
		}

		receiverFlagValue, err := cmd.Flags().GetString(receiverFlagName)
		if err != nil {
			return err, false
		}
		m.Receiver = &receiverFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAlertGetAlertsSilencedFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("silenced") {

		var silencedFlagName string
		if cmdPrefix == "" {
			silencedFlagName = "silenced"
		} else {
			silencedFlagName = fmt.Sprintf("%v.silenced", cmdPrefix)
		}

		silencedFlagValue, err := cmd.Flags().GetBool(silencedFlagName)
		if err != nil {
			return err, false
		}
		m.Silenced = &silencedFlagValue

	}
	return nil, retAdded
}
func retrieveOperationAlertGetAlertsUnprocessedFlag(m *alert.GetAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("unprocessed") {

		var unprocessedFlagName string
		if cmdPrefix == "" {
			unprocessedFlagName = "unprocessed"
		} else {
			unprocessedFlagName = fmt.Sprintf("%v.unprocessed", cmdPrefix)
		}

		unprocessedFlagValue, err := cmd.Flags().GetBool(unprocessedFlagName)
		if err != nil {
			return err, false
		}
		m.Unprocessed = &unprocessedFlagValue

	}
	return nil, retAdded
}

// parseOperationAlertGetAlertsResult parses request result and return the string content
func parseOperationAlertGetAlertsResult(resp0 *alert.GetAlertsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*alert.GetAlertsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*alert.GetAlertsBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*alert.GetAlertsInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
