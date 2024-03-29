// cli/post_silences_operation.go

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/prometheus/alertmanager/api/v2/client/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/spf13/cobra"
)

// makeOperationSilencePostSilencesCmd returns a cmd to handle operation postSilences
func makeOperationSilencePostSilencesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "postSilences",
		Short: `Post a new silence or update an existing one`,
		RunE:  runOperationSilencePostSilences,
	}

	if err := registerOperationSilencePostSilencesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSilencePostSilences uses cmd flags to call endpoint api
func runOperationSilencePostSilences(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := silence.NewPostSilencesParams()
	if err, _ := retrieveOperationSilencePostSilencesSilenceFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSilencePostSilencesResult(appCli.Silence.PostSilences(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSilencePostSilencesParamFlags registers all flags needed to fill params
func registerOperationSilencePostSilencesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSilencePostSilencesSilenceParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSilencePostSilencesSilenceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var silenceFlagName string
	if cmdPrefix == "" {
		silenceFlagName = "silence"
	} else {
		silenceFlagName = fmt.Sprintf("%v.silence", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(silenceFlagName, "", "Optional json string for [silence]. The silence to create")

	// add flags for body
	if err := registerModelPostableSilenceFlags(0, "postableSilence", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationSilencePostSilencesSilenceFlag(m *silence.PostSilencesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("silence") {
		// Read silence string from cmd and unmarshal
		silenceValueStr, err := cmd.Flags().GetString("silence")
		if err != nil {
			return err, false
		}

		silenceValue := models.PostableSilence{}
		if err := json.Unmarshal([]byte(silenceValueStr), &silenceValue); err != nil {
			return fmt.Errorf("cannot unmarshal silence string in models.PostableSilence: %v", err), false
		}
		m.Silence = &silenceValue
	}
	silenceValueModel := m.Silence
	if swag.IsZero(silenceValueModel) {
		silenceValueModel = &models.PostableSilence{}
	}
	err, added := retrieveModelPostableSilenceFlags(0, silenceValueModel, "postableSilence", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Silence = silenceValueModel
	}
	if dryRun && debug {

		silenceValueDebugBytes, err := json.Marshal(m.Silence)
		if err != nil {
			return err, false
		}
		logDebugf("Silence dry-run payload: %v", string(silenceValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationSilencePostSilencesResult parses request result and return the string content
func parseOperationSilencePostSilencesResult(resp0 *silence.PostSilencesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*silence.PostSilencesOK)
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
		resp1, ok := iResp1.(*silence.PostSilencesBadRequest)
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
		resp2, ok := iResp2.(*silence.PostSilencesNotFound)
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
