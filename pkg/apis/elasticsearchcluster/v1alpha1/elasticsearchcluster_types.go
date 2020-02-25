package v1alpha1

import (
	v1 "k8s.io/api/apps/v1"
	storagev1 "k8s.io/api/storage/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers/core"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ElasticsearchNodeSpec struct {
	//Number of node replicas
	replicas int `json:"replicas"`
	//Storageclass used by the node for persistent data storage
	storageClass storagev1.StorageClass `json:"storage_class"`
	//Storage used by node for persistent data
	storage *corev1.PersistentVolumeClaimSpec `json:"storage"`
	//Amount of resources allocated to the node
	resources corev1.ResourceRequirements `json:"resources"`
	//Environment variables for the node
	envs []corev1.EnvVar `json:"envs"`
}

type ClusterTopology struct {
	//Spec of master nodes in the cluster
	MasterNodeSpec *ElasticsearchNodeSpec `json:"master_node_spec"`
	//Spec of data nodes in the cluster
	DataNodeSpec *ElasticsearchNodeSpec `json:"data_node_spec"`
	//Spec of ingest nodes in the cluster
	IngestNodeSpec *ElasticsearchNodeSpec `json:"ingest_node_spec"`
}

type ServiceTemplateSpec struct {
	//List of ports for es
	Ports []ServicePort `json:"ports"`
	//Type of service used for es
	Type corev1.ServiceType `json:"type"`
}

// ServicePort contains information on service's port.
type ServicePort struct {
	//name of the port
	name string `json:"name,omitempty"`

	// The port exposed by this service.
	port int32 `json:"port"`

	//node port for the service
	nodePort int32 `json:"nodePort,omitempty"`
}

// ElasticsearchClusterSpec defines the desired state of ElasticsearchCluster
type ElasticsearchClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	//Name of the elasticsearch cluster
	clusterName string `json:"cluster_name"`
	//Version of the elasticsearch cluster
	version string `json:"version"`
	//Topology of the elasticsearch cluster
	topology *ClusterTopology `json:"topology"`
	//Enable SSL - with searchguard
	enableSSL bool `json:"enable_ssl"`
	//REST endpoint port
	restPort int `json:"rest_port"`
	//Template for es service
	serviceTemplate ServiceTemplateSpec `json:"service_template"`
	//Update strategy for es statefulsets
	updateStrategy v1.StatefulSetUpdateStrategy `json:"update_strategy"`
}

// ElasticsearchClusterStatus defines the observed state of ElasticsearchCluster
type ElasticsearchClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticsearchCluster is the Schema for the elasticsearchclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=elasticsearchclusters,scope=Namespaced
type ElasticsearchCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticsearchClusterSpec   `json:"spec,omitempty"`
	Status ElasticsearchClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ElasticsearchClusterList contains a list of ElasticsearchCluster
type ElasticsearchClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElasticsearchCluster{}, &ElasticsearchClusterList{})
}
