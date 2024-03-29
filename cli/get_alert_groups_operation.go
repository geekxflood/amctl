// cli/get_alert_groups_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/alertgroup"
	"github.com/spf13/cobra"
)

// makeOperationAlertgroupGetAlertGroupsCmd returns a cmd to handle operation getAlertGroups
func makeOperationAlertgroupGetAlertGroupsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getAlertGroups",
		Short: `Get a list of alert groups`,
		RunE:  runOperationAlertgroupGetAlertGroups,
	}

	if err := registerOperationAlertgroupGetAlertGroupsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAlertgroupGetAlertGroups uses cmd flags to call endpoint api
func runOperationAlertgroupGetAlertGroups(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := alertgroup.NewGetAlertGroupsParams()
	if err, _ := retrieveOperationAlertgroupGetAlertGroupsActiveFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertgroupGetAlertGroupsFilterFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertgroupGetAlertGroupsInhibitedFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertgroupGetAlertGroupsReceiverFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationAlertgroupGetAlertGroupsSilencedFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAlertgroupGetAlertGroupsResult(appCli.Alertgroup.GetAlertGroups(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAlertgroupGetAlertGroupsParamFlags registers all flags needed to fill params
func registerOperationAlertgroupGetAlertGroupsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAlertgroupGetAlertGroupsActiveParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertgroupGetAlertGroupsFilterParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertgroupGetAlertGroupsInhibitedParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertgroupGetAlertGroupsReceiverParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationAlertgroupGetAlertGroupsSilencedParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAlertgroupGetAlertGroupsActiveParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAlertgroupGetAlertGroupsFilterParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAlertgroupGetAlertGroupsInhibitedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAlertgroupGetAlertGroupsReceiverParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationAlertgroupGetAlertGroupsSilencedParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationAlertgroupGetAlertGroupsActiveFlag(m *alertgroup.GetAlertGroupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAlertgroupGetAlertGroupsFilterFlag(m *alertgroup.GetAlertGroupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAlertgroupGetAlertGroupsInhibitedFlag(m *alertgroup.GetAlertGroupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAlertgroupGetAlertGroupsReceiverFlag(m *alertgroup.GetAlertGroupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationAlertgroupGetAlertGroupsSilencedFlag(m *alertgroup.GetAlertGroupsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationAlertgroupGetAlertGroupsResult parses request result and return the string content
func parseOperationAlertgroupGetAlertGroupsResult(resp0 *alertgroup.GetAlertGroupsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*alertgroup.GetAlertGroupsOK)
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
		resp1, ok := iResp1.(*alertgroup.GetAlertGroupsBadRequest)
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
		resp2, ok := iResp2.(*alertgroup.GetAlertGroupsInternalServerError)
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
