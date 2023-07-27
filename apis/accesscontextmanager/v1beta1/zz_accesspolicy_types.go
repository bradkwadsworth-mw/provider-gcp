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

type AccessPolicyInitParameters struct {

	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Human readable title. Does not affect behavior.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AccessPolicyObservation struct {

	// Time the AccessPolicy was created in UTC.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource name of the AccessPolicy. Format: {policy_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Human readable title. Does not affect behavior.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// Time the AccessPolicy was updated in UTC.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AccessPolicyParameters struct {

	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Folder or project on which this policy is applicable.
	// Format: folders/{{folder_id}} or projects/{{project_id}}
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Human readable title. Does not affect behavior.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// AccessPolicySpec defines the desired state of AccessPolicy
type AccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider AccessPolicyInitParameters `json:"initProvider,omitempty"`
}

// AccessPolicyStatus defines the observed state of AccessPolicy.
type AccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicy is the Schema for the AccessPolicys API. AccessPolicy is a container for AccessLevels (which define the necessary attributes to use GCP services) and ServicePerimeters (which define regions of services able to freely pass data within a perimeter).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parent) || has(self.initProvider.parent)",message="parent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || has(self.initProvider.title)",message="title is a required parameter"
	Spec   AccessPolicySpec   `json:"spec"`
	Status AccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyList contains a list of AccessPolicys
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicy_Kind             = "AccessPolicy"
	AccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicy_Kind}.String()
	AccessPolicy_KindAPIVersion   = AccessPolicy_Kind + "." + CRDGroupVersion.String()
	AccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
}
