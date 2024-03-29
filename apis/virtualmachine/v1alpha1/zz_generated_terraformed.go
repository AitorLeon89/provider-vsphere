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

// GetTerraformResourceType returns Terraform resource type for this Machine
func (mg *Machine) GetTerraformResourceType() string {
	return "vsphere_virtual_machine"
}

// GetConnectionDetailsMapping for this Machine
func (tr *Machine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"clone[*].customize[*].linux_options[*].script_text": "spec.forProvider.clone[*].customize[*].linuxOptions[*].scriptTextSecretRef", "clone[*].customize[*].windows_options[*].admin_password": "spec.forProvider.clone[*].customize[*].windowsOptions[*].adminPasswordSecretRef", "clone[*].customize[*].windows_options[*].domain_admin_password": "spec.forProvider.clone[*].customize[*].windowsOptions[*].domainAdminPasswordSecretRef", "clone[*].customize[*].windows_options[*].product_key": "spec.forProvider.clone[*].customize[*].windowsOptions[*].productKeySecretRef", "clone[*].customize[*].windows_sysprep_text": "spec.forProvider.clone[*].customize[*].windowsSysprepTextSecretRef"}
}

// GetObservation of this Machine
func (tr *Machine) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Machine
func (tr *Machine) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Machine
func (tr *Machine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Machine
func (tr *Machine) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Machine
func (tr *Machine) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Machine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Machine) LateInitialize(attrs []byte) (bool, error) {
	params := &MachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Machine) GetTerraformSchemaVersion() int {
	return 3
}
