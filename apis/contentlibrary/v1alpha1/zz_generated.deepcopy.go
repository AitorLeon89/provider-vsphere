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
func (in *Library) DeepCopyInto(out *Library) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Library.
func (in *Library) DeepCopy() *Library {
	if in == nil {
		return nil
	}
	out := new(Library)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Library) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryList) DeepCopyInto(out *LibraryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Library, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryList.
func (in *LibraryList) DeepCopy() *LibraryList {
	if in == nil {
		return nil
	}
	out := new(LibraryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibraryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryObservation) DeepCopyInto(out *LibraryObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Publication != nil {
		in, out := &in.Publication, &out.Publication
		*out = make([]PublicationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryObservation.
func (in *LibraryObservation) DeepCopy() *LibraryObservation {
	if in == nil {
		return nil
	}
	out := new(LibraryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryParameters) DeepCopyInto(out *LibraryParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Publication != nil {
		in, out := &in.Publication, &out.Publication
		*out = make([]PublicationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageBacking != nil {
		in, out := &in.StorageBacking, &out.StorageBacking
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Subscription != nil {
		in, out := &in.Subscription, &out.Subscription
		*out = make([]SubscriptionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryParameters.
func (in *LibraryParameters) DeepCopy() *LibraryParameters {
	if in == nil {
		return nil
	}
	out := new(LibraryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibrarySpec) DeepCopyInto(out *LibrarySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibrarySpec.
func (in *LibrarySpec) DeepCopy() *LibrarySpec {
	if in == nil {
		return nil
	}
	out := new(LibrarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibraryStatus) DeepCopyInto(out *LibraryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibraryStatus.
func (in *LibraryStatus) DeepCopy() *LibraryStatus {
	if in == nil {
		return nil
	}
	out := new(LibraryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationObservation) DeepCopyInto(out *PublicationObservation) {
	*out = *in
	if in.PublishURL != nil {
		in, out := &in.PublishURL, &out.PublishURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationObservation.
func (in *PublicationObservation) DeepCopy() *PublicationObservation {
	if in == nil {
		return nil
	}
	out := new(PublicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationParameters) DeepCopyInto(out *PublicationParameters) {
	*out = *in
	if in.AuthenticationMethod != nil {
		in, out := &in.AuthenticationMethod, &out.AuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Published != nil {
		in, out := &in.Published, &out.Published
		*out = new(bool)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationParameters.
func (in *PublicationParameters) DeepCopy() *PublicationParameters {
	if in == nil {
		return nil
	}
	out := new(PublicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionObservation) DeepCopyInto(out *SubscriptionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionObservation.
func (in *SubscriptionObservation) DeepCopy() *SubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionParameters) DeepCopyInto(out *SubscriptionParameters) {
	*out = *in
	if in.AuthenticationMethod != nil {
		in, out := &in.AuthenticationMethod, &out.AuthenticationMethod
		*out = new(string)
		**out = **in
	}
	if in.AutomaticSync != nil {
		in, out := &in.AutomaticSync, &out.AutomaticSync
		*out = new(bool)
		**out = **in
	}
	if in.OnDemand != nil {
		in, out := &in.OnDemand, &out.OnDemand
		*out = new(bool)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionURL != nil {
		in, out := &in.SubscriptionURL, &out.SubscriptionURL
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionParameters.
func (in *SubscriptionParameters) DeepCopy() *SubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}