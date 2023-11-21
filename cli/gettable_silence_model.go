// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
	"github.com/geekxflood/alertmanager-cli/models"

)

// Schema cli for GettableSilence

// register flags to command
func registerModelGettableSilenceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register anonymous fields for AO0

	if err := registerGettableSilenceAnonAO0ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGettableSilenceAnonAO0Status(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGettableSilenceAnonAO0UpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register embedded Silence flags

	if err := registerModelSilenceFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name AO0, type

func registerGettableSilenceAnonAO0ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerGettableSilenceAnonAO0Status(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	if err := registerModelSilenceStatusFlags(depth+1, statusFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerGettableSilenceAnonAO0UpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := `Required. `

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "updatedAt"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.updatedAt", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updatedAtFlagName, "", updatedAtDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGettableSilenceFlags(depth int, m *models.GettableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve allOf AO0 fields

	err, idAdded := retrieveGettableSilenceAnonAO0IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, statusAdded := retrieveGettableSilenceAnonAO0StatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, updatedAtAdded := retrieveGettableSilenceAnonAO0UpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	// retrieve model Silence
	err, aO1Added := retrieveModelSilenceFlags(depth, &m.Silence, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aO1Added

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name AO0

func retrieveGettableSilenceAnonAO0IDFlags(depth int, m *models.GettableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGettableSilenceAnonAO0StatusFlags(depth int, m *models.GettableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {
		// info: complex object status SilenceStatus is retrieved outside this Changed() block
	}
	statusFlagValue := m.Status
	if swag.IsZero(statusFlagValue) {
		statusFlagValue = &models.SilenceStatus{}
	}

	err, statusAdded := retrieveModelSilenceStatusFlags(depth+1, statusFlagValue, statusFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded
	if statusAdded {
		m.Status = statusFlagValue
	}

	return nil, retAdded
}

func retrieveGettableSilenceAnonAO0UpdatedAtFlags(depth int, m *models.GettableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedAtFlagName := fmt.Sprintf("%v.updatedAt", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		var updatedAtFlagName string
		if cmdPrefix == "" {
			updatedAtFlagName = "updatedAt"
		} else {
			updatedAtFlagName = fmt.Sprintf("%v.updatedAt", cmdPrefix)
		}

		updatedAtFlagValueStr, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		var updatedAtFlagValue strfmt.DateTime
		if err := updatedAtFlagValue.UnmarshalText([]byte(updatedAtFlagValueStr)); err != nil {
			return err, false
		}
		m.UpdatedAt = &updatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}