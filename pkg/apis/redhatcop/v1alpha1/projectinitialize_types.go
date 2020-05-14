package v1alpha1

import (
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectInitializeSpec defines the desired state of ProjectInitialize
type ProjectInitializeSpec struct {
	Team        string       `json:"team"`
	Env         string       `json:"env"`
	Cluster     string       `json:"cluster,omitempty"`
	DisplayName string       `json:"displayName"`
	Desc        string       `json:"desc"`
	QuotaSize   string       `json:"quotaSize,omitempty"`
	Git         *Git         `json:"git,omitempty"`
	GitTemplate *GitTemplate `json:"gitTemplate,omitempty"`
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
	// The host platform type
	// +kubebuilder:validation:Enum=GitHub;GitLab;BitBucket
	GitHost GitHost `json:"gittype"`
	// The token or credentials of the account that the template will copy to
	AccountSecret *kapi.LocalObjectReference `json:"accountSecret"`
	// Private or public repository
	Private bool `json:"private"`
	// Description of the repository to create/clone
	Desc string `json:"desc"`
	// Account to copy template
	Owner string `json:"owner"`
	// Ref is the branch/tag/ref to build.
	Ref string `json:"ref,omitempty"`
}

// GitHost specifies what type of API to use for accessing hosting platform
type GitHost string

const (
	GitLab    GitHost = "GitLab"
	BitBucket GitHost = "BitBucket"
	GitHub    GitHost = "GitHub"
)

// Git
// +k8s:openapi-gen=true
type GitTemplate struct {
	// The token of the account that the template will copy to
	AccountSecret *kapi.LocalObjectReference `json:"accountSecret"`
	// Account of the template to copy
	Owner string `json:"owner"`
	// Repo to copy
	Repo string `json:"repo"`
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
