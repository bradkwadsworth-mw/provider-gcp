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

type DatasetInitParameters struct {

	// The location for the Dataset.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type DatasetObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/datasets/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the Dataset.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The fully qualified name of this dataset
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type DatasetParameters struct {

	// The location for the Dataset.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name for the Dataset.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
	// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
	// (e.g., HL7 messages) where no explicit timezone is specified.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

// DatasetSpec defines the desired state of Dataset
type DatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider DatasetInitParameters `json:"initProvider,omitempty"`
}

// DatasetStatus defines the observed state of Dataset.
type DatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dataset is the Schema for the Datasets API. A Healthcare
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   DatasetSpec   `json:"spec"`
	Status DatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetList contains a list of Datasets
type DatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataset `json:"items"`
}

// Repository type metadata.
var (
	Dataset_Kind             = "Dataset"
	Dataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dataset_Kind}.String()
	Dataset_KindAPIVersion   = Dataset_Kind + "." + CRDGroupVersion.String()
	Dataset_GroupVersionKind = CRDGroupVersion.WithKind(Dataset_Kind)
)

func init() {
	SchemeBuilder.Register(&Dataset{}, &DatasetList{})
}
