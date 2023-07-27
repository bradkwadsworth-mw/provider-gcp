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

type ReservationAssignmentInitParameters struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ReservationAssignmentObservation struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/reservations/{{reservation}}/assignments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Output only. The resource name of the assignment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reservation for the resource
	Reservation *string `json:"reservation,omitempty" tf:"reservation,omitempty"`

	// Assignment will remain in PENDING state if no active capacity commitment is present. It will become ACTIVE when some capacity commitment becomes active. Possible values: STATE_UNSPECIFIED, PENDING, ACTIVE
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ReservationAssignmentParameters struct {

	// The resource which will use the reservation. E.g. projects/myproject, folders/123, organizations/456.
	// +kubebuilder:validation:Optional
	Assignee *string `json:"assignee,omitempty" tf:"assignee,omitempty"`

	// Types of job, which could be specified when using the reservation. Possible values: JOB_TYPE_UNSPECIFIED, PIPELINE, QUERY
	// +kubebuilder:validation:Optional
	JobType *string `json:"jobType,omitempty" tf:"job_type,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The reservation for the resource
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Reservation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Reservation *string `json:"reservation,omitempty" tf:"reservation,omitempty"`

	// Reference to a Reservation in bigquery to populate reservation.
	// +kubebuilder:validation:Optional
	ReservationRef *v1.Reference `json:"reservationRef,omitempty" tf:"-"`

	// Selector for a Reservation in bigquery to populate reservation.
	// +kubebuilder:validation:Optional
	ReservationSelector *v1.Selector `json:"reservationSelector,omitempty" tf:"-"`
}

// ReservationAssignmentSpec defines the desired state of ReservationAssignment
type ReservationAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationAssignmentParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider ReservationAssignmentInitParameters `json:"initProvider,omitempty"`
}

// ReservationAssignmentStatus defines the observed state of ReservationAssignment.
type ReservationAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationAssignment is the Schema for the ReservationAssignments API. The BigqueryReservation Assignment resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ReservationAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assignee) || has(self.initProvider.assignee)",message="assignee is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.jobType) || has(self.initProvider.jobType)",message="jobType is a required parameter"
	Spec   ReservationAssignmentSpec   `json:"spec"`
	Status ReservationAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationAssignmentList contains a list of ReservationAssignments
type ReservationAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReservationAssignment `json:"items"`
}

// Repository type metadata.
var (
	ReservationAssignment_Kind             = "ReservationAssignment"
	ReservationAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReservationAssignment_Kind}.String()
	ReservationAssignment_KindAPIVersion   = ReservationAssignment_Kind + "." + CRDGroupVersion.String()
	ReservationAssignment_GroupVersionKind = CRDGroupVersion.WithKind(ReservationAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&ReservationAssignment{}, &ReservationAssignmentList{})
}
