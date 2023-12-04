// cli/get_silence_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/spf13/cobra"
)

// makeOperationSilenceGetSilenceCmd returns a cmd to handle operation getSilence
func makeOperationSilenceGetSilenceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSilence",
		Short: `Get a silence by its ID`,
		RunE:  runOperationSilenceGetSilence,
	}

	if err := registerOperationSilenceGetSilenceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSilenceGetSilence uses cmd flags to call endpoint api
func runOperationSilenceGetSilence(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := silence.NewGetSilenceParams()
	if err, _ := retrieveOperationSilenceGetSilenceSilenceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSilenceGetSilenceResult(appCli.Silence.GetSilence(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSilenceGetSilenceParamFlags registers all flags needed to fill params
func registerOperationSilenceGetSilenceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSilenceGetSilenceSilenceIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSilenceGetSilenceSilenceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	silenceIdDescription := `Required. ID of the silence to get`

	var silenceIdFlagName string
	if cmdPrefix == "" {
		silenceIdFlagName = "silenceID"
	} else {
		silenceIdFlagName = fmt.Sprintf("%v.silenceID", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(silenceIdFlagName, "", silenceIdDescription)

	return nil
}

func retrieveOperationSilenceGetSilenceSilenceIDFlag(m *silence.GetSilenceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("silenceID") {

		var silenceIdFlagName string
		if cmdPrefix == "" {
			silenceIdFlagName = "silenceID"
		} else {
			silenceIdFlagName = fmt.Sprintf("%v.silenceID", cmdPrefix)
		}

		silenceIdFlagValueStr, err := cmd.Flags().GetString(silenceIdFlagName)
		if err != nil {
			return err, false
		}
		var silenceIdFlagValue strfmt.UUID
		if err := silenceIdFlagValue.UnmarshalText([]byte(silenceIdFlagValueStr)); err != nil {
			return err, false
		}
		m.SilenceID = silenceIdFlagValue

	}
	return nil, retAdded
}

// parseOperationSilenceGetSilenceResult parses request result and return the string content
func parseOperationSilenceGetSilenceResult(resp0 *silence.GetSilenceOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*silence.GetSilenceOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getSilenceNotFound is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*silence.GetSilenceInternalServerError)
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
