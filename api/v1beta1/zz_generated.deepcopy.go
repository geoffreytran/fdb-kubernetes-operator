// +build !ignore_autogenerated

/*
Copyright 2020 FoundationDB project authors.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupGenerationStatus) DeepCopyInto(out *BackupGenerationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupGenerationStatus.
func (in *BackupGenerationStatus) DeepCopy() *BackupGenerationStatus {
	if in == nil {
		return nil
	}
	out := new(BackupGenerationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGenerationStatus) DeepCopyInto(out *ClusterGenerationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGenerationStatus.
func (in *ClusterGenerationStatus) DeepCopy() *ClusterGenerationStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterGenerationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealth) DeepCopyInto(out *ClusterHealth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealth.
func (in *ClusterHealth) DeepCopy() *ClusterHealth {
	if in == nil {
		return nil
	}
	out := new(ClusterHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionString) DeepCopyInto(out *ConnectionString) {
	*out = *in
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionString.
func (in *ConnectionString) DeepCopy() *ConnectionString {
	if in == nil {
		return nil
	}
	out := new(ConnectionString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerOverrides) DeepCopyInto(out *ContainerOverrides) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerOverrides.
func (in *ContainerOverrides) DeepCopy() *ContainerOverrides {
	if in == nil {
		return nil
	}
	out := new(ContainerOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCenter) DeepCopyInto(out *DataCenter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCenter.
func (in *DataCenter) DeepCopy() *DataCenter {
	if in == nil {
		return nil
	}
	out := new(DataCenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConfiguration) DeepCopyInto(out *DatabaseConfiguration) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]Region, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.RoleCounts = in.RoleCounts
	out.VersionFlags = in.VersionFlags
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConfiguration.
func (in *DatabaseConfiguration) DeepCopy() *DatabaseConfiguration {
	if in == nil {
		return nil
	}
	out := new(DatabaseConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FdbVersion) DeepCopyInto(out *FdbVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FdbVersion.
func (in *FdbVersion) DeepCopy() *FdbVersion {
	if in == nil {
		return nil
	}
	out := new(FdbVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBBackup) DeepCopyInto(out *FoundationDBBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBBackup.
func (in *FoundationDBBackup) DeepCopy() *FoundationDBBackup {
	if in == nil {
		return nil
	}
	out := new(FoundationDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBBackupList) DeepCopyInto(out *FoundationDBBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FoundationDBBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBBackupList.
func (in *FoundationDBBackupList) DeepCopy() *FoundationDBBackupList {
	if in == nil {
		return nil
	}
	out := new(FoundationDBBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBBackupSpec) DeepCopyInto(out *FoundationDBBackupSpec) {
	*out = *in
	if in.AgentCount != nil {
		in, out := &in.AgentCount, &out.AgentCount
		*out = new(int)
		**out = **in
	}
	if in.SnapshotPeriodSeconds != nil {
		in, out := &in.SnapshotPeriodSeconds, &out.SnapshotPeriodSeconds
		*out = new(int)
		**out = **in
	}
	if in.BackupDeploymentMetadata != nil {
		in, out := &in.BackupDeploymentMetadata, &out.BackupDeploymentMetadata
		*out = new(v1.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.PodTemplateSpec != nil {
		in, out := &in.PodTemplateSpec, &out.PodTemplateSpec
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBBackupSpec.
func (in *FoundationDBBackupSpec) DeepCopy() *FoundationDBBackupSpec {
	if in == nil {
		return nil
	}
	out := new(FoundationDBBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBBackupStatus) DeepCopyInto(out *FoundationDBBackupStatus) {
	*out = *in
	if in.BackupDetails != nil {
		in, out := &in.BackupDetails, &out.BackupDetails
		*out = new(FoundationDBBackupStatusBackupDetails)
		**out = **in
	}
	out.Generations = in.Generations
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBBackupStatus.
func (in *FoundationDBBackupStatus) DeepCopy() *FoundationDBBackupStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBBackupStatusBackupDetails) DeepCopyInto(out *FoundationDBBackupStatusBackupDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBBackupStatusBackupDetails.
func (in *FoundationDBBackupStatusBackupDetails) DeepCopy() *FoundationDBBackupStatusBackupDetails {
	if in == nil {
		return nil
	}
	out := new(FoundationDBBackupStatusBackupDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBCluster) DeepCopyInto(out *FoundationDBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBCluster.
func (in *FoundationDBCluster) DeepCopy() *FoundationDBCluster {
	if in == nil {
		return nil
	}
	out := new(FoundationDBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterAutomationOptions) DeepCopyInto(out *FoundationDBClusterAutomationOptions) {
	*out = *in
	if in.ConfigureDatabase != nil {
		in, out := &in.ConfigureDatabase, &out.ConfigureDatabase
		*out = new(bool)
		**out = **in
	}
	if in.KillProcesses != nil {
		in, out := &in.KillProcesses, &out.KillProcesses
		*out = new(bool)
		**out = **in
	}
	if in.DeletePods != nil {
		in, out := &in.DeletePods, &out.DeletePods
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterAutomationOptions.
func (in *FoundationDBClusterAutomationOptions) DeepCopy() *FoundationDBClusterAutomationOptions {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterAutomationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterFaultDomain) DeepCopyInto(out *FoundationDBClusterFaultDomain) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterFaultDomain.
func (in *FoundationDBClusterFaultDomain) DeepCopy() *FoundationDBClusterFaultDomain {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterFaultDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterList) DeepCopyInto(out *FoundationDBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FoundationDBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterList.
func (in *FoundationDBClusterList) DeepCopy() *FoundationDBClusterList {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterSpec) DeepCopyInto(out *FoundationDBClusterSpec) {
	*out = *in
	if in.SidecarVersions != nil {
		in, out := &in.SidecarVersions, &out.SidecarVersions
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.DatabaseConfiguration.DeepCopyInto(&out.DatabaseConfiguration)
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make(map[string]ProcessSettings, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.ProcessCounts = in.ProcessCounts
	out.FaultDomain = in.FaultDomain
	if in.InstancesToRemove != nil {
		in, out := &in.InstancesToRemove, &out.InstancesToRemove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(corev1.ConfigMap)
		(*in).DeepCopyInto(*out)
	}
	in.MainContainer.DeepCopyInto(&out.MainContainer)
	in.SidecarContainer.DeepCopyInto(&out.SidecarContainer)
	if in.TrustedCAs != nil {
		in, out := &in.TrustedCAs, &out.TrustedCAs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SidecarVariables != nil {
		in, out := &in.SidecarVariables, &out.SidecarVariables
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.AutomationOptions.DeepCopyInto(&out.AutomationOptions)
	in.LockOptions.DeepCopyInto(&out.LockOptions)
	in.Services.DeepCopyInto(&out.Services)
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaim != nil {
		in, out := &in.VolumeClaim, &out.VolumeClaim
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomParameters != nil {
		in, out := &in.CustomParameters, &out.CustomParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PendingRemovals != nil {
		in, out := &in.PendingRemovals, &out.PendingRemovals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterSpec.
func (in *FoundationDBClusterSpec) DeepCopy() *FoundationDBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBClusterStatus) DeepCopyInto(out *FoundationDBClusterStatus) {
	*out = *in
	out.ProcessCounts = in.ProcessCounts
	if in.IncorrectProcesses != nil {
		in, out := &in.IncorrectProcesses, &out.IncorrectProcesses
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IncorrectPods != nil {
		in, out := &in.IncorrectPods, &out.IncorrectPods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MissingProcesses != nil {
		in, out := &in.MissingProcesses, &out.MissingProcesses
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.DatabaseConfiguration.DeepCopyInto(&out.DatabaseConfiguration)
	out.Generations = in.Generations
	out.Health = in.Health
	out.RequiredAddresses = in.RequiredAddresses
	if in.PendingRemovals != nil {
		in, out := &in.PendingRemovals, &out.PendingRemovals
		*out = make(map[string]PendingRemovalState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBClusterStatus.
func (in *FoundationDBClusterStatus) DeepCopy() *FoundationDBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBLiveBackupStatus) DeepCopyInto(out *FoundationDBLiveBackupStatus) {
	*out = *in
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBLiveBackupStatus.
func (in *FoundationDBLiveBackupStatus) DeepCopy() *FoundationDBLiveBackupStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBLiveBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBLiveBackupStatusState) DeepCopyInto(out *FoundationDBLiveBackupStatusState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBLiveBackupStatusState.
func (in *FoundationDBLiveBackupStatusState) DeepCopy() *FoundationDBLiveBackupStatusState {
	if in == nil {
		return nil
	}
	out := new(FoundationDBLiveBackupStatusState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBRestore) DeepCopyInto(out *FoundationDBRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBRestore.
func (in *FoundationDBRestore) DeepCopy() *FoundationDBRestore {
	if in == nil {
		return nil
	}
	out := new(FoundationDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBRestoreList) DeepCopyInto(out *FoundationDBRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FoundationDBRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBRestoreList.
func (in *FoundationDBRestoreList) DeepCopy() *FoundationDBRestoreList {
	if in == nil {
		return nil
	}
	out := new(FoundationDBRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FoundationDBRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBRestoreSpec) DeepCopyInto(out *FoundationDBRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBRestoreSpec.
func (in *FoundationDBRestoreSpec) DeepCopy() *FoundationDBRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(FoundationDBRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBRestoreStatus) DeepCopyInto(out *FoundationDBRestoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBRestoreStatus.
func (in *FoundationDBRestoreStatus) DeepCopy() *FoundationDBRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatus) DeepCopyInto(out *FoundationDBStatus) {
	*out = *in
	in.Client.DeepCopyInto(&out.Client)
	in.Cluster.DeepCopyInto(&out.Cluster)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatus.
func (in *FoundationDBStatus) DeepCopy() *FoundationDBStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusBackupInfo) DeepCopyInto(out *FoundationDBStatusBackupInfo) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]FoundationDBStatusBackupTag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusBackupInfo.
func (in *FoundationDBStatusBackupInfo) DeepCopy() *FoundationDBStatusBackupInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusBackupInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusBackupTag) DeepCopyInto(out *FoundationDBStatusBackupTag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusBackupTag.
func (in *FoundationDBStatusBackupTag) DeepCopy() *FoundationDBStatusBackupTag {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusBackupTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusClientDBStatus) DeepCopyInto(out *FoundationDBStatusClientDBStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusClientDBStatus.
func (in *FoundationDBStatusClientDBStatus) DeepCopy() *FoundationDBStatusClientDBStatus {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusClientDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusClusterClientInfo) DeepCopyInto(out *FoundationDBStatusClusterClientInfo) {
	*out = *in
	if in.SupportedVersions != nil {
		in, out := &in.SupportedVersions, &out.SupportedVersions
		*out = make([]FoundationDBStatusSupportedVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusClusterClientInfo.
func (in *FoundationDBStatusClusterClientInfo) DeepCopy() *FoundationDBStatusClusterClientInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusClusterClientInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusClusterInfo) DeepCopyInto(out *FoundationDBStatusClusterInfo) {
	*out = *in
	in.DatabaseConfiguration.DeepCopyInto(&out.DatabaseConfiguration)
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make(map[string]FoundationDBStatusProcessInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Data = in.Data
	in.Clients.DeepCopyInto(&out.Clients)
	in.Layers.DeepCopyInto(&out.Layers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusClusterInfo.
func (in *FoundationDBStatusClusterInfo) DeepCopy() *FoundationDBStatusClusterInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusConnectedClient) DeepCopyInto(out *FoundationDBStatusConnectedClient) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusConnectedClient.
func (in *FoundationDBStatusConnectedClient) DeepCopy() *FoundationDBStatusConnectedClient {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusConnectedClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusCoordinator) DeepCopyInto(out *FoundationDBStatusCoordinator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusCoordinator.
func (in *FoundationDBStatusCoordinator) DeepCopy() *FoundationDBStatusCoordinator {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusCoordinator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusCoordinatorInfo) DeepCopyInto(out *FoundationDBStatusCoordinatorInfo) {
	*out = *in
	if in.Coordinators != nil {
		in, out := &in.Coordinators, &out.Coordinators
		*out = make([]FoundationDBStatusCoordinator, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusCoordinatorInfo.
func (in *FoundationDBStatusCoordinatorInfo) DeepCopy() *FoundationDBStatusCoordinatorInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusCoordinatorInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusDataStatistics) DeepCopyInto(out *FoundationDBStatusDataStatistics) {
	*out = *in
	out.MovingData = in.MovingData
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusDataStatistics.
func (in *FoundationDBStatusDataStatistics) DeepCopy() *FoundationDBStatusDataStatistics {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusDataStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusLayerInfo) DeepCopyInto(out *FoundationDBStatusLayerInfo) {
	*out = *in
	in.Backup.DeepCopyInto(&out.Backup)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusLayerInfo.
func (in *FoundationDBStatusLayerInfo) DeepCopy() *FoundationDBStatusLayerInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusLayerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusLocalClientInfo) DeepCopyInto(out *FoundationDBStatusLocalClientInfo) {
	*out = *in
	in.Coordinators.DeepCopyInto(&out.Coordinators)
	out.DatabaseStatus = in.DatabaseStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusLocalClientInfo.
func (in *FoundationDBStatusLocalClientInfo) DeepCopy() *FoundationDBStatusLocalClientInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusLocalClientInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusMovingData) DeepCopyInto(out *FoundationDBStatusMovingData) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusMovingData.
func (in *FoundationDBStatusMovingData) DeepCopy() *FoundationDBStatusMovingData {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusMovingData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusProcessInfo) DeepCopyInto(out *FoundationDBStatusProcessInfo) {
	*out = *in
	if in.Locality != nil {
		in, out := &in.Locality, &out.Locality
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusProcessInfo.
func (in *FoundationDBStatusProcessInfo) DeepCopy() *FoundationDBStatusProcessInfo {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusProcessInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDBStatusSupportedVersion) DeepCopyInto(out *FoundationDBStatusSupportedVersion) {
	*out = *in
	if in.ConnectedClients != nil {
		in, out := &in.ConnectedClients, &out.ConnectedClients
		*out = make([]FoundationDBStatusConnectedClient, len(*in))
		copy(*out, *in)
	}
	if in.MaxProtocolClients != nil {
		in, out := &in.MaxProtocolClients, &out.MaxProtocolClients
		*out = make([]FoundationDBStatusConnectedClient, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDBStatusSupportedVersion.
func (in *FoundationDBStatusSupportedVersion) DeepCopy() *FoundationDBStatusSupportedVersion {
	if in == nil {
		return nil
	}
	out := new(FoundationDBStatusSupportedVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LockOptions) DeepCopyInto(out *LockOptions) {
	*out = *in
	if in.DisableLocks != nil {
		in, out := &in.DisableLocks, &out.DisableLocks
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LockOptions.
func (in *LockOptions) DeepCopy() *LockOptions {
	if in == nil {
		return nil
	}
	out := new(LockOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingRemovalState) DeepCopyInto(out *PendingRemovalState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingRemovalState.
func (in *PendingRemovalState) DeepCopy() *PendingRemovalState {
	if in == nil {
		return nil
	}
	out := new(PendingRemovalState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessAddress) DeepCopyInto(out *ProcessAddress) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessAddress.
func (in *ProcessAddress) DeepCopy() *ProcessAddress {
	if in == nil {
		return nil
	}
	out := new(ProcessAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessCounts) DeepCopyInto(out *ProcessCounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessCounts.
func (in *ProcessCounts) DeepCopy() *ProcessCounts {
	if in == nil {
		return nil
	}
	out := new(ProcessCounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessSettings) DeepCopyInto(out *ProcessSettings) {
	*out = *in
	if in.PodTemplate != nil {
		in, out := &in.PodTemplate, &out.PodTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaim != nil {
		in, out := &in.VolumeClaim, &out.VolumeClaim
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomParameters != nil {
		in, out := &in.CustomParameters, &out.CustomParameters
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessSettings.
func (in *ProcessSettings) DeepCopy() *ProcessSettings {
	if in == nil {
		return nil
	}
	out := new(ProcessSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Region) DeepCopyInto(out *Region) {
	*out = *in
	if in.DataCenters != nil {
		in, out := &in.DataCenters, &out.DataCenters
		*out = make([]DataCenter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Region.
func (in *Region) DeepCopy() *Region {
	if in == nil {
		return nil
	}
	out := new(Region)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredAddressSet) DeepCopyInto(out *RequiredAddressSet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredAddressSet.
func (in *RequiredAddressSet) DeepCopy() *RequiredAddressSet {
	if in == nil {
		return nil
	}
	out := new(RequiredAddressSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleCounts) DeepCopyInto(out *RoleCounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleCounts.
func (in *RoleCounts) DeepCopy() *RoleCounts {
	if in == nil {
		return nil
	}
	out := new(RoleCounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
	if in.Headless != nil {
		in, out := &in.Headless, &out.Headless
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionFlags) DeepCopyInto(out *VersionFlags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionFlags.
func (in *VersionFlags) DeepCopy() *VersionFlags {
	if in == nil {
		return nil
	}
	out := new(VersionFlags)
	in.DeepCopyInto(out)
	return out
}
