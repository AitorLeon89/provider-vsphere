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

type DiskObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiskParameters struct {

	// The adapter type for this virtual disk. Can be
	// one of ide, lsiLogic, or busLogic.  Default: lsiLogic.
	// +kubebuilder:validation:Optional
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// Tells the resource to create any
	// directories that are a part of the vmdk_path parameter if they are missing.
	// Default: false.
	// +kubebuilder:validation:Optional
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// The name of the datacenter in which to create the
	// disk. Can be omitted when when ESXi or if there is only one datacenter in
	// your infrastructure.
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// The name of the datastore in which to create the
	// disk.
	// +kubebuilder:validation:Required
	Datastore *string `json:"datastore" tf:"datastore,omitempty"`

	// Size of the disk (in GB).
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// The type of disk to create. Can be one of
	// eagerZeroedThick, lazy, or thin. Default: eagerZeroedThick. For
	// information on what each kind of disk provisioning policy means, click
	// here.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters `json:"forProvider"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Disk is the Schema for the Disks API. Provides a vSphere virtual disk resource.  This can be used to create and delete virtual disks.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskSpec   `json:"spec"`
	Status            DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
