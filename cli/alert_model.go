// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/geekxflood/alertmanager-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Alert

// register flags to command
func registerModelAlertFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAlertGeneratorURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAlertLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAlertGeneratorURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive generatorURL strfmt.URI is not supported by go-swagger cli yet

	return nil
}

func registerAlertLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: labels LabelSet map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAlertFlags(depth int, m *models.Alert, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, generatorUrlAdded := retrieveAlertGeneratorURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || generatorUrlAdded

	err, labelsAdded := retrieveAlertLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	return nil, retAdded
}

func retrieveAlertGeneratorURLFlags(depth int, m *models.Alert, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	generatorUrlFlagName := fmt.Sprintf("%v.generatorURL", cmdPrefix)
	if cmd.Flags().Changed(generatorUrlFlagName) {

		// warning: primitive generatorURL strfmt.URI is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAlertLabelsFlags(depth int, m *models.Alert, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	labelsFlagName := fmt.Sprintf("%v.labels", cmdPrefix)
	if cmd.Flags().Changed(labelsFlagName) {
		// warning: labels map type LabelSet is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
