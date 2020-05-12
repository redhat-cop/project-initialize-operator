package v1alpha1

import (
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectInitializeSpec defines the desired state of ProjectInitialize
type ProjectInitializeSpec struct {
	Team        string     `json:"team"`
	Env         string     `json:"env"`
	Cluster     string     `json:"cluster,omitempty"`
	DisplayName string     `json:"displayName"`
	Desc        string     `json:"desc"`
	QuotaSize   string     `json:"quotaSize,omitempty"`
	GitSource   *GitSource `json:"gitSource,omitempty"`
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

// GitSource
// +k8s:openapi-gen=true
type GitSource struct {
	// The host platform type
	// +kubebuilder:validation:Enum=GitHub;GitLab;BitBucket
	GitHost GitHost `json:"gittype"`
	// The token or credentials of the users account
	AccountSecret *kapi.LocalObjectReference `json:"accountSecret"`
	// Private or public repository
	Private bool `json:"private"`
	// Description of the repository to create/clone
	Desc string `json:"desc"`
	// Optional GIT clone source for cloning existing template instead of creating new blank repo
	GitClone *GitClone `json:"gittype,omitempty"`
}

// GitClone
// +k8s:openapi-gen=true
type GitClone struct {
	// URI points to the source that will be built. The structure of the source
	// will depend on the type of build to run
	CloneURI string `json:"uri,omitempty"`
	// Ref is the branch/tag/ref to build.
	CloneRef string `json:"ref,omitempty"`
	// ProxyConfig defines the proxies to use for the git clone operation
	CloneProxyConfig `json:"proxyconfig,omitempty"`
	// SourceSecret is the name of a Secret that would be used for setting
	// up the authentication for cloning private repository.
	// The secret contains valid credentials for remote repository, where the
	// data's key represent the authentication method to be used and value is
	// the base64 encoded credentials. Supported auth methods are: ssh-privatekey.
	CloneSourceSecret *kapi.LocalObjectReference `json:"sourceSecret,omitempty"`
}

// GitHost specifies what type of API to use for accessing hosting platform
type GitHost string

const (
	GitLab    GitHost = "GitLab"
	BitBucket GitHost = "BitBucket"
	GitHub    GitHost = "GitHub"
)

// ProxyConfig defines what proxies to use for an operation
type ProxyConfig struct {
	// HTTPProxy is a proxy used to reach the git repository over http
	HTTPProxy *string

	// HTTPSProxy is a proxy used to reach the git repository over https
	HTTPSProxy *string

	// NoProxy is the list of domains for which the proxy should not be used
	NoProxy *string
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
