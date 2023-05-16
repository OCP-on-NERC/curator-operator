package v1alpha1

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FetchDataSpec defines the desired state of FetchData
type FetchDataSpec struct {
	//Namespace in which cron job is created
	CronjobNamespace string `json:"cronjobNamespace,omitempty"`

	//Schedule period for the CronJob
	Schedule string `json:"schedule,omitempty"`

	//Koku metrics pvc zipped files storage path
	BackupSrc string `json:"backupSrc,omitempty"`

	//Koku-metrics-pvc path to unzip files
	UnzipDir string `json:"unzipDir,omitempty"`

	// Value for the Database Name Environment Variable
	DatabaseName string `json:"databaseName,omitempty"`

	//Value for the Database Password Environment Variable
	DatabasePassword string `json:"databasePassword,omitempty"`

	// Value for the Database User Environment Variable
	DatabaseUser string `json:"databaseUser,omitempty"`

	// Value for the Database HostName Environment Variable
	DatabaseHostName string `json:"databaseHostName,omitempty"`

	// Value for the Database Environment Variable in order to define the port which it should use. It will be used in its container as well
	DatabasePort string `json:"databasePort,omitempty"`
}

// FetchDataStatus defines the observed state of FetchData
type FetchDataStatus struct {
	//Name of the CronJob object created and managed by it
	CronJobName string `json:"cronJobName"`

	//CronJobStatus represents the current state of a cronjob
	CronJobStatus batchv1.CronJobStatus `json:"cronJobStatus"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FetchData is the Schema for the fetchdata API
type FetchData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FetchDataSpec   `json:"spec,omitempty"`
	Status FetchDataStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FetchDataList contains a list of FetchData
type FetchDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FetchData `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FetchData{}, &FetchDataList{})
}
