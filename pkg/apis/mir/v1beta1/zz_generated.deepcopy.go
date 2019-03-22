// +build !ignore_autogenerated

/*
Copyright 2019 Microsoft Corporation.

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelInferenceResource) DeepCopyInto(out *ModelInferenceResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelInferenceResource.
func (in *ModelInferenceResource) DeepCopy() *ModelInferenceResource {
	if in == nil {
		return nil
	}
	out := new(ModelInferenceResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelInferenceResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelInferenceResourceList) DeepCopyInto(out *ModelInferenceResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelInferenceResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelInferenceResourceList.
func (in *ModelInferenceResourceList) DeepCopy() *ModelInferenceResourceList {
	if in == nil {
		return nil
	}
	out := new(ModelInferenceResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelInferenceResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelInferenceResourceSpec) DeepCopyInto(out *ModelInferenceResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelInferenceResourceSpec.
func (in *ModelInferenceResourceSpec) DeepCopy() *ModelInferenceResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ModelInferenceResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelInferenceResourceStatus) DeepCopyInto(out *ModelInferenceResourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelInferenceResourceStatus.
func (in *ModelInferenceResourceStatus) DeepCopy() *ModelInferenceResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ModelInferenceResourceStatus)
	in.DeepCopyInto(out)
	return out
}
