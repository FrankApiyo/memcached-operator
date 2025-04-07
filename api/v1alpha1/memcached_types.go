/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcachedSpec defines the desired state of Memcached.
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

  // Size configures number of Memcached instances
  // The following marker use OpenAPI v3 schema to validate the values
  // +kubebuilder:validation:Minimum=1
  // +kubebuilder:validation:Maximum=3
  // +kubebuilder:validation:ExclusiveMaximum=false
  Size int32 `json:"size,omitempty"`
}

// MemcachedStatus defines the observed state of Memcached.
type MemcachedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
  // Memcached.status.conditions.types are: "Available", "Processing", and "Degraded"
  // Memcached.status.conditions.status are one of True, False, Unknown.
  // Memcached.status.condition.reason the value should be a camelCase string and producers of specific
  // condition types may define expected values and meanings for this field, and whether the values
  // are considered a guaranteed API.
  // Memcached.status.conditions.message is a human readable message indicating the details about the transition.

  Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Memcached is the Schema for the memcacheds API.
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemcachedList contains a list of Memcached.
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
