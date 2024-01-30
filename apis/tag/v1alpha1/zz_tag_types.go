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

type TagObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TagParameters struct {

	// The unique identifier of the parent category in
	// which this tag will be created. Forces a new resource if changed.
	// The unique identifier of the parent category in which this tag will be created.
	// +kubebuilder:validation:Required
	CategoryID *string `json:"categoryId" tf:"category_id,omitempty"`

	// A description for the tag.
	// The description of the tag.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The display name of the tag. The name must be unique
	// within its category.
	// The display name of the tag. The name must be unique within its category.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tag is the Schema for the Tags API. Provides a vSphere tag resource. This can be used to manage tags in vSphere.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagSpec   `json:"spec"`
	Status            TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}