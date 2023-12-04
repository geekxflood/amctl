// cli/get_receivers_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/receiver"
	"github.com/spf13/cobra"
)

// makeOperationReceiverGetReceiversCmd returns a cmd to handle operation getReceivers
func makeOperationReceiverGetReceiversCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getReceivers",
		Short: `Get list of all receivers (name of notification integrations)`,
		RunE:  runOperationReceiverGetReceivers,
	}

	if err := registerOperationReceiverGetReceiversParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationReceiverGetReceivers uses cmd flags to call endpoint api
func runOperationReceiverGetReceivers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := receiver.NewGetReceiversParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationReceiverGetReceiversResult(appCli.Receiver.GetReceivers(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationReceiverGetReceiversParamFlags registers all flags needed to fill params
func registerOperationReceiverGetReceiversParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationReceiverGetReceiversResult parses request result and return the string content
func parseOperationReceiverGetReceiversResult(resp0 *receiver.GetReceiversOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*receiver.GetReceiversOK)
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
