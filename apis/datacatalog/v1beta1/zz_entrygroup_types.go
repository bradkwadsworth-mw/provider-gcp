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

type EntryGroupInitParameters struct {

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupID *string `json:"entryGroupId,omitempty" tf:"entry_group_id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// EntryGroup location region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EntryGroupObservation struct {

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupID *string `json:"entryGroupId,omitempty" tf:"entry_group_id,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// EntryGroup location region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EntryGroupParameters struct {

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	// +kubebuilder:validation:Optional
	EntryGroupID *string `json:"entryGroupId,omitempty" tf:"entry_group_id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// EntryGroup location region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// EntryGroupSpec defines the desired state of EntryGroup
type EntryGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntryGroupParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EntryGroupInitParameters `json:"initProvider,omitempty"`
}

// EntryGroupStatus defines the observed state of EntryGroup.
type EntryGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntryGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EntryGroup is the Schema for the EntryGroups API. An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type EntryGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entryGroupId) || has(self.initProvider.entryGroupId)",message="entryGroupId is a required parameter"
	Spec   EntryGroupSpec   `json:"spec"`
	Status EntryGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntryGroupList contains a list of EntryGroups
type EntryGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EntryGroup `json:"items"`
}

// Repository type metadata.
var (
	EntryGroup_Kind             = "EntryGroup"
	EntryGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EntryGroup_Kind}.String()
	EntryGroup_KindAPIVersion   = EntryGroup_Kind + "." + CRDGroupVersion.String()
	EntryGroup_GroupVersionKind = CRDGroupVersion.WithKind(EntryGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&EntryGroup{}, &EntryGroupList{})
}
