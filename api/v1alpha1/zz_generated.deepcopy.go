//go:build !ignore_autogenerated

/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftBuild) DeepCopyInto(out *OpenShiftBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftBuild.
func (in *OpenShiftBuild) DeepCopy() *OpenShiftBuild {
	if in == nil {
		return nil
	}
	out := new(OpenShiftBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftBuildList) DeepCopyInto(out *OpenShiftBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenShiftBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftBuildList.
func (in *OpenShiftBuildList) DeepCopy() *OpenShiftBuildList {
	if in == nil {
		return nil
	}
	out := new(OpenShiftBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftBuildSpec) DeepCopyInto(out *OpenShiftBuildSpec) {
	*out = *in
	if in.Shipwright != nil {
		in, out := &in.Shipwright, &out.Shipwright
		*out = new(Shipwright)
		(*in).DeepCopyInto(*out)
	}
	if in.SharedResource != nil {
		in, out := &in.SharedResource, &out.SharedResource
		*out = new(SharedResource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftBuildSpec.
func (in *OpenShiftBuildSpec) DeepCopy() *OpenShiftBuildSpec {
	if in == nil {
		return nil
	}
	out := new(OpenShiftBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftBuildStatus) DeepCopyInto(out *OpenShiftBuildStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftBuildStatus.
func (in *OpenShiftBuildStatus) DeepCopy() *OpenShiftBuildStatus {
	if in == nil {
		return nil
	}
	out := new(OpenShiftBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedResource) DeepCopyInto(out *SharedResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedResource.
func (in *SharedResource) DeepCopy() *SharedResource {
	if in == nil {
		return nil
	}
	out := new(SharedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Shipwright) DeepCopyInto(out *Shipwright) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(ShipwrightBuild)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Shipwright.
func (in *Shipwright) DeepCopy() *Shipwright {
	if in == nil {
		return nil
	}
	out := new(Shipwright)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipwrightBuild) DeepCopyInto(out *ShipwrightBuild) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipwrightBuild.
func (in *ShipwrightBuild) DeepCopy() *ShipwrightBuild {
	if in == nil {
		return nil
	}
	out := new(ShipwrightBuild)
	in.DeepCopyInto(out)
	return out
}
