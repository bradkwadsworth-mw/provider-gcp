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

type SchemaInitParameters struct {

	// The definition of the schema.
	// This should contain a string representing the full definition of the schema
	// that is a valid schema definition of the type specified in type.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of the schema definition
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, PROTOCOL_BUFFER, AVRO.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemaObservation struct {

	// The definition of the schema.
	// This should contain a string representing the full definition of the schema
	// that is a valid schema definition of the type specified in type.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// an identifier for the resource with format projects/{{project}}/schemas/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of the schema definition
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, PROTOCOL_BUFFER, AVRO.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemaParameters struct {

	// The definition of the schema.
	// This should contain a string representing the full definition of the schema
	// that is a valid schema definition of the type specified in type.
	// +kubebuilder:validation:Optional
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of the schema definition
	// Default value is TYPE_UNSPECIFIED.
	// Possible values are: TYPE_UNSPECIFIED, PROTOCOL_BUFFER, AVRO.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SchemaSpec defines the desired state of Schema
type SchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchemaParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider SchemaInitParameters `json:"initProvider,omitempty"`
}

// SchemaStatus defines the observed state of Schema.
type SchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schema is the Schema for the Schemas API. A schema is a format that messages must follow, creating a contract between publisher and subscriber that Pub/Sub will enforce.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Schema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemaSpec   `json:"spec"`
	Status            SchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaList contains a list of Schemas
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schema `json:"items"`
}

// Repository type metadata.
var (
	Schema_Kind             = "Schema"
	Schema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schema_Kind}.String()
	Schema_KindAPIVersion   = Schema_Kind + "." + CRDGroupVersion.String()
	Schema_GroupVersionKind = CRDGroupVersion.WithKind(Schema_Kind)
)

func init() {
	SchemeBuilder.Register(&Schema{}, &SchemaList{})
}
