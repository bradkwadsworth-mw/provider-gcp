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

type ProjectExclusionInitParameters struct {

	// A human-readable description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectExclusionObservation struct {

	// A human-readable description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// an identifier for the resource with format projects/{{project}}/exclusions/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ProjectExclusionParameters struct {

	// A human-readable description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See Advanced Log Filters for information on how to
	// write a filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// The project to create the exclusion in. If omitted, the project associated with the provider is
	// used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ProjectExclusionSpec defines the desired state of ProjectExclusion
type ProjectExclusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectExclusionParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ProjectExclusionInitParameters `json:"initProvider,omitempty"`
}

// ProjectExclusionStatus defines the observed state of ProjectExclusion.
type ProjectExclusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectExclusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectExclusion is the Schema for the ProjectExclusions API. Manages a project-level logging exclusion.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ProjectExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filter) || has(self.initProvider.filter)",message="filter is a required parameter"
	Spec   ProjectExclusionSpec   `json:"spec"`
	Status ProjectExclusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectExclusionList contains a list of ProjectExclusions
type ProjectExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectExclusion `json:"items"`
}

// Repository type metadata.
var (
	ProjectExclusion_Kind             = "ProjectExclusion"
	ProjectExclusion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectExclusion_Kind}.String()
	ProjectExclusion_KindAPIVersion   = ProjectExclusion_Kind + "." + CRDGroupVersion.String()
	ProjectExclusion_GroupVersionKind = CRDGroupVersion.WithKind(ProjectExclusion_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectExclusion{}, &ProjectExclusionList{})
}
