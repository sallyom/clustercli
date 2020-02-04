package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterCLI is the Schema for the cluster cli API
// +k8s:openapi-gen=true
// +genclient
type ClusterCLI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterCLISpec   `json:"spec,omitempty"`
	Status ClusterCLIStatus `json:"status,omitempty"`
}

// ClusterCLISpec defines the desired state of ClusterCLI
type ClusterCLISpec struct {
	// Description of ClusterCLI
	Description string `json:"description"`
	// AvailableTarget defines extract targets for ClusterCLIs
	AvailableTargets []ExtractTarget `json:"items"`
}

// ClusterCLIStatus defines the observed state of ClusterCLI
type ClusterCLIStatus struct {
	Installed bool   `json:"installed"`
	Version   string `json:"version"`
}

// ExtractTarget defines the targets available for a ClusterCLI
type ExtractTarget struct {
	OS      string  `json:"os"`
	Command string  `json:"command"`
	Mapping Mapping `json:"mapping"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterCLIList contains a list of ClusterCLI
type ClusterCLIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCLI `json:"items"`
}

type Mapping struct {
	// Name is for associating callback metadata with a mapping
	Name string `json:"name"`
	// Image is the raw input image to extract
	Image string `json:"image"`
	// From is the directory or file in the image to extract
	From string
	// To is the directory to extract the contents of the directory or the named file into.
	To string
}
