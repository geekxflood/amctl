// cli/silence_model.go

package cli

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/spf13/cobra"
)

// Schema cli for Silence

// register flags to command
func registerModelSilenceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSilenceComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSilenceCreatedBy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSilenceEndsAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSilenceMatchers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSilenceStartsAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSilenceComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := `Required. `

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerSilenceCreatedBy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdByDescription := `Required. `

	var createdByFlagName string
	if cmdPrefix == "" {
		createdByFlagName = "createdBy"
	} else {
		createdByFlagName = fmt.Sprintf("%v.createdBy", cmdPrefix)
	}

	var createdByFlagDefault string

	_ = cmd.PersistentFlags().String(createdByFlagName, createdByFlagDefault, createdByDescription)

	return nil
}

func registerSilenceEndsAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endsAtDescription := `Required. `

	var endsAtFlagName string
	if cmdPrefix == "" {
		endsAtFlagName = "endsAt"
	} else {
		endsAtFlagName = fmt.Sprintf("%v.endsAt", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(endsAtFlagName, "", endsAtDescription)

	return nil
}

func registerSilenceMatchers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: matchers Matchers array type is not supported by go-swagger cli yet

	return nil
}

func registerSilenceStartsAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startsAtDescription := `Required. `

	var startsAtFlagName string
	if cmdPrefix == "" {
		startsAtFlagName = "startsAt"
	} else {
		startsAtFlagName = fmt.Sprintf("%v.startsAt", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(startsAtFlagName, "", startsAtDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSilenceFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commentAdded := retrieveSilenceCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, createdByAdded := retrieveSilenceCreatedByFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdByAdded

	err, endsAtAdded := retrieveSilenceEndsAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endsAtAdded

	err, matchersAdded := retrieveSilenceMatchersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || matchersAdded

	err, startsAtAdded := retrieveSilenceStartsAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startsAtAdded

	return nil, retAdded
}

func retrieveSilenceCommentFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commentFlagName := fmt.Sprintf("%v.comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = &commentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSilenceCreatedByFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdByFlagName := fmt.Sprintf("%v.createdBy", cmdPrefix)
	if cmd.Flags().Changed(createdByFlagName) {

		var createdByFlagName string
		if cmdPrefix == "" {
			createdByFlagName = "createdBy"
		} else {
			createdByFlagName = fmt.Sprintf("%v.createdBy", cmdPrefix)
		}

		createdByFlagValue, err := cmd.Flags().GetString(createdByFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedBy = &createdByFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSilenceEndsAtFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endsAtFlagName := fmt.Sprintf("%v.endsAt", cmdPrefix)
	if cmd.Flags().Changed(endsAtFlagName) {

		var endsAtFlagName string
		if cmdPrefix == "" {
			endsAtFlagName = "endsAt"
		} else {
			endsAtFlagName = fmt.Sprintf("%v.endsAt", cmdPrefix)
		}

		endsAtFlagValueStr, err := cmd.Flags().GetString(endsAtFlagName)
		if err != nil {
			return err, false
		}
		var endsAtFlagValue strfmt.DateTime
		if err := endsAtFlagValue.UnmarshalText([]byte(endsAtFlagValueStr)); err != nil {
			return err, false
		}
		m.EndsAt = &endsAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSilenceMatchersFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	return nil, retAdded
}

func retrieveSilenceStartsAtFlags(depth int, m *models.Silence, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startsAtFlagName := fmt.Sprintf("%v.startsAt", cmdPrefix)
	if cmd.Flags().Changed(startsAtFlagName) {

		var startsAtFlagName string
		if cmdPrefix == "" {
			startsAtFlagName = "startsAt"
		} else {
			startsAtFlagName = fmt.Sprintf("%v.startsAt", cmdPrefix)
		}

		startsAtFlagValueStr, err := cmd.Flags().GetString(startsAtFlagName)
		if err != nil {
			return err, false
		}
		var startsAtFlagValue strfmt.DateTime
		if err := startsAtFlagValue.UnmarshalText([]byte(startsAtFlagValueStr)); err != nil {
			return err, false
		}
		m.StartsAt = &startsAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}
