// cli/post_alerts_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/alert"
	"github.com/spf13/cobra"
)

// makeOperationAlertPostAlertsCmd returns a cmd to handle operation postAlerts
func makeOperationAlertPostAlertsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postAlerts",
		Short: `Create new Alerts`,
		RunE:  runOperationAlertPostAlerts,
	}

	if err := registerOperationAlertPostAlertsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationAlertPostAlerts uses cmd flags to call endpoint api
func runOperationAlertPostAlerts(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := alert.NewPostAlertsParams()
	if err, _ := retrieveOperationAlertPostAlertsAlertsFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationAlertPostAlertsResult(appCli.Alert.PostAlerts(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationAlertPostAlertsParamFlags registers all flags needed to fill params
func registerOperationAlertPostAlertsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationAlertPostAlertsAlertsParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationAlertPostAlertsAlertsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	// warning: array alerts models.PostableAlerts is not supported by go-swagger cli yet

	return nil
}

func retrieveOperationAlertPostAlertsAlertsFlag(m *alert.PostAlertsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	return nil, retAdded
}

// parseOperationAlertPostAlertsResult parses request result and return the string content
func parseOperationAlertPostAlertsResult(resp0 *alert.PostAlertsOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning postAlertsOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*alert.PostAlertsBadRequest)
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
		resp2, ok := iResp2.(*alert.PostAlertsInternalServerError)
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

	// warning: non schema response postAlertsOK is not supported by go-swagger cli yet.

	return "", nil
}
