// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/geekxflood/alertmanager-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for AlertmanagerConfig

// register flags to command
func registerModelAlertmanagerConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAlertmanagerConfigOriginal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAlertmanagerConfigOriginal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	originalDescription := `Required. `

	var originalFlagName string
	if cmdPrefix == "" {
		originalFlagName = "original"
	} else {
		originalFlagName = fmt.Sprintf("%v.original", cmdPrefix)
	}

	var originalFlagDefault string

	_ = cmd.PersistentFlags().String(originalFlagName, originalFlagDefault, originalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAlertmanagerConfigFlags(depth int, m *models.AlertmanagerConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, originalAdded := retrieveAlertmanagerConfigOriginalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || originalAdded

	return nil, retAdded
}

func retrieveAlertmanagerConfigOriginalFlags(depth int, m *models.AlertmanagerConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	originalFlagName := fmt.Sprintf("%v.original", cmdPrefix)
	if cmd.Flags().Changed(originalFlagName) {

		var originalFlagName string
		if cmdPrefix == "" {
			originalFlagName = "original"
		} else {
			originalFlagName = fmt.Sprintf("%v.original", cmdPrefix)
		}

		originalFlagValue, err := cmd.Flags().GetString(originalFlagName)
		if err != nil {
			return err, false
		}
		m.Original = &originalFlagValue

		retAdded = true
	}

	return nil, retAdded
}