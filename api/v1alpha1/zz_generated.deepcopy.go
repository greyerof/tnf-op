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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRun) DeepCopyInto(out *CnfCertificationSuiteRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRun.
func (in *CnfCertificationSuiteRun) DeepCopy() *CnfCertificationSuiteRun {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnfCertificationSuiteRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunList) DeepCopyInto(out *CnfCertificationSuiteRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnfCertificationSuiteRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunList.
func (in *CnfCertificationSuiteRunList) DeepCopy() *CnfCertificationSuiteRunList {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnfCertificationSuiteRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunSpec) DeepCopyInto(out *CnfCertificationSuiteRunSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunSpec.
func (in *CnfCertificationSuiteRunSpec) DeepCopy() *CnfCertificationSuiteRunSpec {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnfCertificationSuiteRunStatus) DeepCopyInto(out *CnfCertificationSuiteRunStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnfCertificationSuiteRunStatus.
func (in *CnfCertificationSuiteRunStatus) DeepCopy() *CnfCertificationSuiteRunStatus {
	if in == nil {
		return nil
	}
	out := new(CnfCertificationSuiteRunStatus)
	in.DeepCopyInto(out)
	return out
}
