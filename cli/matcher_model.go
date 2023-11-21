// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/geekxflood/alertmanager-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Matcher

// register flags to command
func registerModelMatcherFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMatcherIsEqual(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMatcherIsRegex(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMatcherName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMatcherValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMatcherIsEqual(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isEqualDescription := ``

	var isEqualFlagName string
	if cmdPrefix == "" {
		isEqualFlagName = "isEqual"
	} else {
		isEqualFlagName = fmt.Sprintf("%v.isEqual", cmdPrefix)
	}

	var isEqualFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(isEqualFlagName, isEqualFlagDefault, isEqualDescription)

	return nil
}

func registerMatcherIsRegex(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	isRegexDescription := `Required. `

	var isRegexFlagName string
	if cmdPrefix == "" {
		isRegexFlagName = "isRegex"
	} else {
		isRegexFlagName = fmt.Sprintf("%v.isRegex", cmdPrefix)
	}

	var isRegexFlagDefault bool

	_ = cmd.PersistentFlags().Bool(isRegexFlagName, isRegexFlagDefault, isRegexDescription)

	return nil
}

func registerMatcherName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerMatcherValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := `Required. `

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMatcherFlags(depth int, m *models.Matcher, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, isEqualAdded := retrieveMatcherIsEqualFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isEqualAdded

	err, isRegexAdded := retrieveMatcherIsRegexFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || isRegexAdded

	err, nameAdded := retrieveMatcherNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, valueAdded := retrieveMatcherValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveMatcherIsEqualFlags(depth int, m *models.Matcher, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isEqualFlagName := fmt.Sprintf("%v.isEqual", cmdPrefix)
	if cmd.Flags().Changed(isEqualFlagName) {

		var isEqualFlagName string
		if cmdPrefix == "" {
			isEqualFlagName = "isEqual"
		} else {
			isEqualFlagName = fmt.Sprintf("%v.isEqual", cmdPrefix)
		}

		isEqualFlagValue, err := cmd.Flags().GetBool(isEqualFlagName)
		if err != nil {
			return err, false
		}
		m.IsEqual = &isEqualFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMatcherIsRegexFlags(depth int, m *models.Matcher, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	isRegexFlagName := fmt.Sprintf("%v.isRegex", cmdPrefix)
	if cmd.Flags().Changed(isRegexFlagName) {

		var isRegexFlagName string
		if cmdPrefix == "" {
			isRegexFlagName = "isRegex"
		} else {
			isRegexFlagName = fmt.Sprintf("%v.isRegex", cmdPrefix)
		}

		isRegexFlagValue, err := cmd.Flags().GetBool(isRegexFlagName)
		if err != nil {
			return err, false
		}
		m.IsRegex = &isRegexFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMatcherNameFlags(depth int, m *models.Matcher, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMatcherValueFlags(depth int, m *models.Matcher, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = &valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}
