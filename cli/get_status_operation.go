// cli/get_status_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/spf13/cobra"
)

// makeOperationGeneralGetStatusCmd returns a cmd to handle operation getStatus
func makeOperationGeneralGetStatusCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStatus",
		Short: `Get current status of an Alertmanager instance and its cluster`,
		RunE:  runOperationGeneralGetStatus,
	}

	if err := registerOperationGeneralGetStatusParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationGeneralGetStatus uses cmd flags to call endpoint api
func runOperationGeneralGetStatus(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := general.NewGetStatusParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationGeneralGetStatusResult(appCli.General.GetStatus(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationGeneralGetStatusParamFlags registers all flags needed to fill params
func registerOperationGeneralGetStatusParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationGeneralGetStatusResult parses request result and return the string content
func parseOperationGeneralGetStatusResult(resp0 *general.GetStatusOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*general.GetStatusOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
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
