/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ClusterVMAntiAffinityRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ClusterVMAntiAffinityRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ClusterVMAntiAffinityRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ClusterVMAntiAffinityRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ClusterVMAntiAffinityRule.
func (mg *ClusterVMAntiAffinityRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
