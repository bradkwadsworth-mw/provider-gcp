/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PerInstanceConfigInitParameters struct {

	// The minimal action to perform on the instance during an update.
	// Default is NONE. Possible values are:
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action,omitempty"`

	// The most disruptive action to perform on the instance during an update.
	// Default is REPLACE. Possible values are:
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState []PreservedStateInitParameters `json:"preservedState,omitempty" tf:"preserved_state,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will not immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy,omitempty"`
}

type PerInstanceConfigObservation struct {

	// an identifier for the resource with format {{project}}/{{zone}}/{{instance_group_manager}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance group manager this instance config is part of.
	InstanceGroupManager *string `json:"instanceGroupManager,omitempty" tf:"instance_group_manager,omitempty"`

	// The minimal action to perform on the instance during an update.
	// Default is NONE. Possible values are:
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action,omitempty"`

	// The most disruptive action to perform on the instance during an update.
	// Default is REPLACE. Possible values are:
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The preserved state for this instance.
	// Structure is documented below.
	PreservedState []PreservedStateObservation `json:"preservedState,omitempty" tf:"preserved_state,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will not immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy,omitempty"`

	// Zone where the containing instance group manager is located
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PerInstanceConfigParameters struct {

	// The instance group manager this instance config is part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.InstanceGroupManager
	// +kubebuilder:validation:Optional
	InstanceGroupManager *string `json:"instanceGroupManager,omitempty" tf:"instance_group_manager,omitempty"`

	// Reference to a InstanceGroupManager in compute to populate instanceGroupManager.
	// +kubebuilder:validation:Optional
	InstanceGroupManagerRef *v1.Reference `json:"instanceGroupManagerRef,omitempty" tf:"-"`

	// Selector for a InstanceGroupManager in compute to populate instanceGroupManager.
	// +kubebuilder:validation:Optional
	InstanceGroupManagerSelector *v1.Selector `json:"instanceGroupManagerSelector,omitempty" tf:"-"`

	// The minimal action to perform on the instance during an update.
	// Default is NONE. Possible values are:
	// +kubebuilder:validation:Optional
	MinimalAction *string `json:"minimalAction,omitempty" tf:"minimal_action,omitempty"`

	// The most disruptive action to perform on the instance during an update.
	// Default is REPLACE. Possible values are:
	// +kubebuilder:validation:Optional
	MostDisruptiveAllowedAction *string `json:"mostDisruptiveAllowedAction,omitempty" tf:"most_disruptive_allowed_action,omitempty"`

	// The name for this per-instance config and its corresponding instance.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The preserved state for this instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PreservedState []PreservedStateParameters `json:"preservedState,omitempty" tf:"preserved_state,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// When true, deleting this config will immediately remove any specified state from the underlying instance.
	// When false, deleting this config will not immediately remove any state from the underlying instance.
	// State will be removed on the next instance recreation or update.
	// +kubebuilder:validation:Optional
	RemoveInstanceStateOnDestroy *bool `json:"removeInstanceStateOnDestroy,omitempty" tf:"remove_instance_state_on_destroy,omitempty"`

	// Zone where the containing instance group manager is located
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.InstanceGroupManager
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("zone",false)
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`

	// Reference to a InstanceGroupManager in compute to populate zone.
	// +kubebuilder:validation:Optional
	ZoneRef *v1.Reference `json:"zoneRef,omitempty" tf:"-"`

	// Selector for a InstanceGroupManager in compute to populate zone.
	// +kubebuilder:validation:Optional
	ZoneSelector *v1.Selector `json:"zoneSelector,omitempty" tf:"-"`
}

type PreservedStateDiskInitParameters struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION.
	// NEVER - detach the disk when the VM is deleted, but do not delete the disk.
	// ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently
	// deleted from the instance group.
	// Default value is NEVER.
	// Possible values are: NEVER, ON_PERMANENT_INSTANCE_DELETION.
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The mode of the disk.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type PreservedStateDiskObservation struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION.
	// NEVER - detach the disk when the VM is deleted, but do not delete the disk.
	// ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently
	// deleted from the instance group.
	// Default value is NEVER.
	// Possible values are: NEVER, ON_PERMANENT_INSTANCE_DELETION.
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The mode of the disk.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URI of an existing persistent disk to attach under the specified device-name in the format
	// projects/project-id/zones/zone/disks/disk-name.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type PreservedStateDiskParameters struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	// The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION.
	// NEVER - detach the disk when the VM is deleted, but do not delete the disk.
	// ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently
	// deleted from the instance group.
	// Default value is NEVER.
	// Possible values are: NEVER, ON_PERMANENT_INSTANCE_DELETION.
	// +kubebuilder:validation:Optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The mode of the disk.
	// Default value is READ_WRITE.
	// Possible values are: READ_ONLY, READ_WRITE.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The URI of an existing persistent disk to attach under the specified device-name in the format
	// projects/project-id/zones/zone/disks/disk-name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Reference to a Disk in compute to populate source.
	// +kubebuilder:validation:Optional
	SourceRef *v1.Reference `json:"sourceRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate source.
	// +kubebuilder:validation:Optional
	SourceSelector *v1.Selector `json:"sourceSelector,omitempty" tf:"-"`
}

type PreservedStateInitParameters struct {

	// Stateful disks for the instance.
	// Structure is documented below.
	Disk []PreservedStateDiskInitParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type PreservedStateObservation struct {

	// Stateful disks for the instance.
	// Structure is documented below.
	Disk []PreservedStateDiskObservation `json:"disk,omitempty" tf:"disk,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type PreservedStateParameters struct {

	// Stateful disks for the instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Disk []PreservedStateDiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

// PerInstanceConfigSpec defines the desired state of PerInstanceConfig
type PerInstanceConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PerInstanceConfigParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider PerInstanceConfigInitParameters `json:"initProvider,omitempty"`
}

// PerInstanceConfigStatus defines the observed state of PerInstanceConfig.
type PerInstanceConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PerInstanceConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PerInstanceConfig is the Schema for the PerInstanceConfigs API. A config defined for a single managed instance that belongs to an instance group manager.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type PerInstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   PerInstanceConfigSpec   `json:"spec"`
	Status PerInstanceConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerInstanceConfigList contains a list of PerInstanceConfigs
type PerInstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerInstanceConfig `json:"items"`
}

// Repository type metadata.
var (
	PerInstanceConfig_Kind             = "PerInstanceConfig"
	PerInstanceConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PerInstanceConfig_Kind}.String()
	PerInstanceConfig_KindAPIVersion   = PerInstanceConfig_Kind + "." + CRDGroupVersion.String()
	PerInstanceConfig_GroupVersionKind = CRDGroupVersion.WithKind(PerInstanceConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&PerInstanceConfig{}, &PerInstanceConfigList{})
}
