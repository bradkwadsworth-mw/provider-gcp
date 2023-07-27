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

type EnvironmentInitParameters struct {

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs []VersionConfigsInitParameters `json:"versionConfigs,omitempty" tf:"version_configs,omitempty"`
}

type EnvironmentObservation struct {

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format {{parent}}/environments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the environment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Agent to create an Environment for.
	// Format: projects//locations//agents/.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs []VersionConfigsObservation `json:"versionConfigs,omitempty" tf:"version_configs,omitempty"`
}

type EnvironmentParameters struct {

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Agent to create an Environment for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VersionConfigs []VersionConfigsParameters `json:"versionConfigs,omitempty" tf:"version_configs,omitempty"`
}

type VersionConfigsInitParameters struct {
}

type VersionConfigsObservation struct {

	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VersionConfigsParameters struct {

	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Version
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// Reference to a Version in dialogflowcx to populate version.
	// +kubebuilder:validation:Optional
	VersionRef *v1.Reference `json:"versionRef,omitempty" tf:"-"`

	// Selector for a Version in dialogflowcx to populate version.
	// +kubebuilder:validation:Optional
	VersionSelector *v1.Selector `json:"versionSelector,omitempty" tf:"-"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider EnvironmentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API. Represents an environment for an agent.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.versionConfigs) || has(self.initProvider.versionConfigs)",message="versionConfigs is a required parameter"
	Spec   EnvironmentSpec   `json:"spec"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	Environment_Kind             = "Environment"
	Environment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Environment_Kind}.String()
	Environment_KindAPIVersion   = Environment_Kind + "." + CRDGroupVersion.String()
	Environment_GroupVersionKind = CRDGroupVersion.WithKind(Environment_Kind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
