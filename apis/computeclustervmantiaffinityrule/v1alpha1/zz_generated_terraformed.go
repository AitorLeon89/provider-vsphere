/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ClusterVMAntiAffinityRule
func (mg *ClusterVMAntiAffinityRule) GetTerraformResourceType() string {
	return "vsphere_compute_cluster_vm_anti_affinity_rule"
}

// GetConnectionDetailsMapping for this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ClusterVMAntiAffinityRule
func (tr *ClusterVMAntiAffinityRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ClusterVMAntiAffinityRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ClusterVMAntiAffinityRule) LateInitialize(attrs []byte) (bool, error) {
	params := &ClusterVMAntiAffinityRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ClusterVMAntiAffinityRule) GetTerraformSchemaVersion() int {
	return 0
}
