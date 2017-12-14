// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppliedClusterResourceQuota).DeepCopyInto(out.(*AppliedClusterResourceQuota))
			return nil
		}, InType: reflect.TypeOf(new(AppliedClusterResourceQuota))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AppliedClusterResourceQuotaList).DeepCopyInto(out.(*AppliedClusterResourceQuotaList))
			return nil
		}, InType: reflect.TypeOf(new(AppliedClusterResourceQuotaList))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterResourceQuota).DeepCopyInto(out.(*ClusterResourceQuota))
			return nil
		}, InType: reflect.TypeOf(new(ClusterResourceQuota))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterResourceQuotaList).DeepCopyInto(out.(*ClusterResourceQuotaList))
			return nil
		}, InType: reflect.TypeOf(new(ClusterResourceQuotaList))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterResourceQuotaSelector).DeepCopyInto(out.(*ClusterResourceQuotaSelector))
			return nil
		}, InType: reflect.TypeOf(new(ClusterResourceQuotaSelector))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterResourceQuotaSpec).DeepCopyInto(out.(*ClusterResourceQuotaSpec))
			return nil
		}, InType: reflect.TypeOf(new(ClusterResourceQuotaSpec))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterResourceQuotaStatus).DeepCopyInto(out.(*ClusterResourceQuotaStatus))
			return nil
		}, InType: reflect.TypeOf(new(ClusterResourceQuotaStatus))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotaStatusByNamespace).DeepCopyInto(out.(*ResourceQuotaStatusByNamespace))
			return nil
		}, InType: reflect.TypeOf(new(ResourceQuotaStatusByNamespace))},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ResourceQuotasStatusByNamespace).DeepCopyInto(out.(*ResourceQuotasStatusByNamespace))
			return nil
		}, InType: reflect.TypeOf(new(ResourceQuotasStatusByNamespace))},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterResourceQuota) DeepCopyInto(out *AppliedClusterResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterResourceQuota.
func (in *AppliedClusterResourceQuota) DeepCopy() *AppliedClusterResourceQuota {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterResourceQuotaList) DeepCopyInto(out *AppliedClusterResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterResourceQuotaList.
func (in *AppliedClusterResourceQuotaList) DeepCopy() *AppliedClusterResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuota) DeepCopyInto(out *ClusterResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuota.
func (in *ClusterResourceQuota) DeepCopy() *ClusterResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaList) DeepCopyInto(out *ClusterResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaList.
func (in *ClusterResourceQuotaList) DeepCopy() *ClusterResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaSelector) DeepCopyInto(out *ClusterResourceQuotaSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AnnotationSelector != nil {
		in, out := &in.AnnotationSelector, &out.AnnotationSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaSelector.
func (in *ClusterResourceQuotaSelector) DeepCopy() *ClusterResourceQuotaSelector {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaSpec) DeepCopyInto(out *ClusterResourceQuotaSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Quota.DeepCopyInto(&out.Quota)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaSpec.
func (in *ClusterResourceQuotaSpec) DeepCopy() *ClusterResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceQuotaStatus) DeepCopyInto(out *ClusterResourceQuotaStatus) {
	*out = *in
	in.Total.DeepCopyInto(&out.Total)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(ResourceQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceQuotaStatus.
func (in *ClusterResourceQuotaStatus) DeepCopy() *ClusterResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaStatusByNamespace) DeepCopyInto(out *ResourceQuotaStatusByNamespace) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaStatusByNamespace.
func (in *ResourceQuotaStatusByNamespace) DeepCopy() *ResourceQuotaStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaStatusByNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotasStatusByNamespace) DeepCopyInto(out *ResourceQuotasStatusByNamespace) {
	{
		in := (*[]ResourceQuotaStatusByNamespace)(unsafe.Pointer(in))
		out := (*[]ResourceQuotaStatusByNamespace)(unsafe.Pointer(out))
		*out = make([]ResourceQuotaStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotasStatusByNamespace.
func (in *ResourceQuotasStatusByNamespace) DeepCopy() *ResourceQuotasStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotasStatusByNamespace)
	in.DeepCopyInto(out)
	return out
}
