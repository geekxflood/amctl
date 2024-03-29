// cli/delete_silence_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/spf13/cobra"
)

// makeOperationSilenceDeleteSilenceCmd returns a cmd to handle operation deleteSilence
func makeOperationSilenceDeleteSilenceCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteSilence",
		Short: `Delete a silence by its ID`,
		RunE:  runOperationSilenceDeleteSilence,
	}

	if err := registerOperationSilenceDeleteSilenceParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSilenceDeleteSilence uses cmd flags to call endpoint api
func runOperationSilenceDeleteSilence(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := silence.NewDeleteSilenceParams()
	if err, _ := retrieveOperationSilenceDeleteSilenceSilenceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSilenceDeleteSilenceResult(appCli.Silence.DeleteSilence(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSilenceDeleteSilenceParamFlags registers all flags needed to fill params
func registerOperationSilenceDeleteSilenceParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSilenceDeleteSilenceSilenceIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSilenceDeleteSilenceSilenceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSilenceDeleteSilenceSilenceIDFlag(m *silence.DeleteSilenceParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSilenceDeleteSilenceResult parses request result and return the string content
func parseOperationSilenceDeleteSilenceResult(resp0 *silence.DeleteSilenceOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteSilenceOK is not supported

		// Non schema case: warning deleteSilenceNotFound is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*silence.DeleteSilenceInternalServerError)
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

	// warning: non schema response deleteSilenceOK is not supported by go-swagger cli yet.

	return "", nil
}
