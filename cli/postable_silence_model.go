// cli/postable_silence_model.go

package cli

import (
	"fmt"

	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for PostableSilence

// register flags to command
func registerModelPostableSilenceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register anonymous fields for AO0

	if err := registerPostableSilenceAnonAO0ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register embedded Silence flags

	if err := registerModelSilenceFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name AO0, type

func registerPostableSilenceAnonAO0ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostableSilenceFlags(depth int, m *models.PostableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve allOf AO0 fields

	err, idAdded := retrievePostableSilenceAnonAO0IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	// retrieve model Silence
	err, aO1Added := retrieveModelSilenceFlags(depth, &m.Silence, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aO1Added

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name AO0

func retrievePostableSilenceAnonAO0IDFlags(depth int, m *models.PostableSilence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}
