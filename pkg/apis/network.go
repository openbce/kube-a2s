package apis

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Node represents a Kubernetes node
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeSpec   `json:"spec,omitempty"`
	Status            NodeStatus `json:"status,omitempty"`
}

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	// HostName is the name of the host this node is running on
	HostName string `json:"hostName"`
	// Labels represents the labels on the node
	Labels map[string]string `json:"labels,omitempty"`
}

// NodeStatus defines the observed state of Node
type NodeStatus struct {
	// Phase represents the current phase of the node
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the node's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// HCA represents a High Computing Accelerator
type HCA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HCASpec   `json:"spec,omitempty"`
	Status            HCAStatus `json:"status,omitempty"`
}

// HCASpec defines the desired state of HCA
type HCASpec struct {
	// HostName is the name of the host this HCA is installed on
	HostName string `json:"hostName"`
	// Type represents the type of HCA
	Type string `json:"type"`
	// Resources represents the available resources on this HCA
	Resources map[string]string `json:"resources,omitempty"`
}

// HCAStatus defines the observed state of HCA
type HCAStatus struct {
	// Phase represents the current phase of the HCA
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the HCA's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Partition represents a resource partition in the system
type Partition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartitionSpec   `json:"spec,omitempty"`
	Status            PartitionStatus `json:"status,omitempty"`
}

// PartitionSpec defines the desired state of Partition
type PartitionSpec struct {
	// NodeSelector represents the nodes that belong to this partition
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Resources represents the available resources in this partition
	Resources map[string]string `json:"resources,omitempty"`
}

// PartitionStatus defines the observed state of Partition
type PartitionStatus struct {
	// Phase represents the current phase of the partition
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the partition's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Port represents a network port configuration
type Port struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PortSpec   `json:"spec,omitempty"`
	Status            PortStatus `json:"status,omitempty"`
}

// PortSpec defines the desired state of Port
type PortSpec struct {
	// Number represents the port number
	Number int32 `json:"number"`
	// Protocol represents the protocol used (TCP/UDP)
	Protocol string `json:"protocol"`
	// TargetPort represents the target port number
	TargetPort int32 `json:"targetPort"`
}

// PortStatus defines the observed state of Port
type PortStatus struct {
	// Phase represents the current phase of the port
	Phase string `json:"phase,omitempty"`
	// Conditions represents the latest available observations of the port's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
