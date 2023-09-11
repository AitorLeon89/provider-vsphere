//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.SdrsAdvancedOptions != nil {
		in, out := &in.SdrsAdvancedOptions, &out.SdrsAdvancedOptions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.SdrsAutomationLevel != nil {
		in, out := &in.SdrsAutomationLevel, &out.SdrsAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsDefaultIntraVMAffinity != nil {
		in, out := &in.SdrsDefaultIntraVMAffinity, &out.SdrsDefaultIntraVMAffinity
		*out = new(bool)
		**out = **in
	}
	if in.SdrsEnabled != nil {
		in, out := &in.SdrsEnabled, &out.SdrsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SdrsFreeSpaceThreshold != nil {
		in, out := &in.SdrsFreeSpaceThreshold, &out.SdrsFreeSpaceThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsFreeSpaceThresholdMode != nil {
		in, out := &in.SdrsFreeSpaceThresholdMode, &out.SdrsFreeSpaceThresholdMode
		*out = new(string)
		**out = **in
	}
	if in.SdrsFreeSpaceUtilizationDifference != nil {
		in, out := &in.SdrsFreeSpaceUtilizationDifference, &out.SdrsFreeSpaceUtilizationDifference
		*out = new(float64)
		**out = **in
	}
	if in.SdrsIoBalanceAutomationLevel != nil {
		in, out := &in.SdrsIoBalanceAutomationLevel, &out.SdrsIoBalanceAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsIoLatencyThreshold != nil {
		in, out := &in.SdrsIoLatencyThreshold, &out.SdrsIoLatencyThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsIoLoadBalanceEnabled != nil {
		in, out := &in.SdrsIoLoadBalanceEnabled, &out.SdrsIoLoadBalanceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SdrsIoLoadImbalanceThreshold != nil {
		in, out := &in.SdrsIoLoadImbalanceThreshold, &out.SdrsIoLoadImbalanceThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsIoReservableIopsThreshold != nil {
		in, out := &in.SdrsIoReservableIopsThreshold, &out.SdrsIoReservableIopsThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsIoReservablePercentThreshold != nil {
		in, out := &in.SdrsIoReservablePercentThreshold, &out.SdrsIoReservablePercentThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsIoReservableThresholdMode != nil {
		in, out := &in.SdrsIoReservableThresholdMode, &out.SdrsIoReservableThresholdMode
		*out = new(string)
		**out = **in
	}
	if in.SdrsLoadBalanceInterval != nil {
		in, out := &in.SdrsLoadBalanceInterval, &out.SdrsLoadBalanceInterval
		*out = new(float64)
		**out = **in
	}
	if in.SdrsPolicyEnforcementAutomationLevel != nil {
		in, out := &in.SdrsPolicyEnforcementAutomationLevel, &out.SdrsPolicyEnforcementAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsRuleEnforcementAutomationLevel != nil {
		in, out := &in.SdrsRuleEnforcementAutomationLevel, &out.SdrsRuleEnforcementAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsSpaceBalanceAutomationLevel != nil {
		in, out := &in.SdrsSpaceBalanceAutomationLevel, &out.SdrsSpaceBalanceAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.SdrsSpaceUtilizationThreshold != nil {
		in, out := &in.SdrsSpaceUtilizationThreshold, &out.SdrsSpaceUtilizationThreshold
		*out = new(float64)
		**out = **in
	}
	if in.SdrsVMEvacuationAutomationLevel != nil {
		in, out := &in.SdrsVMEvacuationAutomationLevel, &out.SdrsVMEvacuationAutomationLevel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}