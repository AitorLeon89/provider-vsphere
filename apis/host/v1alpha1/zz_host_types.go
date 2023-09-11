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

type HostObservation struct {

	// The ID of the host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HostParameters struct {

	// The ID of the Compute Cluster this host should
	// be added to. This should not be set if datacenter is set. Conflicts with:
	// cluster.
	// ID of the vSphere cluster the host will belong to.
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Can be set to true if compute cluster
	// membership will be managed through the compute_cluster resource rather
	// than thehost resource. Conflicts with: cluster.
	// Must be set if host is a member of a managed compute_cluster resource.
	// +kubebuilder:validation:Optional
	ClusterManaged *bool `json:"clusterManaged,omitempty" tf:"cluster_managed,omitempty"`

	// If set to false then the host will be disconnected.
	// Default is false.
	// Set the state of the host. If set to false then the host will be asked to disconnect.
	// +kubebuilder:validation:Optional
	Connected *bool `json:"connected,omitempty" tf:"connected,omitempty"`

	// A map of custom attribute IDs and string
	// values to apply to the resource. Please refer to the
	// vsphere_custom_attributes resource for more information on applying
	// tags to resources.
	// A list of custom attributes to set on this resource.
	// +kubebuilder:validation:Optional
	CustomAttributes map[string]*string `json:"customAttributes,omitempty" tf:"custom_attributes,omitempty"`

	// The ID of the datacenter this host should
	// be added to. This should not be set if cluster is set.
	// ID of the vSphere datacenter the host will belong to.
	// +crossplane:generate:reference:type=github.com/AitorLeon89/provider-vsphere/apis/datacenter/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// Reference to a Datacenter in datacenter to populate datacenter.
	// +kubebuilder:validation:Optional
	DatacenterRef *v1.Reference `json:"datacenterRef,omitempty" tf:"-"`

	// Selector for a Datacenter in datacenter to populate datacenter.
	// +kubebuilder:validation:Optional
	DatacenterSelector *v1.Selector `json:"datacenterSelector,omitempty" tf:"-"`

	// If set to true then it will force the host to be added,
	// even if the host is already connected to a different vCenter Server instance.
	// Default is false.
	// Force add the host to the vSphere inventory even if it's already managed by a different vCenter Server instance.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// The license key that will be applied to the host.
	// The license key is expected to be present in vSphere.
	// License key that will be applied to this host.
	// +kubebuilder:validation:Optional
	License *string `json:"license,omitempty" tf:"license,omitempty"`

	// Set the lockdown state of the host. Valid options are
	// disabled, normal, and strict. Default is disabled.
	// Set the host's lockdown status. Default is disabled. Valid options are 'disabled', 'normal', 'strict'
	// +kubebuilder:validation:Optional
	Lockdown *string `json:"lockdown,omitempty" tf:"lockdown,omitempty"`

	// Set the management state of the host.
	// Default is false.
	// Set the host's maintenance mode. Default is false
	// +kubebuilder:validation:Optional
	Maintenance *bool `json:"maintenance,omitempty" tf:"maintenance,omitempty"`

	// Password that will be used by vSphere to authenticate
	// to the host.
	// Password of the administration account of the host.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The IDs of any tags to attach to this resource. Please
	// refer to the vsphere_tag resource for more information on applying
	// tags to resources.
	// A list of tag IDs to apply to this object.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Host's certificate SHA-1 thumbprint. If not set the
	// CA that signed the host's certificate should be trusted. If the CA is not
	// trusted and no thumbprint is set then the operation will fail. See data source
	// vsphere_host_thumbprint.
	// Host's certificate SHA-1 thumbprint. If not set then the CA that signed the host's certificate must be trusted.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	// Username that will be used by vSphere to authenticate
	// to the host.
	// Username of the administration account of the host.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostParameters `json:"forProvider"`
}

// HostStatus defines the observed state of Host.
type HostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Host is the Schema for the Hosts API. Provides a VMware vSphere host resource. This represents an ESXi host that can be used as a member of a cluster or as a standalone host.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec"`
	Status            HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostList contains a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// Repository type metadata.
var (
	Host_Kind             = "Host"
	Host_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Host_Kind}.String()
	Host_KindAPIVersion   = Host_Kind + "." + CRDGroupVersion.String()
	Host_GroupVersionKind = CRDGroupVersion.WithKind(Host_Kind)
)

func init() {
	SchemeBuilder.Register(&Host{}, &HostList{})
}