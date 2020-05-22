package v1alpha1

import (
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectInitializeSpec defines the desired state of ProjectInitialize
type ProjectInitializeSpec struct {
	Team             string            `json:"team"`
	Env              string            `json:"env"`
	Cluster          string            `json:"cluster,omitempty"`
	DisplayName      string            `json:"displayName"`
	Desc             string            `json:"desc"`
	QuotaSize        string            `json:"quotaSize,omitempty"`
	Git              *Git              `json:"git,omitempty"`
	GitTemplate      *GitTemplate      `json:"gitTemplate,omitempty"`
	NamespaceDetails *NamespaceDetails `json:"namespaceDetails,omitempty"`
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

// Git
// +k8s:openapi-gen=true
type Git struct {
	// The host provider type
	// +kubebuilder:validation:Enum=GitHub;GitLab;BitBucket
	Provider Provider `json:"provider"`
	// Private or public repository
	Private bool `json:"private"`
	// Description of the repository to create/clone
	Desc string `json:"desc"`
	// Account to copy template
	Owner string `json:"owner"`
	// the suffix that files the teamname in the format teamname-suffix
	Suffix string `json:"suffix"`
	// The token of the account that the template will copy to
	AccountSecret *kapi.ObjectReference `json:"accountSecret"`
}

// GitHost specifies what type of API to use for accessing hosting platform
type Provider string

const (
	GitLab    Provider = "GitLab"
	BitBucket Provider = "BitBucket"
	GitHub    Provider = "GitHub"
)

// Git
// +k8s:openapi-gen=true
type GitTemplate struct {
	// Account of the template to copy
	Owner string `json:"owner"`
	// Repo to copy
	Repo string `json:"repo"`
}

type NamespaceDetails struct {
	Labels      map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
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
