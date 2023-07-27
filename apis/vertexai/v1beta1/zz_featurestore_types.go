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

type FeaturestoreEncryptionSpecInitParameters struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type FeaturestoreEncryptionSpecObservation struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type FeaturestoreEncryptionSpecParameters struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the compute resource is created.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type FeaturestoreInitParameters struct {

	// If set, both of the online and offline data storage will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []FeaturestoreEncryptionSpecInitParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A set of key/value label pairs to assign to this Featurestore.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Config for online serving resources.
	// Structure is documented below.
	OnlineServingConfig []OnlineServingConfigInitParameters `json:"onlineServingConfig,omitempty" tf:"online_serving_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type FeaturestoreObservation struct {

	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// If set, both of the online and offline data storage will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []FeaturestoreEncryptionSpecObservation `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// Used to perform consistent read-modify-write updates.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/featurestores/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to this Featurestore.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Config for online serving resources.
	// Structure is documented below.
	OnlineServingConfig []OnlineServingConfigObservation `json:"onlineServingConfig,omitempty" tf:"online_serving_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type FeaturestoreParameters struct {

	// If set, both of the online and offline data storage will be secured by this key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionSpec []FeaturestoreEncryptionSpecParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A set of key/value label pairs to assign to this Featurestore.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Config for online serving resources.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	OnlineServingConfig []OnlineServingConfigParameters `json:"onlineServingConfig,omitempty" tf:"online_serving_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type OnlineServingConfigInitParameters struct {

	// The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	FixedNodeCount *float64 `json:"fixedNodeCount,omitempty" tf:"fixed_node_count,omitempty"`

	// Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.
	// Structure is documented below.
	Scaling []ScalingInitParameters `json:"scaling,omitempty" tf:"scaling,omitempty"`
}

type OnlineServingConfigObservation struct {

	// The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	FixedNodeCount *float64 `json:"fixedNodeCount,omitempty" tf:"fixed_node_count,omitempty"`

	// Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.
	// Structure is documented below.
	Scaling []ScalingObservation `json:"scaling,omitempty" tf:"scaling,omitempty"`
}

type OnlineServingConfigParameters struct {

	// The number of nodes for each cluster. The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	// +kubebuilder:validation:Optional
	FixedNodeCount *float64 `json:"fixedNodeCount,omitempty" tf:"fixed_node_count,omitempty"`

	// Online serving scaling configuration. Only one of fixedNodeCount and scaling can be set. Setting one will reset the other.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Scaling []ScalingParameters `json:"scaling,omitempty" tf:"scaling,omitempty"`
}

type ScalingInitParameters struct {

	// The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The minimum number of nodes to scale down to. Must be greater than or equal to 1.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

type ScalingObservation struct {

	// The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The minimum number of nodes to scale down to. Must be greater than or equal to 1.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

type ScalingParameters struct {

	// The maximum number of nodes to scale up to. Must be greater than minNodeCount, and less than or equal to 10 times of 'minNodeCount'.
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// The minimum number of nodes to scale down to. Must be greater than or equal to 1.
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

// FeaturestoreSpec defines the desired state of Featurestore
type FeaturestoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FeaturestoreParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider FeaturestoreInitParameters `json:"initProvider,omitempty"`
}

// FeaturestoreStatus defines the observed state of Featurestore.
type FeaturestoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FeaturestoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Featurestore is the Schema for the Featurestores API. A collection of DataItems and Annotations on them.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Featurestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FeaturestoreSpec   `json:"spec"`
	Status            FeaturestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FeaturestoreList contains a list of Featurestores
type FeaturestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Featurestore `json:"items"`
}

// Repository type metadata.
var (
	Featurestore_Kind             = "Featurestore"
	Featurestore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Featurestore_Kind}.String()
	Featurestore_KindAPIVersion   = Featurestore_Kind + "." + CRDGroupVersion.String()
	Featurestore_GroupVersionKind = CRDGroupVersion.WithKind(Featurestore_Kind)
)

func init() {
	SchemeBuilder.Register(&Featurestore{}, &FeaturestoreList{})
}
