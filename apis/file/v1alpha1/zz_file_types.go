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

type FileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FileParameters struct {

	// Create directories in destination_file
	// path parameter on first apply if any are missing for copy operation.
	// +kubebuilder:validation:Optional
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// The name of a datacenter to which the file will be
	// uploaded.
	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// The name of the datastore to which to upload the
	// file.
	// +kubebuilder:validation:Required
	Datastore *string `json:"datastore" tf:"datastore,omitempty"`

	// The path to where the file should be uploaded
	// or copied to on the destination datastore in vSphere.
	// +kubebuilder:validation:Required
	DestinationFile *string `json:"destinationFile" tf:"destination_file,omitempty"`

	// The name of a datacenter from which the file
	// will be copied. Forces a new resource if changed.
	// +kubebuilder:validation:Optional
	SourceDatacenter *string `json:"sourceDatacenter,omitempty" tf:"source_datacenter,omitempty"`

	// The name of the datastore from which file will
	// be copied. Forces a new resource if changed.
	// +kubebuilder:validation:Optional
	SourceDatastore *string `json:"sourceDatastore,omitempty" tf:"source_datastore,omitempty"`

	// Forces a new resource if changed.
	// +kubebuilder:validation:Required
	SourceFile *string `json:"sourceFile" tf:"source_file,omitempty"`
}

// FileSpec defines the desired state of File
type FileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileParameters `json:"forProvider"`
}

// FileStatus defines the observed state of File.
type FileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// File is the Schema for the Files API. Provides a VMware vSphere file resource. This can be used to upload files (e.g. .iso and .
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSpec   `json:"spec"`
	Status            FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of Files
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// Repository type metadata.
var (
	File_Kind             = "File"
	File_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: File_Kind}.String()
	File_KindAPIVersion   = File_Kind + "." + CRDGroupVersion.String()
	File_GroupVersionKind = CRDGroupVersion.WithKind(File_Kind)
)

func init() {
	SchemeBuilder.Register(&File{}, &FileList{})
}
