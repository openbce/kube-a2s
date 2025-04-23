package apis

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Host represents a physical or virtual host in the system
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec,omitempty"`
	Status            HostStatus `json:"status,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	// NodeName is the name of the node this host is associated with
	NodeName string `json:"nodeName"`
	// HCAs represents the list of HCAs on this host
	HCAs []string `json:"hcas,omitempty"`
	// Resources represents the available resources on this host
	Resources map[string]string `json:"resources,omitempty"`
}

// HostStatus defines the observed state of Host
type HostStatus struct {
	// Phase represents the current phase of the host
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the host's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// PodGroup represents a group of pods that share common characteristics
type PodGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PodGroupSpec   `json:"spec,omitempty"`
	Status            PodGroupStatus `json:"status,omitempty"`
}

// PodGroupSpec defines the desired state of PodGroup
type PodGroupSpec struct {
	// MinMember represents the minimum number of pods in the group
	MinMember int32 `json:"minMember"`
	// MaxMember represents the maximum number of pods in the group
	MaxMember int32 `json:"maxMember"`
	// Priority represents the priority of the pod group
	Priority int32 `json:"priority"`
	// Resources represents the resource requirements for the pod group
	Resources map[string]string `json:"resources,omitempty"`
}

// PodGroupStatus defines the observed state of PodGroup
type PodGroupStatus struct {
	// Phase represents the current phase of the pod group
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the pod group's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
