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
func (in *HostObservation) DeepCopyInto(out *HostObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostObservation.
func (in *HostObservation) DeepCopy() *HostObservation {
	if in == nil {
		return nil
	}
	out := new(HostObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostParameters) DeepCopyInto(out *HostParameters) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HostSystemID != nil {
		in, out := &in.HostSystemID, &out.HostSystemID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostParameters.
func (in *HostParameters) DeepCopy() *HostParameters {
	if in == nil {
		return nil
	}
	out := new(HostParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PvlanMappingObservation) DeepCopyInto(out *PvlanMappingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PvlanMappingObservation.
func (in *PvlanMappingObservation) DeepCopy() *PvlanMappingObservation {
	if in == nil {
		return nil
	}
	out := new(PvlanMappingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PvlanMappingParameters) DeepCopyInto(out *PvlanMappingParameters) {
	*out = *in
	if in.PrimaryVlanID != nil {
		in, out := &in.PrimaryVlanID, &out.PrimaryVlanID
		*out = new(float64)
		**out = **in
	}
	if in.PvlanType != nil {
		in, out := &in.PvlanType, &out.PvlanType
		*out = new(string)
		**out = **in
	}
	if in.SecondaryVlanID != nil {
		in, out := &in.SecondaryVlanID, &out.SecondaryVlanID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PvlanMappingParameters.
func (in *PvlanMappingParameters) DeepCopy() *PvlanMappingParameters {
	if in == nil {
		return nil
	}
	out := new(PvlanMappingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitch) DeepCopyInto(out *VirtualSwitch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitch.
func (in *VirtualSwitch) DeepCopy() *VirtualSwitch {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualSwitch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchList) DeepCopyInto(out *VirtualSwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualSwitch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchList.
func (in *VirtualSwitchList) DeepCopy() *VirtualSwitchList {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualSwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchObservation) DeepCopyInto(out *VirtualSwitchObservation) {
	*out = *in
	if in.ConfigVersion != nil {
		in, out := &in.ConfigVersion, &out.ConfigVersion
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchObservation.
func (in *VirtualSwitchObservation) DeepCopy() *VirtualSwitchObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchParameters) DeepCopyInto(out *VirtualSwitchParameters) {
	*out = *in
	if in.ActiveUplinks != nil {
		in, out := &in.ActiveUplinks, &out.ActiveUplinks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowForgedTransmits != nil {
		in, out := &in.AllowForgedTransmits, &out.AllowForgedTransmits
		*out = new(bool)
		**out = **in
	}
	if in.AllowMacChanges != nil {
		in, out := &in.AllowMacChanges, &out.AllowMacChanges
		*out = new(bool)
		**out = **in
	}
	if in.AllowPromiscuous != nil {
		in, out := &in.AllowPromiscuous, &out.AllowPromiscuous
		*out = new(bool)
		**out = **in
	}
	if in.BackupnfcMaximumMbit != nil {
		in, out := &in.BackupnfcMaximumMbit, &out.BackupnfcMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.BackupnfcReservationMbit != nil {
		in, out := &in.BackupnfcReservationMbit, &out.BackupnfcReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.BackupnfcShareCount != nil {
		in, out := &in.BackupnfcShareCount, &out.BackupnfcShareCount
		*out = new(float64)
		**out = **in
	}
	if in.BackupnfcShareLevel != nil {
		in, out := &in.BackupnfcShareLevel, &out.BackupnfcShareLevel
		*out = new(string)
		**out = **in
	}
	if in.BlockAllPorts != nil {
		in, out := &in.BlockAllPorts, &out.BlockAllPorts
		*out = new(bool)
		**out = **in
	}
	if in.CheckBeacon != nil {
		in, out := &in.CheckBeacon, &out.CheckBeacon
		*out = new(bool)
		**out = **in
	}
	if in.ContactDetail != nil {
		in, out := &in.ContactDetail, &out.ContactDetail
		*out = new(string)
		**out = **in
	}
	if in.ContactName != nil {
		in, out := &in.ContactName, &out.ContactName
		*out = new(string)
		**out = **in
	}
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DirectpathGen2Allowed != nil {
		in, out := &in.DirectpathGen2Allowed, &out.DirectpathGen2Allowed
		*out = new(bool)
		**out = **in
	}
	if in.EgressShapingAverageBandwidth != nil {
		in, out := &in.EgressShapingAverageBandwidth, &out.EgressShapingAverageBandwidth
		*out = new(float64)
		**out = **in
	}
	if in.EgressShapingBurstSize != nil {
		in, out := &in.EgressShapingBurstSize, &out.EgressShapingBurstSize
		*out = new(float64)
		**out = **in
	}
	if in.EgressShapingEnabled != nil {
		in, out := &in.EgressShapingEnabled, &out.EgressShapingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.EgressShapingPeakBandwidth != nil {
		in, out := &in.EgressShapingPeakBandwidth, &out.EgressShapingPeakBandwidth
		*out = new(float64)
		**out = **in
	}
	if in.Failback != nil {
		in, out := &in.Failback, &out.Failback
		*out = new(bool)
		**out = **in
	}
	if in.FaulttoleranceMaximumMbit != nil {
		in, out := &in.FaulttoleranceMaximumMbit, &out.FaulttoleranceMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.FaulttoleranceReservationMbit != nil {
		in, out := &in.FaulttoleranceReservationMbit, &out.FaulttoleranceReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.FaulttoleranceShareCount != nil {
		in, out := &in.FaulttoleranceShareCount, &out.FaulttoleranceShareCount
		*out = new(float64)
		**out = **in
	}
	if in.FaulttoleranceShareLevel != nil {
		in, out := &in.FaulttoleranceShareLevel, &out.FaulttoleranceShareLevel
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.HbrMaximumMbit != nil {
		in, out := &in.HbrMaximumMbit, &out.HbrMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.HbrReservationMbit != nil {
		in, out := &in.HbrReservationMbit, &out.HbrReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.HbrShareCount != nil {
		in, out := &in.HbrShareCount, &out.HbrShareCount
		*out = new(float64)
		**out = **in
	}
	if in.HbrShareLevel != nil {
		in, out := &in.HbrShareLevel, &out.HbrShareLevel
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = make([]HostParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPv4Address != nil {
		in, out := &in.IPv4Address, &out.IPv4Address
		*out = new(string)
		**out = **in
	}
	if in.ISCSIMaximumMbit != nil {
		in, out := &in.ISCSIMaximumMbit, &out.ISCSIMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.ISCSIReservationMbit != nil {
		in, out := &in.ISCSIReservationMbit, &out.ISCSIReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.ISCSIShareCount != nil {
		in, out := &in.ISCSIShareCount, &out.ISCSIShareCount
		*out = new(float64)
		**out = **in
	}
	if in.ISCSIShareLevel != nil {
		in, out := &in.ISCSIShareLevel, &out.ISCSIShareLevel
		*out = new(string)
		**out = **in
	}
	if in.IgnoreOtherPvlanMappings != nil {
		in, out := &in.IgnoreOtherPvlanMappings, &out.IgnoreOtherPvlanMappings
		*out = new(bool)
		**out = **in
	}
	if in.IngressShapingAverageBandwidth != nil {
		in, out := &in.IngressShapingAverageBandwidth, &out.IngressShapingAverageBandwidth
		*out = new(float64)
		**out = **in
	}
	if in.IngressShapingBurstSize != nil {
		in, out := &in.IngressShapingBurstSize, &out.IngressShapingBurstSize
		*out = new(float64)
		**out = **in
	}
	if in.IngressShapingEnabled != nil {
		in, out := &in.IngressShapingEnabled, &out.IngressShapingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IngressShapingPeakBandwidth != nil {
		in, out := &in.IngressShapingPeakBandwidth, &out.IngressShapingPeakBandwidth
		*out = new(float64)
		**out = **in
	}
	if in.LacpAPIVersion != nil {
		in, out := &in.LacpAPIVersion, &out.LacpAPIVersion
		*out = new(string)
		**out = **in
	}
	if in.LacpEnabled != nil {
		in, out := &in.LacpEnabled, &out.LacpEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LacpMode != nil {
		in, out := &in.LacpMode, &out.LacpMode
		*out = new(string)
		**out = **in
	}
	if in.LinkDiscoveryOperation != nil {
		in, out := &in.LinkDiscoveryOperation, &out.LinkDiscoveryOperation
		*out = new(string)
		**out = **in
	}
	if in.LinkDiscoveryProtocol != nil {
		in, out := &in.LinkDiscoveryProtocol, &out.LinkDiscoveryProtocol
		*out = new(string)
		**out = **in
	}
	if in.ManagementMaximumMbit != nil {
		in, out := &in.ManagementMaximumMbit, &out.ManagementMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.ManagementReservationMbit != nil {
		in, out := &in.ManagementReservationMbit, &out.ManagementReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.ManagementShareCount != nil {
		in, out := &in.ManagementShareCount, &out.ManagementShareCount
		*out = new(float64)
		**out = **in
	}
	if in.ManagementShareLevel != nil {
		in, out := &in.ManagementShareLevel, &out.ManagementShareLevel
		*out = new(string)
		**out = **in
	}
	if in.MaxMtu != nil {
		in, out := &in.MaxMtu, &out.MaxMtu
		*out = new(float64)
		**out = **in
	}
	if in.MulticastFilteringMode != nil {
		in, out := &in.MulticastFilteringMode, &out.MulticastFilteringMode
		*out = new(string)
		**out = **in
	}
	if in.NFSMaximumMbit != nil {
		in, out := &in.NFSMaximumMbit, &out.NFSMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.NFSReservationMbit != nil {
		in, out := &in.NFSReservationMbit, &out.NFSReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.NFSShareCount != nil {
		in, out := &in.NFSShareCount, &out.NFSShareCount
		*out = new(float64)
		**out = **in
	}
	if in.NFSShareLevel != nil {
		in, out := &in.NFSShareLevel, &out.NFSShareLevel
		*out = new(string)
		**out = **in
	}
	if in.NetflowActiveFlowTimeout != nil {
		in, out := &in.NetflowActiveFlowTimeout, &out.NetflowActiveFlowTimeout
		*out = new(float64)
		**out = **in
	}
	if in.NetflowCollectorIPAddress != nil {
		in, out := &in.NetflowCollectorIPAddress, &out.NetflowCollectorIPAddress
		*out = new(string)
		**out = **in
	}
	if in.NetflowCollectorPort != nil {
		in, out := &in.NetflowCollectorPort, &out.NetflowCollectorPort
		*out = new(float64)
		**out = **in
	}
	if in.NetflowEnabled != nil {
		in, out := &in.NetflowEnabled, &out.NetflowEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NetflowIdleFlowTimeout != nil {
		in, out := &in.NetflowIdleFlowTimeout, &out.NetflowIdleFlowTimeout
		*out = new(float64)
		**out = **in
	}
	if in.NetflowInternalFlowsOnly != nil {
		in, out := &in.NetflowInternalFlowsOnly, &out.NetflowInternalFlowsOnly
		*out = new(bool)
		**out = **in
	}
	if in.NetflowObservationDomainID != nil {
		in, out := &in.NetflowObservationDomainID, &out.NetflowObservationDomainID
		*out = new(float64)
		**out = **in
	}
	if in.NetflowSamplingRate != nil {
		in, out := &in.NetflowSamplingRate, &out.NetflowSamplingRate
		*out = new(float64)
		**out = **in
	}
	if in.NetworkResourceControlEnabled != nil {
		in, out := &in.NetworkResourceControlEnabled, &out.NetworkResourceControlEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NetworkResourceControlVersion != nil {
		in, out := &in.NetworkResourceControlVersion, &out.NetworkResourceControlVersion
		*out = new(string)
		**out = **in
	}
	if in.NotifySwitches != nil {
		in, out := &in.NotifySwitches, &out.NotifySwitches
		*out = new(bool)
		**out = **in
	}
	if in.PortPrivateSecondaryVlanID != nil {
		in, out := &in.PortPrivateSecondaryVlanID, &out.PortPrivateSecondaryVlanID
		*out = new(float64)
		**out = **in
	}
	if in.PvlanMapping != nil {
		in, out := &in.PvlanMapping, &out.PvlanMapping
		*out = make([]PvlanMappingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StandbyUplinks != nil {
		in, out := &in.StandbyUplinks, &out.StandbyUplinks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.TeamingPolicy != nil {
		in, out := &in.TeamingPolicy, &out.TeamingPolicy
		*out = new(string)
		**out = **in
	}
	if in.TxUplink != nil {
		in, out := &in.TxUplink, &out.TxUplink
		*out = new(bool)
		**out = **in
	}
	if in.Uplinks != nil {
		in, out := &in.Uplinks, &out.Uplinks
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VdpMaximumMbit != nil {
		in, out := &in.VdpMaximumMbit, &out.VdpMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.VdpReservationMbit != nil {
		in, out := &in.VdpReservationMbit, &out.VdpReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.VdpShareCount != nil {
		in, out := &in.VdpShareCount, &out.VdpShareCount
		*out = new(float64)
		**out = **in
	}
	if in.VdpShareLevel != nil {
		in, out := &in.VdpShareLevel, &out.VdpShareLevel
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VirtualmachineMaximumMbit != nil {
		in, out := &in.VirtualmachineMaximumMbit, &out.VirtualmachineMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.VirtualmachineReservationMbit != nil {
		in, out := &in.VirtualmachineReservationMbit, &out.VirtualmachineReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.VirtualmachineShareCount != nil {
		in, out := &in.VirtualmachineShareCount, &out.VirtualmachineShareCount
		*out = new(float64)
		**out = **in
	}
	if in.VirtualmachineShareLevel != nil {
		in, out := &in.VirtualmachineShareLevel, &out.VirtualmachineShareLevel
		*out = new(string)
		**out = **in
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(float64)
		**out = **in
	}
	if in.VlanRange != nil {
		in, out := &in.VlanRange, &out.VlanRange
		*out = make([]VlanRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VmotionMaximumMbit != nil {
		in, out := &in.VmotionMaximumMbit, &out.VmotionMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.VmotionReservationMbit != nil {
		in, out := &in.VmotionReservationMbit, &out.VmotionReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.VmotionShareCount != nil {
		in, out := &in.VmotionShareCount, &out.VmotionShareCount
		*out = new(float64)
		**out = **in
	}
	if in.VmotionShareLevel != nil {
		in, out := &in.VmotionShareLevel, &out.VmotionShareLevel
		*out = new(string)
		**out = **in
	}
	if in.VsanMaximumMbit != nil {
		in, out := &in.VsanMaximumMbit, &out.VsanMaximumMbit
		*out = new(float64)
		**out = **in
	}
	if in.VsanReservationMbit != nil {
		in, out := &in.VsanReservationMbit, &out.VsanReservationMbit
		*out = new(float64)
		**out = **in
	}
	if in.VsanShareCount != nil {
		in, out := &in.VsanShareCount, &out.VsanShareCount
		*out = new(float64)
		**out = **in
	}
	if in.VsanShareLevel != nil {
		in, out := &in.VsanShareLevel, &out.VsanShareLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchParameters.
func (in *VirtualSwitchParameters) DeepCopy() *VirtualSwitchParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchSpec) DeepCopyInto(out *VirtualSwitchSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchSpec.
func (in *VirtualSwitchSpec) DeepCopy() *VirtualSwitchSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualSwitchStatus) DeepCopyInto(out *VirtualSwitchStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualSwitchStatus.
func (in *VirtualSwitchStatus) DeepCopy() *VirtualSwitchStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualSwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanRangeObservation) DeepCopyInto(out *VlanRangeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanRangeObservation.
func (in *VlanRangeObservation) DeepCopy() *VlanRangeObservation {
	if in == nil {
		return nil
	}
	out := new(VlanRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanRangeParameters) DeepCopyInto(out *VlanRangeParameters) {
	*out = *in
	if in.MaxVlan != nil {
		in, out := &in.MaxVlan, &out.MaxVlan
		*out = new(float64)
		**out = **in
	}
	if in.MinVlan != nil {
		in, out := &in.MinVlan, &out.MinVlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanRangeParameters.
func (in *VlanRangeParameters) DeepCopy() *VlanRangeParameters {
	if in == nil {
		return nil
	}
	out := new(VlanRangeParameters)
	in.DeepCopyInto(out)
	return out
}