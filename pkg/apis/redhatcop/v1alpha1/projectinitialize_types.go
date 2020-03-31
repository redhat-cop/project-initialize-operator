package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectInitializeSpec defines the desired state of ProjectInitialize
type ProjectInitializeSpec struct {
	Team        string `json:"team"`
	Env         string `json:"env"`
	Cluster     string `json:"cluster,omitempty"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"desc"`
	QuotaSize   string `json:"quotaSize,omitempty"`
}

// ProjectInitializeStatus defines the observed state of ProjectInitialize
type ProjectInitializeStatus struct {
	NamespaceCreated bool `json:"namespaceCreated,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectInitialize is the Schema for the projectinitializes API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=projectinitializes,scope=Cluster
type ProjectInitialize struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectInitializeSpec   `json:"spec,omitempty"`
	Status ProjectInitializeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectInitializeList contains a list of ProjectInitialize
type ProjectInitializeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectInitialize `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectInitialize{}, &ProjectInitializeList{})
}
