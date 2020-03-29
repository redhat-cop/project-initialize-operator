package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectInitializeQuotaSpec defines the desired state of ProjectInitializeQuota
type ProjectInitializeQuotaSpec struct {
	corev1.ResourceQuotaSpec `json:"resourceQuotaSpec"`
}

// ProjectInitializeQuotaStatus defines the observed state of ProjectInitializeQuota
type ProjectInitializeQuotaStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectInitializeQuota is the Schema for the projectinitializequota API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=projectinitializequota,scope=Cluster
type ProjectInitializeQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectInitializeQuotaSpec   `json:"spec,omitempty"`
	Status ProjectInitializeQuotaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectInitializeQuotaList contains a list of ProjectInitializeQuota
type ProjectInitializeQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectInitializeQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectInitializeQuota{}, &ProjectInitializeQuotaList{})
}
