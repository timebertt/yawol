//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerDebugSettings) DeepCopyInto(out *LoadBalancerDebugSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerDebugSettings.
func (in *LoadBalancerDebugSettings) DeepCopy() *LoadBalancerDebugSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerDebugSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerEndpoint) DeepCopyInto(out *LoadBalancerEndpoint) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerEndpoint.
func (in *LoadBalancerEndpoint) DeepCopy() *LoadBalancerEndpoint {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerInfrastructure) DeepCopyInto(out *LoadBalancerInfrastructure) {
	*out = *in
	if in.FloatingNetID != nil {
		in, out := &in.FloatingNetID, &out.FloatingNetID
		*out = new(string)
		**out = **in
	}
	if in.Flavor != nil {
		in, out := &in.Flavor, &out.Flavor
		*out = new(OpenstackFlavorRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(OpenstackImageRef)
		(*in).DeepCopyInto(*out)
	}
	out.AuthSecretRef = in.AuthSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerInfrastructure.
func (in *LoadBalancerInfrastructure) DeepCopy() *LoadBalancerInfrastructure {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerInfrastructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerList.
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerLogForward) DeepCopyInto(out *LoadBalancerLogForward) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerLogForward.
func (in *LoadBalancerLogForward) DeepCopy() *LoadBalancerLogForward {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerLogForward)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachine) DeepCopyInto(out *LoadBalancerMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachine.
func (in *LoadBalancerMachine) DeepCopy() *LoadBalancerMachine {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineList) DeepCopyInto(out *LoadBalancerMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineList.
func (in *LoadBalancerMachineList) DeepCopy() *LoadBalancerMachineList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineMetric) DeepCopyInto(out *LoadBalancerMachineMetric) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineMetric.
func (in *LoadBalancerMachineMetric) DeepCopy() *LoadBalancerMachineMetric {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineSpec) DeepCopyInto(out *LoadBalancerMachineSpec) {
	*out = *in
	in.Infrastructure.DeepCopyInto(&out.Infrastructure)
	out.LoadBalancerRef = in.LoadBalancerRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineSpec.
func (in *LoadBalancerMachineSpec) DeepCopy() *LoadBalancerMachineSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineStatus) DeepCopyInto(out *LoadBalancerMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new([]v1.NodeCondition)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.NodeCondition, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new([]LoadBalancerMachineMetric)
		if **in != nil {
			in, out := *in, *out
			*out = make([]LoadBalancerMachineMetric, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.LastOpenstackReconcile != nil {
		in, out := &in.LastOpenstackReconcile, &out.LastOpenstackReconcile
		*out = (*in).DeepCopy()
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.RoleName != nil {
		in, out := &in.RoleName, &out.RoleName
		*out = new(string)
		**out = **in
	}
	if in.RoleBindingName != nil {
		in, out := &in.RoleBindingName, &out.RoleBindingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineStatus.
func (in *LoadBalancerMachineStatus) DeepCopy() *LoadBalancerMachineStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerMachineTemplateSpec) DeepCopyInto(out *LoadBalancerMachineTemplateSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerMachineTemplateSpec.
func (in *LoadBalancerMachineTemplateSpec) DeepCopy() *LoadBalancerMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerOptions) DeepCopyInto(out *LoadBalancerOptions) {
	*out = *in
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TCPProxyProtocolPortsFilter != nil {
		in, out := &in.TCPProxyProtocolPortsFilter, &out.TCPProxyProtocolPortsFilter
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.TCPIdleTimeout != nil {
		in, out := &in.TCPIdleTimeout, &out.TCPIdleTimeout
		*out = new(int)
		**out = **in
	}
	if in.UDPIdleTimeout != nil {
		in, out := &in.UDPIdleTimeout, &out.UDPIdleTimeout
		*out = new(int)
		**out = **in
	}
	out.LogForward = in.LogForward
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerOptions.
func (in *LoadBalancerOptions) DeepCopy() *LoadBalancerOptions {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerRef) DeepCopyInto(out *LoadBalancerRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerRef.
func (in *LoadBalancerRef) DeepCopy() *LoadBalancerRef {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSet) DeepCopyInto(out *LoadBalancerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSet.
func (in *LoadBalancerSet) DeepCopy() *LoadBalancerSet {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetList) DeepCopyInto(out *LoadBalancerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetList.
func (in *LoadBalancerSetList) DeepCopy() *LoadBalancerSetList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetSpec) DeepCopyInto(out *LoadBalancerSetSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetSpec.
func (in *LoadBalancerSetSpec) DeepCopy() *LoadBalancerSetSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSetStatus) DeepCopyInto(out *LoadBalancerSetStatus) {
	*out = *in
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSetStatus.
func (in *LoadBalancerSetStatus) DeepCopy() *LoadBalancerSetStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.ExistingFloatingIP != nil {
		in, out := &in.ExistingFloatingIP, &out.ExistingFloatingIP
		*out = new(string)
		**out = **in
	}
	out.DebugSettings = in.DebugSettings
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]LoadBalancerEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Infrastructure.DeepCopyInto(&out.Infrastructure)
	in.Options.DeepCopyInto(&out.Options)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = new(string)
		**out = **in
	}
	if in.FloatingID != nil {
		in, out := &in.FloatingID, &out.FloatingID
		*out = new(string)
		**out = **in
	}
	if in.FloatingName != nil {
		in, out := &in.FloatingName, &out.FloatingName
		*out = new(string)
		**out = **in
	}
	if in.PortID != nil {
		in, out := &in.PortID, &out.PortID
		*out = new(string)
		**out = **in
	}
	if in.PortName != nil {
		in, out := &in.PortName, &out.PortName
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupName != nil {
		in, out := &in.SecurityGroupName, &out.SecurityGroupName
		*out = new(string)
		**out = **in
	}
	if in.LastOpenstackReconcile != nil {
		in, out := &in.LastOpenstackReconcile, &out.LastOpenstackReconcile
		*out = (*in).DeepCopy()
	}
	if in.OpenstackReconcileHash != nil {
		in, out := &in.OpenstackReconcileHash, &out.OpenstackReconcileHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackFlavorRef) DeepCopyInto(out *OpenstackFlavorRef) {
	*out = *in
	if in.FlavorID != nil {
		in, out := &in.FlavorID, &out.FlavorID
		*out = new(string)
		**out = **in
	}
	if in.FlavorName != nil {
		in, out := &in.FlavorName, &out.FlavorName
		*out = new(string)
		**out = **in
	}
	if in.FlavorSearch != nil {
		in, out := &in.FlavorSearch, &out.FlavorSearch
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackFlavorRef.
func (in *OpenstackFlavorRef) DeepCopy() *OpenstackFlavorRef {
	if in == nil {
		return nil
	}
	out := new(OpenstackFlavorRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackImageRef) DeepCopyInto(out *OpenstackImageRef) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.ImageSearch != nil {
		in, out := &in.ImageSearch, &out.ImageSearch
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackImageRef.
func (in *OpenstackImageRef) DeepCopy() *OpenstackImageRef {
	if in == nil {
		return nil
	}
	out := new(OpenstackImageRef)
	in.DeepCopyInto(out)
	return out
}
