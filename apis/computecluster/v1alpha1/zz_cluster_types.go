/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterObservation struct {

	// : The managed object ID of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The managed object ID of the primary
	// resource pool for this cluster. This can be passed directly to the
	// resource_pool_id
	// attribute of the
	// vsphere_virtual_machine resource.
	// The managed object ID of the cluster's root resource pool.
	ResourcePoolID *string `json:"resourcePoolId,omitempty" tf:"resource_pool_id,omitempty"`
}

type ClusterParameters struct {

	// A map of custom attribute ids to attribute
	// value strings to set for the datastore cluster. See
	// here for a reference on how to set values
	// for custom attributes.
	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The managed object ID of
	// the datacenter to create the cluster in. Forces a new resource if changed.
	// The managed object ID of the datacenter to put the cluster in.
	// +kubebuilder:validation:Required
	DatacenterID *string `json:"datacenterId" tf:"datacenter_id,omitempty"`

	// The automation level for host power
	// operations in this cluster. Can be one of manual or automated. Default:
	// manual.
	// The automation level for host power operations in this cluster. Can be one of manual or automated.
	// +kubebuilder:validation:Optional
	DpmAutomationLevel *string `json:"dpmAutomationLevel,omitempty" tf:"dpm_automation_level,omitempty"`

	// Enable DPM support for DRS in this cluster.
	// Requires drs_enabled to be true in order to be effective.
	// Default: false.
	// Enable DPM support for DRS. This allows you to dynamically control the power of hosts depending on the needs of virtual machines in the cluster. Requires that DRS be enabled.
	// +kubebuilder:validation:Optional
	DpmEnabled *bool `json:"dpmEnabled,omitempty" tf:"dpm_enabled,omitempty"`

	// A value between 1 and 5 indicating the
	// threshold of load within the cluster that influences host power operations.
	// This affects both power on and power off operations - a lower setting will
	// tolerate more of a surplus/deficit than a higher setting. Default: 3.
	// A value between 1 and 5 indicating the threshold of load within the cluster that influences host power operations. This affects both power on and power off operations - a lower setting will tolerate more of a surplus/deficit than a higher setting.
	// +kubebuilder:validation:Optional
	DpmThreshold *float64 `json:"dpmThreshold,omitempty" tf:"dpm_threshold,omitempty"`

	// A key/value map that specifies advanced
	// options for DRS and DPM.
	// Advanced configuration options for DRS and DPM.
	// +kubebuilder:validation:Optional
	DrsAdvancedOptions map[string]*string `json:"drsAdvancedOptions,omitempty" tf:"drs_advanced_options,omitempty"`

	// The default automation level for all
	// virtual machines in this cluster. Can be one of manual,
	// partiallyAutomated, or fullyAutomated. Default: manual.
	// The default automation level for all virtual machines in this cluster. Can be one of manual, partiallyAutomated, or fullyAutomated.
	// +kubebuilder:validation:Optional
	DrsAutomationLevel *string `json:"drsAutomationLevel,omitempty" tf:"drs_automation_level,omitempty"`

	// When true, enables DRS to use data
	// from vRealize Operations Manager to make proactive DRS
	// recommendations. *
	// When true, enables DRS to use data from vRealize Operations Manager to make proactive DRS recommendations.
	// +kubebuilder:validation:Optional
	DrsEnablePredictiveDrs *bool `json:"drsEnablePredictiveDrs,omitempty" tf:"drs_enable_predictive_drs,omitempty"`

	// Allow individual DRS overrides to be
	// set for virtual machines in the cluster. Default: true.
	// When true, allows individual VM overrides within this cluster to be set.
	// +kubebuilder:validation:Optional
	DrsEnableVMOverrides *bool `json:"drsEnableVmOverrides,omitempty" tf:"drs_enable_vm_overrides,omitempty"`

	// Enable DRS for this cluster. Default: false.
	// Enable DRS for this cluster.
	// +kubebuilder:validation:Optional
	DrsEnabled *bool `json:"drsEnabled,omitempty" tf:"drs_enabled,omitempty"`

	// A value between 1 and 5 indicating
	// the threshold of imbalance tolerated between hosts. A lower setting will
	// tolerate more imbalance while a higher setting will tolerate less. Default:
	// 3.
	// A value between 1 and 5 indicating the threshold of imbalance tolerated between hosts. A lower setting will tolerate more imbalance while a higher setting will tolerate less.
	// +kubebuilder:validation:Optional
	DrsMigrationThreshold *float64 `json:"drsMigrationThreshold,omitempty" tf:"drs_migration_threshold,omitempty"`

	// Enable scalable shares for all
	// resource pools in the cluster. Can be one of disabled or
	// scaleCpuAndMemoryShares. Default: disabled.
	// Enable scalable shares for all descendants of this cluster.
	// +kubebuilder:validation:Optional
	DrsScaleDescendantsShares *string `json:"drsScaleDescendantsShares,omitempty" tf:"drs_scale_descendants_shares,omitempty"`

	// The relative path to a folder to put this cluster in.
	// This is a path relative to the datacenter you are deploying the cluster to.
	// The name of the folder to locate the cluster in.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// When destroying the resource, setting this to
	// true will auto-remove any hosts that are currently a member of the cluster,
	// as if they were removed by taking their entry out of host_system_ids (see
	// below). This is an advanced
	// option and should only be used for testing. Default: false.
	// Force removal of all hosts in the cluster during destroy and make them standalone hosts. Use of this flag mainly exists for testing and is not recommended in normal use.
	// +kubebuilder:validation:Optional
	ForceEvacuateOnDestroy *bool `json:"forceEvacuateOnDestroy,omitempty" tf:"force_evacuate_on_destroy,omitempty"`

	// Defines the
	// managed object IDs of hosts to use as dedicated failover
	// hosts. These hosts are kept as available as possible - admission control will
	// block access to the host, and DRS will ignore the host when making
	// recommendations.
	// When ha_admission_control_policy is failoverHosts, this defines the managed object IDs of hosts to use as dedicated failover hosts. These hosts are kept as available as possible - admission control will block access to the host, and DRS will ignore the host when making recommendations.
	// +kubebuilder:validation:Optional
	HaAdmissionControlFailoverHostSystemIds []*string `json:"haAdmissionControlFailoverHostSystemIds,omitempty" tf:"ha_admission_control_failover_host_system_ids,omitempty"`

	// The maximum number
	// of failed hosts that admission control tolerates when making decisions on
	// whether to permit virtual machine operations. The maximum is one less than
	// the number of hosts in the cluster. Default: 1.
	// *
	// The maximum number of failed hosts that admission control tolerates when making decisions on whether to permit virtual machine operations. The maximum is one less than the number of hosts in the cluster.
	// +kubebuilder:validation:Optional
	HaAdmissionControlHostFailureTolerance *float64 `json:"haAdmissionControlHostFailureTolerance,omitempty" tf:"ha_admission_control_host_failure_tolerance,omitempty"`

	// The percentage of
	// resource reduction that a cluster of virtual machines can tolerate in case of
	// a failover. A value of 0 produces warnings only, whereas a value of 100
	// disables the setting. Default: 100 (disabled).
	// The percentage of resource reduction that a cluster of VMs can tolerate in case of a failover. A value of 0 produces warnings only, whereas a value of 100 disables the setting.
	// +kubebuilder:validation:Optional
	HaAdmissionControlPerformanceTolerance *float64 `json:"haAdmissionControlPerformanceTolerance,omitempty" tf:"ha_admission_control_performance_tolerance,omitempty"`

	// The type of admission control
	// policy to use with vSphere HA. Can be one of resourcePercentage,
	// slotPolicy, failoverHosts, or disabled. Default: resourcePercentage.
	// The type of admission control policy to use with vSphere HA, which controls whether or not specific VM operations are permitted in the cluster in order to protect the reliability of the cluster. Can be one of resourcePercentage, slotPolicy, failoverHosts, or disabled. Note that disabling admission control is not recommended and can lead to service issues.
	// +kubebuilder:validation:Optional
	HaAdmissionControlPolicy *string `json:"haAdmissionControlPolicy,omitempty" tf:"ha_admission_control_policy,omitempty"`

	// Automatically determine available resource percentages by subtracting the
	// average number of host resources represented by the
	// ha_admission_control_host_failure_tolerance
	// setting from the total amount of resources in the cluster. Disable to supply
	// user-defined values. Default: true.
	// *
	// When ha_admission_control_policy is resourcePercentage, automatically determine available resource percentages by subtracting the average number of host resources represented by the ha_admission_control_host_failure_tolerance setting from the total amount of resources in the cluster. Disable to supply user-defined values.
	// +kubebuilder:validation:Optional
	HaAdmissionControlResourcePercentageAutoCompute *bool `json:"haAdmissionControlResourcePercentageAutoCompute,omitempty" tf:"ha_admission_control_resource_percentage_auto_compute,omitempty"`

	// Controls the
	// user-defined percentage of CPU resources in the cluster to reserve for
	// failover. Default: 100.
	// When ha_admission_control_policy is resourcePercentage, this controls the user-defined percentage of CPU resources in the cluster to reserve for failover.
	// +kubebuilder:validation:Optional
	HaAdmissionControlResourcePercentageCPU *float64 `json:"haAdmissionControlResourcePercentageCpu,omitempty" tf:"ha_admission_control_resource_percentage_cpu,omitempty"`

	// Controls the
	// user-defined percentage of memory resources in the cluster to reserve for
	// failover. Default: 100.
	// When ha_admission_control_policy is resourcePercentage, this controls the user-defined percentage of memory resources in the cluster to reserve for failover.
	// +kubebuilder:validation:Optional
	HaAdmissionControlResourcePercentageMemory *float64 `json:"haAdmissionControlResourcePercentageMemory,omitempty" tf:"ha_admission_control_resource_percentage_memory,omitempty"`

	// Controls the
	// user-defined CPU slot size, in MHz. Default: 32.
	// When ha_admission_control_policy is slotPolicy, this controls the user-defined CPU slot size, in MHz.
	// +kubebuilder:validation:Optional
	HaAdmissionControlSlotPolicyExplicitCPU *float64 `json:"haAdmissionControlSlotPolicyExplicitCpu,omitempty" tf:"ha_admission_control_slot_policy_explicit_cpu,omitempty"`

	// Controls the
	// user-defined memory slot size, in MB. Default: 100.
	// When ha_admission_control_policy is slotPolicy, this controls the user-defined memory slot size, in MB.
	// +kubebuilder:validation:Optional
	HaAdmissionControlSlotPolicyExplicitMemory *float64 `json:"haAdmissionControlSlotPolicyExplicitMemory,omitempty" tf:"ha_admission_control_slot_policy_explicit_memory,omitempty"`

	// Controls
	// whether or not you wish to supply explicit values to CPU and memory slot
	// sizes. The default is false, which tells vSphere to gather a automatic
	// average based on all powered-on virtual machines currently in the cluster.
	// When ha_admission_control_policy is slotPolicy, this setting controls whether or not you wish to supply explicit values to CPU and memory slot sizes. The default is to gather a automatic average based on all powered-on virtual machines currently in the cluster.
	// +kubebuilder:validation:Optional
	HaAdmissionControlSlotPolicyUseExplicitSize *bool `json:"haAdmissionControlSlotPolicyUseExplicitSize,omitempty" tf:"ha_admission_control_slot_policy_use_explicit_size,omitempty"`

	// A key/value map that specifies advanced
	// options for vSphere HA.
	// Advanced configuration options for vSphere HA.
	// +kubebuilder:validation:Optional
	HaAdvancedOptions map[string]*string `json:"haAdvancedOptions,omitempty" tf:"ha_advanced_options,omitempty"`

	// Controls the action to take
	// on virtual machines if an APD status on an affected datastore clears in the
	// middle of an APD event. Can be one of none or reset. Default: none.
	// *
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines if an APD status on an affected datastore clears in the middle of an APD event. Can be one of none or reset.
	// +kubebuilder:validation:Optional
	HaDatastoreApdRecoveryAction *string `json:"haDatastoreApdRecoveryAction,omitempty" tf:"ha_datastore_apd_recovery_action,omitempty"`

	// Controls the action to take on
	// virtual machines when the cluster has detected loss to all paths to a
	// relevant datastore. Can be one of disabled, warning,
	// restartConservative, or restartAggressive.  Default: disabled.
	// *
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines when the cluster has detected loss to all paths to a relevant datastore. Can be one of disabled, warning, restartConservative, or restartAggressive.
	// +kubebuilder:validation:Optional
	HaDatastoreApdResponse *string `json:"haDatastoreApdResponse,omitempty" tf:"ha_datastore_apd_response,omitempty"`

	// The time, in seconds,
	// to wait after an APD timeout event to run the response action defined in
	// ha_datastore_apd_response. Default: 180
	// seconds (3 minutes). *
	// When ha_vm_component_protection is enabled, controls the delay in seconds to wait after an APD timeout event to execute the response action defined in ha_datastore_apd_response.
	// +kubebuilder:validation:Optional
	HaDatastoreApdResponseDelay *float64 `json:"haDatastoreApdResponseDelay,omitempty" tf:"ha_datastore_apd_response_delay,omitempty"`

	// Controls the action to take on
	// virtual machines when the cluster has detected a permanent device loss to a
	// relevant datastore. Can be one of disabled, warning, or
	// restartAggressive. Default: disabled.
	// *
	// When ha_vm_component_protection is enabled, controls the action to take on virtual machines when the cluster has detected a permanent device loss to a relevant datastore. Can be one of disabled, warning, or restartAggressive.
	// +kubebuilder:validation:Optional
	HaDatastorePdlResponse *string `json:"haDatastorePdlResponse,omitempty" tf:"ha_datastore_pdl_response,omitempty"`

	// Enable vSphere HA for this cluster. Default:
	// false.
	// Enable vSphere HA for this cluster.
	// +kubebuilder:validation:Optional
	HaEnabled *bool `json:"haEnabled,omitempty" tf:"ha_enabled,omitempty"`

	// The list of managed object IDs for
	// preferred datastores to use for HA heartbeating. This setting is only useful
	// when ha_heartbeat_datastore_policy is set
	// to either userSelectedDs or allFeasibleDsWithUserPreference.
	// The list of managed object IDs for preferred datastores to use for HA heartbeating. This setting is only useful when ha_heartbeat_datastore_policy is set to either userSelectedDs or allFeasibleDsWithUserPreference.
	// +kubebuilder:validation:Optional
	HaHeartbeatDatastoreIds []*string `json:"haHeartbeatDatastoreIds,omitempty" tf:"ha_heartbeat_datastore_ids,omitempty"`

	// The selection policy for HA
	// heartbeat datastores. Can be one of allFeasibleDs, userSelectedDs, or
	// allFeasibleDsWithUserPreference. Default:
	// allFeasibleDsWithUserPreference.
	// The selection policy for HA heartbeat datastores. Can be one of allFeasibleDs, userSelectedDs, or allFeasibleDsWithUserPreference.
	// +kubebuilder:validation:Optional
	HaHeartbeatDatastorePolicy *string `json:"haHeartbeatDatastorePolicy,omitempty" tf:"ha_heartbeat_datastore_policy,omitempty"`

	// The action to take on virtual
	// machines when a host has detected that it has been isolated from the rest of
	// the cluster. Can be one of none, powerOff, or shutdown. Default:
	// none.
	// The action to take on virtual machines when a host has detected that it has been isolated from the rest of the cluster. Can be one of none, powerOff, or shutdown.
	// +kubebuilder:validation:Optional
	HaHostIsolationResponse *string `json:"haHostIsolationResponse,omitempty" tf:"ha_host_isolation_response,omitempty"`

	// Global setting that controls whether
	// vSphere HA remediates virtual machines on host failure. Can be one of enabled
	// or disabled. Default: enabled.
	// Global setting that controls whether vSphere HA remediates VMs on host failure. Can be one of enabled or disabled.
	// +kubebuilder:validation:Optional
	HaHostMonitoring *string `json:"haHostMonitoring,omitempty" tf:"ha_host_monitoring,omitempty"`

	// Controls vSphere VM component
	// protection for virtual machines in this cluster. Can be one of enabled or
	// disabled. Default: enabled.
	// *
	// Controls vSphere VM component protection for virtual machines in this cluster. This allows vSphere HA to react to failures between hosts and specific virtual machine components, such as datastores. Can be one of enabled or disabled.
	// +kubebuilder:validation:Optional
	HaVMComponentProtection *string `json:"haVmComponentProtection,omitempty" tf:"ha_vm_component_protection,omitempty"`

	// The condition used to
	// determine whether or not virtual machines in a certain restart priority class
	// are online, allowing HA to move on to restarting virtual machines on the next
	// priority. Can be one of none, poweredOn, guestHbStatusGreen, or
	// appHbStatusGreen. The default is none, which means that a virtual machine
	// is considered ready immediately after a host is found to start it on.
	// *
	// The condition used to determine whether or not VMs in a certain restart priority class are online, allowing HA to move on to restarting VMs on the next priority. Can be one of none, poweredOn, guestHbStatusGreen, or appHbStatusGreen.
	// +kubebuilder:validation:Optional
	HaVMDependencyRestartCondition *string `json:"haVmDependencyRestartCondition,omitempty" tf:"ha_vm_dependency_restart_condition,omitempty"`

	// The time interval, in seconds, a heartbeat
	// from a virtual machine is not received within this configured interval,
	// the virtual machine is marked as failed. Default: 30 seconds.
	// If a heartbeat from a virtual machine is not received within this configured interval, the virtual machine is marked as failed. The value is in seconds.
	// +kubebuilder:validation:Optional
	HaVMFailureInterval *float64 `json:"haVmFailureInterval,omitempty" tf:"ha_vm_failure_interval,omitempty"`

	// The time, in seconds, for the reset window in
	// which ha_vm_maximum_resets can operate. When this
	// window expires, no more resets are attempted regardless of the setting
	// configured in ha_vm_maximum_resets. -1 means no window, meaning an
	// unlimited reset time is allotted. Default: -1 (no window).
	// The length of the reset window in which ha_vm_maximum_resets can operate. When this window expires, no more resets are attempted regardless of the setting configured in ha_vm_maximum_resets. -1 means no window, meaning an unlimited reset time is allotted.
	// +kubebuilder:validation:Optional
	HaVMMaximumFailureWindow *float64 `json:"haVmMaximumFailureWindow,omitempty" tf:"ha_vm_maximum_failure_window,omitempty"`

	// The maximum number of resets that HA will
	// perform to a virtual machine when responding to a failure event. Default: 3
	// The maximum number of resets that HA will perform to a virtual machine when responding to a failure event.
	// +kubebuilder:validation:Optional
	HaVMMaximumResets *float64 `json:"haVmMaximumResets,omitempty" tf:"ha_vm_maximum_resets,omitempty"`

	// The time, in seconds, that HA waits after
	// powering on a virtual machine before monitoring for heartbeats. Default:
	// 120 seconds (2 minutes).
	// The time, in seconds, that HA waits after powering on a virtual machine before monitoring for heartbeats.
	// +kubebuilder:validation:Optional
	HaVMMinimumUptime *float64 `json:"haVmMinimumUptime,omitempty" tf:"ha_vm_minimum_uptime,omitempty"`

	// The type of virtual machine monitoring to use
	// when HA is enabled in the cluster. Can be one of vmMonitoringDisabled,
	// vmMonitoringOnly, or vmAndAppMonitoring. Default: vmMonitoringDisabled.
	// The type of virtual machine monitoring to use when HA is enabled in the cluster. Can be one of vmMonitoringDisabled, vmMonitoringOnly, or vmAndAppMonitoring.
	// +kubebuilder:validation:Optional
	HaVMMonitoring *string `json:"haVmMonitoring,omitempty" tf:"ha_vm_monitoring,omitempty"`

	// Additional delay, in seconds,
	// after ready condition is met. A VM is considered ready at this point.
	// Default: 0 seconds (no delay). *
	// Additional delay in seconds after ready condition is met. A VM is considered ready at this point.
	// +kubebuilder:validation:Optional
	HaVMRestartAdditionalDelay *float64 `json:"haVmRestartAdditionalDelay,omitempty" tf:"ha_vm_restart_additional_delay,omitempty"`

	// The default restart priority
	// for affected virtual machines when vSphere detects a host failure. Can be one
	// of lowest, low, medium, high, or highest. Default: medium.
	// The default restart priority for affected VMs when vSphere detects a host failure. Can be one of lowest, low, medium, high, or highest.
	// +kubebuilder:validation:Optional
	HaVMRestartPriority *string `json:"haVmRestartPriority,omitempty" tf:"ha_vm_restart_priority,omitempty"`

	// The maximum time, in seconds,
	// that vSphere HA will wait for virtual machines in one priority to be ready
	// before proceeding with the next priority. Default: 600 seconds (10 minutes).
	// *
	// The maximum time, in seconds, that vSphere HA will wait for virtual machines in one priority to be ready before proceeding with the next priority.
	// +kubebuilder:validation:Optional
	HaVMRestartTimeout *float64 `json:"haVmRestartTimeout,omitempty" tf:"ha_vm_restart_timeout,omitempty"`

	// The timeout, in seconds, for each host maintenance
	// mode operation when removing hosts from a cluster. Default: 3600 seconds (1 hour).
	// The timeout for each host maintenance mode operation when removing hosts from a cluster.
	// +kubebuilder:validation:Optional
	HostClusterExitTimeout *float64 `json:"hostClusterExitTimeout,omitempty" tf:"host_cluster_exit_timeout,omitempty"`

	// Can be set to true if compute cluster
	// membership will be managed through the host resource rather than the
	// compute_cluster resource. Conflicts with: host_system_ids.
	// Must be set if cluster enrollment is managed from host resource.
	// +kubebuilder:validation:Optional
	HostManaged *bool `json:"hostManaged,omitempty" tf:"host_managed,omitempty"`

	// The managed object IDs of
	// the hosts to put in the cluster. Conflicts with: host_managed.
	// The managed object IDs of the hosts to put in the cluster.
	// +kubebuilder:validation:Optional
	HostSystemIds []*string `json:"hostSystemIds,omitempty" tf:"host_system_ids,omitempty"`

	// Determines how the host
	// quarantine, maintenance mode, or virtual machine migration recommendations
	// made by proactive HA are to be handled. Can be one of Automated or
	// Manual. Default: Manual. *
	// The DRS behavior for proactive HA recommendations. Can be one of Automated or Manual.
	// +kubebuilder:validation:Optional
	ProactiveHaAutomationLevel *string `json:"proactiveHaAutomationLevel,omitempty" tf:"proactive_ha_automation_level,omitempty"`

	// Enables Proactive HA. Default: false.
	// *
	// Enables proactive HA, allowing for vSphere to get HA data from external providers and use DRS to perform remediation.
	// +kubebuilder:validation:Optional
	ProactiveHaEnabled *bool `json:"proactiveHaEnabled,omitempty" tf:"proactive_ha_enabled,omitempty"`

	// The configured remediation
	// for moderately degraded hosts. Can be one of MaintenanceMode or
	// QuarantineMode. Note that this cannot be set to MaintenanceMode when
	// proactive_ha_severe_remediation is set
	// to QuarantineMode. Default: QuarantineMode.
	// *
	// The configured remediation for moderately degraded hosts. Can be one of MaintenanceMode or QuarantineMode. Note that this cannot be set to MaintenanceMode when proactive_ha_severe_remediation is set to QuarantineMode.
	// +kubebuilder:validation:Optional
	ProactiveHaModerateRemediation *string `json:"proactiveHaModerateRemediation,omitempty" tf:"proactive_ha_moderate_remediation,omitempty"`

	// The list of IDs for health update
	// providers configured for this cluster.
	// *
	// The list of IDs for health update providers configured for this cluster.
	// +kubebuilder:validation:Optional
	ProactiveHaProviderIds []*string `json:"proactiveHaProviderIds,omitempty" tf:"proactive_ha_provider_ids,omitempty"`

	// The configured remediation for
	// severely degraded hosts. Can be one of MaintenanceMode or QuarantineMode.
	// Note that this cannot be set to QuarantineMode when
	// proactive_ha_moderate_remediation is
	// set to MaintenanceMode. Default: QuarantineMode.
	// *
	// The configured remediation for severely degraded hosts. Can be one of MaintenanceMode or QuarantineMode. Note that this cannot be set to QuarantineMode when proactive_ha_moderate_remediation is set to MaintenanceMode.
	// +kubebuilder:validation:Optional
	ProactiveHaSevereRemediation *string `json:"proactiveHaSevereRemediation,omitempty" tf:"proactive_ha_severe_remediation,omitempty"`

	// The IDs of any tags to attach to this resource. See
	// here for a reference on how to apply tags.
	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Enables vSAN compression on the
	// cluster.
	// Whether the vSAN compression service is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanCompressionEnabled *bool `json:"vsanCompressionEnabled,omitempty" tf:"vsan_compression_enabled,omitempty"`

	// Enables vSAN deduplication on the cluster.
	// Cannot be independently set to true. When vSAN deduplication is enabled, vSAN
	// compression must also be enabled.
	// Whether the vSAN deduplication service is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanDedupEnabled *bool `json:"vsanDedupEnabled,omitempty" tf:"vsan_dedup_enabled,omitempty"`

	// Represents the configuration of a host disk
	// group in the cluster.
	// A list of disk UUIDs to add to the vSAN cluster.
	// +kubebuilder:validation:Optional
	VsanDiskGroup []VsanDiskGroupParameters `json:"vsanDiskGroup,omitempty" tf:"vsan_disk_group,omitempty"`

	// Enables vSAN data-in-transit
	// encryption on the cluster. Conflicts with vsan_remote_datastore_ids, i.e.,
	// vSAN data-in-transit feature cannot be enabled with the vSAN HCI Mesh feature
	// at the same time.
	// Whether the vSAN data-in-transit encryption is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanDitEncryptionEnabled *bool `json:"vsanDitEncryptionEnabled,omitempty" tf:"vsan_dit_encryption_enabled,omitempty"`

	// Indicates the rekey interval in
	// minutes for data-in-transit encryption. The valid rekey interval is 30 to
	// 10800 (feature defaults to 1440). Conflicts with vsan_remote_datastore_ids.
	// When vsan_dit_encryption_enabled is enabled, sets the rekey interval of data-in-transit encryption (in minutes).
	// +kubebuilder:validation:Optional
	VsanDitRekeyInterval *float64 `json:"vsanDitRekeyInterval,omitempty" tf:"vsan_dit_rekey_interval,omitempty"`

	// Enables vSAN on the cluster.
	// Whether the vSAN service is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanEnabled *bool `json:"vsanEnabled,omitempty" tf:"vsan_enabled,omitempty"`

	// Enables network
	// diagnostic mode for vSAN performance service on the cluster.
	// Whether the vSAN network diagnostic mode is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanNetworkDiagnosticModeEnabled *bool `json:"vsanNetworkDiagnosticModeEnabled,omitempty" tf:"vsan_network_diagnostic_mode_enabled,omitempty"`

	// Enables vSAN performance service on
	// the cluster. Default: true.
	// Whether the vSAN performance service is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanPerformanceEnabled *bool `json:"vsanPerformanceEnabled,omitempty" tf:"vsan_performance_enabled,omitempty"`

	// The remote vSAN datastore IDs to be
	// mounted to this cluster. Conflicts with vsan_dit_encryption_enabled and
	// vsan_dit_rekey_interval, i.e., vSAN HCI Mesh feature cannot be enabled with
	// data-in-transit encryption feature at the same time.
	// The managed object IDs of the vSAN datastore to be mounted on the cluster.
	// +kubebuilder:validation:Optional
	VsanRemoteDatastoreIds []*string `json:"vsanRemoteDatastoreIds,omitempty" tf:"vsan_remote_datastore_ids,omitempty"`

	// Enables vSAN unmap on the cluster.
	// Whether the vSAN unmap service is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanUnmapEnabled *bool `json:"vsanUnmapEnabled,omitempty" tf:"vsan_unmap_enabled,omitempty"`

	// Enables verbose mode for vSAN
	// performance service on the cluster.
	// Whether the vSAN verbose mode is enabled for the cluster.
	// +kubebuilder:validation:Optional
	VsanVerboseModeEnabled *bool `json:"vsanVerboseModeEnabled,omitempty" tf:"vsan_verbose_mode_enabled,omitempty"`
}

type VsanDiskGroupObservation struct {
}

type VsanDiskGroupParameters struct {

	// The canonical name of the disk to use for vSAN cache.
	// Cache disk.
	// +kubebuilder:validation:Optional
	Cache *string `json:"cache,omitempty" tf:"cache,omitempty"`

	// An array of disk canonical names for vSAN storage.
	// List of storage disks.
	// +kubebuilder:validation:Optional
	Storage []*string `json:"storage,omitempty" tf:"storage,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Provides a vSphere cluster resource. This can be used to create and manage clusters of hosts.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
