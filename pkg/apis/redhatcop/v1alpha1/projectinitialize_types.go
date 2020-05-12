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
// Referencing Origin BuildConfig
type GitSource struct {
	// URI points to the source that will be built. The structure of the source
	// will depend on the type of build to run
	URI string `json:"uri"`

	// Ref is the branch/tag/ref to build.
	Ref string `json:"ref,omitempty"`

	// ProxyConfig defines the proxies to use for the git clone operation
	ProxyConfig `json:"proxyconfig,omitempty"`

	// SourceSecret is the name of a Secret that would be used for setting
	// up the authentication for cloning private repository.
	// The secret contains valid credentials for remote repository, where the
	// data's key represent the authentication method to be used and value is
	// the base64 encoded credentials. Supported auth methods are: ssh-privatekey.
	SourceSecret *kapi.LocalObjectReference `json:"sourceSecret,omitempty"`
}

// ProxyConfig defines what proxies to use for an operation
// Referencing Origin BuildConfig
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
