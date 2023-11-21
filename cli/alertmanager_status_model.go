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

// Schema cli for AlertmanagerStatus

// register flags to command
func registerModelAlertmanagerStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAlertmanagerStatusCluster(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAlertmanagerStatusConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAlertmanagerStatusUptime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAlertmanagerStatusVersionInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAlertmanagerStatusCluster(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var clusterFlagName string
	if cmdPrefix == "" {
		clusterFlagName = "cluster"
	} else {
		clusterFlagName = fmt.Sprintf("%v.cluster", cmdPrefix)
	}

	if err := registerModelClusterStatusFlags(depth+1, clusterFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAlertmanagerStatusConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var configFlagName string
	if cmdPrefix == "" {
		configFlagName = "config"
	} else {
		configFlagName = fmt.Sprintf("%v.config", cmdPrefix)
	}

	if err := registerModelAlertmanagerConfigFlags(depth+1, configFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAlertmanagerStatusUptime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uptimeDescription := `Required. `

	var uptimeFlagName string
	if cmdPrefix == "" {
		uptimeFlagName = "uptime"
	} else {
		uptimeFlagName = fmt.Sprintf("%v.uptime", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(uptimeFlagName, "", uptimeDescription)

	return nil
}

func registerAlertmanagerStatusVersionInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var versionInfoFlagName string
	if cmdPrefix == "" {
		versionInfoFlagName = "versionInfo"
	} else {
		versionInfoFlagName = fmt.Sprintf("%v.versionInfo", cmdPrefix)
	}

	if err := registerModelVersionInfoFlags(depth+1, versionInfoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAlertmanagerStatusFlags(depth int, m *models.AlertmanagerStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, clusterAdded := retrieveAlertmanagerStatusClusterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterAdded

	err, configAdded := retrieveAlertmanagerStatusConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configAdded

	err, uptimeAdded := retrieveAlertmanagerStatusUptimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uptimeAdded

	err, versionInfoAdded := retrieveAlertmanagerStatusVersionInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionInfoAdded

	return nil, retAdded
}

func retrieveAlertmanagerStatusClusterFlags(depth int, m *models.AlertmanagerStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterFlagName := fmt.Sprintf("%v.cluster", cmdPrefix)
	if cmd.Flags().Changed(clusterFlagName) {
		// info: complex object cluster ClusterStatus is retrieved outside this Changed() block
	}
	clusterFlagValue := m.Cluster
	if swag.IsZero(clusterFlagValue) {
		clusterFlagValue = &models.ClusterStatus{}
	}

	err, clusterAdded := retrieveModelClusterStatusFlags(depth+1, clusterFlagValue, clusterFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterAdded
	if clusterAdded {
		m.Cluster = clusterFlagValue
	}

	return nil, retAdded
}

func retrieveAlertmanagerStatusConfigFlags(depth int, m *models.AlertmanagerStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	configFlagName := fmt.Sprintf("%v.config", cmdPrefix)
	if cmd.Flags().Changed(configFlagName) {
		// info: complex object config AlertmanagerConfig is retrieved outside this Changed() block
	}
	configFlagValue := m.Config
	if swag.IsZero(configFlagValue) {
		configFlagValue = &models.AlertmanagerConfig{}
	}

	err, configAdded := retrieveModelAlertmanagerConfigFlags(depth+1, configFlagValue, configFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configAdded
	if configAdded {
		m.Config = configFlagValue
	}

	return nil, retAdded
}

func retrieveAlertmanagerStatusUptimeFlags(depth int, m *models.AlertmanagerStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uptimeFlagName := fmt.Sprintf("%v.uptime", cmdPrefix)
	if cmd.Flags().Changed(uptimeFlagName) {

		var uptimeFlagName string
		if cmdPrefix == "" {
			uptimeFlagName = "uptime"
		} else {
			uptimeFlagName = fmt.Sprintf("%v.uptime", cmdPrefix)
		}

		uptimeFlagValueStr, err := cmd.Flags().GetString(uptimeFlagName)
		if err != nil {
			return err, false
		}
		var uptimeFlagValue strfmt.DateTime
		if err := uptimeFlagValue.UnmarshalText([]byte(uptimeFlagValueStr)); err != nil {
			return err, false
		}
		m.Uptime = &uptimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAlertmanagerStatusVersionInfoFlags(depth int, m *models.AlertmanagerStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionInfoFlagName := fmt.Sprintf("%v.versionInfo", cmdPrefix)
	if cmd.Flags().Changed(versionInfoFlagName) {
		// info: complex object versionInfo VersionInfo is retrieved outside this Changed() block
	}
	versionInfoFlagValue := m.VersionInfo
	if swag.IsZero(versionInfoFlagValue) {
		versionInfoFlagValue = &models.VersionInfo{}
	}

	err, versionInfoAdded := retrieveModelVersionInfoFlags(depth+1, versionInfoFlagValue, versionInfoFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionInfoAdded
	if versionInfoAdded {
		m.VersionInfo = versionInfoFlagValue
	}

	return nil, retAdded
}
