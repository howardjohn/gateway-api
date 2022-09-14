/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api,scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:unservedversion
// +kubebuilder:printcolumn:name="Controller",type=string,JSONPath=`.spec.controllerName`
// +kubebuilder:printcolumn:name="Accepted",type=string,JSONPath=`.status.conditions[?(@.type=="Accepted")].status`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`,priority=1

type Mesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of MeshSpec.
	Spec MeshSpec `json:"spec"`

	// Status defines the current state of Mesh.
	//
	// +kubebuilder:default={conditions: {{type: "Accepted", status: "Unknown", message: "Waiting for controller", reason: "Waiting", lastTransitionTime: "1970-01-01T00:00:00Z"}}}
	Status MeshStatus `json:"status,omitempty"`
}


type MeshSpec struct {
	ControllerName GatewayController `json:"controllerName"`

	// ParametersRef is a reference to a resource that contains the configuration
	// parameters corresponding to the Mesh. This is optional if the
	// controller does not require any additional configuration.
	//
	// ParametersRef can reference a standard Kubernetes resource, i.e. ConfigMap,
	// or an implementation-specific custom resource. The resource can be
	// cluster-scoped or namespace-scoped.
	//
	// If the referent cannot be found, the Mesh's "InvalidParameters"
	// status condition will be true.
	//
	// Support: Custom
	//
	// +optional
	ParametersRef *ParametersReference `json:"parametersRef,omitempty"`

	// Description helps describe a Mesh with more details.
	//
	// +kubebuilder:validation:MaxLength=64
	// +optional
	Description *string `json:"description,omitempty"`
}

// MeshConditionType is the type for status conditions on
// Mesh resources. This type should be used with the
// MeshStatus.Conditions field.
type MeshConditionType string

// MeshConditionReason defines the set of reasons that explain why a
// particular Mesh condition type has been raised.
type MeshConditionReason string

const (
	MeshConditionStatusAccepted MeshConditionType = "Accepted"

	MeshReasonAccepted MeshConditionReason = "Accepted"

	MeshReasonInvalidParameters MeshConditionReason = "InvalidParameters"

	MeshReasonWaiting MeshConditionReason = "Waiting"
)

// MeshStatus is the current status for the Mesh.
type MeshStatus struct {
	// Conditions is the current status from the controller for
	// this Mesh.
	//
	// Controllers should prefer to publish conditions using values
	// of MeshConditionType for the type of each Condition.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Accepted", status: "Unknown", message: "Waiting for controller", reason: "Waiting", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// MeshList contains a list of Mesh
type MeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mesh `json:"items"`
}
