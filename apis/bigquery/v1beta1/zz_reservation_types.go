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

type AutoscaleInitParameters struct {

	// Number of slots to be scaled when needed.
	MaxSlots *float64 `json:"maxSlots,omitempty" tf:"max_slots,omitempty"`
}

type AutoscaleObservation struct {

	// (Output)
	// The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].
	CurrentSlots *float64 `json:"currentSlots,omitempty" tf:"current_slots,omitempty"`

	// Number of slots to be scaled when needed.
	MaxSlots *float64 `json:"maxSlots,omitempty" tf:"max_slots,omitempty"`
}

type AutoscaleParameters struct {

	// Number of slots to be scaled when needed.
	// +kubebuilder:validation:Optional
	MaxSlots *float64 `json:"maxSlots,omitempty" tf:"max_slots,omitempty"`
}

type ReservationInitParameters struct {

	// The configuration parameters for the auto scaling feature.
	// Structure is documented below.
	Autoscale []AutoscaleInitParameters `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency *float64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`

	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty" tf:"ignore_idle_slots,omitempty"`

	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
	// If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty" tf:"multi_region_auxiliary,omitempty"`

	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *float64 `json:"slotCapacity,omitempty" tf:"slot_capacity,omitempty"`
}

type ReservationObservation struct {

	// The configuration parameters for the auto scaling feature.
	// Structure is documented below.
	Autoscale []AutoscaleObservation `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency *float64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`

	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/reservations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty" tf:"ignore_idle_slots,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
	// If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty" tf:"multi_region_auxiliary,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *float64 `json:"slotCapacity,omitempty" tf:"slot_capacity,omitempty"`
}

type ReservationParameters struct {

	// The configuration parameters for the auto scaling feature.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Autoscale []AutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	// +kubebuilder:validation:Optional
	Concurrency *float64 `json:"concurrency,omitempty" tf:"concurrency,omitempty"`

	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	// +kubebuilder:validation:Optional
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty" tf:"ignore_idle_slots,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
	// If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	// +kubebuilder:validation:Optional
	MultiRegionAuxiliary *bool `json:"multiRegionAuxiliary,omitempty" tf:"multi_region_auxiliary,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	// +kubebuilder:validation:Optional
	SlotCapacity *float64 `json:"slotCapacity,omitempty" tf:"slot_capacity,omitempty"`
}

// ReservationSpec defines the desired state of Reservation
type ReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ReservationInitParameters `json:"initProvider,omitempty"`
}

// ReservationStatus defines the observed state of Reservation.
type ReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reservation is the Schema for the Reservations API. A reservation is a mechanism used to guarantee BigQuery slots to users.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Reservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.slotCapacity) || has(self.initProvider.slotCapacity)",message="slotCapacity is a required parameter"
	Spec   ReservationSpec   `json:"spec"`
	Status ReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationList contains a list of Reservations
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reservation `json:"items"`
}

// Repository type metadata.
var (
	Reservation_Kind             = "Reservation"
	Reservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reservation_Kind}.String()
	Reservation_KindAPIVersion   = Reservation_Kind + "." + CRDGroupVersion.String()
	Reservation_GroupVersionKind = CRDGroupVersion.WithKind(Reservation_Kind)
)

func init() {
	SchemeBuilder.Register(&Reservation{}, &ReservationList{})
}
