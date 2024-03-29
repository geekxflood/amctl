// cli/get_silences_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/spf13/cobra"
)

// makeOperationSilenceGetSilencesCmd returns a cmd to handle operation getSilences
func makeOperationSilenceGetSilencesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSilences",
		Short: `Get a list of silences`,
		RunE:  runOperationSilenceGetSilences,
	}

	if err := registerOperationSilenceGetSilencesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSilenceGetSilences uses cmd flags to call endpoint api
func runOperationSilenceGetSilences(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := silence.NewGetSilencesParams()
	if err, _ := retrieveOperationSilenceGetSilencesFilterFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSilenceGetSilencesResult(appCli.Silence.GetSilences(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSilenceGetSilencesParamFlags registers all flags needed to fill params
func registerOperationSilenceGetSilencesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSilenceGetSilencesFilterParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSilenceGetSilencesFilterParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filterDescription := `A list of matchers to filter silences by`

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

func retrieveOperationSilenceGetSilencesFilterFlag(m *silence.GetSilencesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSilenceGetSilencesResult parses request result and return the string content
func parseOperationSilenceGetSilencesResult(resp0 *silence.GetSilencesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*silence.GetSilencesOK)
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
		resp1, ok := iResp1.(*silence.GetSilencesInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
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
