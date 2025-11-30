package types

// Patient represents detailed information about a DICOM patient
type Patient struct {
	// Unique identifier of the patient in Orthanc
	ID string `json:"ID"`

	// Whether the patient is stable (no more instances expected)
	IsStable bool `json:"IsStable"`

	// Last update timestamp
	LastUpdate string `json:"LastUpdate"`

	// Main DICOM tags for the patient
	MainDicomTags PatientDicomTags `json:"MainDicomTags"`

	// List of study IDs for this patient
	Studies []string `json:"Studies"`

	// Type of resource (should be "Patient")
	Type string `json:"Type"`
}

// PatientDicomTags represents the main DICOM tags for a patient
type PatientDicomTags struct {
	// Patient Name
	PatientName string `json:"PatientName,omitempty"`

	// Patient ID
	PatientID string `json:"PatientID,omitempty"`

	// Patient Birth Date
	PatientBirthDate string `json:"PatientBirthDate,omitempty"`

	// Patient Sex
	PatientSex string `json:"PatientSex,omitempty"`

	// Other Patient IDs
	OtherPatientIDs string `json:"OtherPatientIDs,omitempty"`

	// Patient's Age
	PatientAge string `json:"PatientAge,omitempty"`

	// Patient's Size
	PatientSize string `json:"PatientSize,omitempty"`

	// Patient's Weight
	PatientWeight string `json:"PatientWeight,omitempty"`

	// Patient Comments
	PatientComments string `json:"PatientComments,omitempty"`

	// Issuer of Patient ID
	IssuerOfPatientID string `json:"IssuerOfPatientID,omitempty"`
}

// PatientQueryParams represents query parameters for GET /patients
type PatientQueryParams struct {
	// Expand the response to include full details
	Expand bool

	// Return results starting from this index
	Since int

	// Maximum number of results to return
	Limit int
}

// PatientStatistics represents statistics about a patient
type PatientStatistics struct {
	// Total disk size in bytes
	DiskSize string `json:"DiskSize"`

	// Disk size in megabytes
	DiskSizeMB int `json:"DiskSizeMB"`

	// Number of instances
	CountInstances int `json:"CountInstances"`

	// Number of series
	CountSeries int `json:"CountSeries"`

	// Number of studies
	CountStudies int `json:"CountStudies"`

	// Uncompressed size in bytes
	UncompressedSize string `json:"UncompressedSize,omitempty"`

	// Uncompressed size in megabytes
	UncompressedSizeMB int `json:"UncompressedSizeMB,omitempty"`
}

// AnonymizeRequest represents a request to anonymize a series
type PatientAnonymizeRequest struct {
	// If true, the REST API will return a Job ID and the job will be put in a queue
	Asynchronous *bool `json:"Asynchronous,omitempty"`

	// DicomVersion to use for anonymization
	DicomVersion string `json:"DicomVersion,omitempty"`

	// Force operation even if it would create an invalid DICOM file
	Force *bool `json:"Force,omitempty"`

	// By default orthanc does not exclude the source series, use set this to falso to delete after the anonymization proccess
	KeepSource *bool `json:"KeepSource,omitempty"`

	// If true, ignore errors during the individual steps of the job.
	Permissive *bool `json:"Permissive,omitempty"`

	// Defines the priority of the job (Only work on async mode, which isn't implemented)
	Priority int `json:"Priority,omitempty"`

	// Transcode the DICOM instance to the provided transfersyntax (https://orthanc.uclouvain.be/book/faq/transcoding.html)
	Transcode string `json:"Transcode,omitempty"`
}

// AnonymizeResponse represents a response to anonymize a series in a synchronous mode
type PatientAnonymizeResponse struct {
	// Orthanc Series ID or the new study
	ID string `json:"ID,omitempty"`

	// Anonymized instance count
	InstancesCount int `json:"InstancesCount,omitempty"`

	// Anonymized failed instance count
	FailedInstancesCount int `json:"FailedInstancesCount,omitempty"`

	// If true, means the resource was anonymized
	IsAnonymization bool `json:"IsAnonymization,omitempty"`

	// Parent resource (Maybe deleted depending on KeepSource)
	ParentResources []string `json:"ParentResources,omitempty"`

	// Path to access the new series
	Path string `json:"Path,omitempty"`

	// Orthanc Patient ID or the new series
	PatientID string `json:"PatientID,omitempty"`

	// Type of the resource, can be "Study", "Series", "Instance" or "Patient"
	Type string `json:"Type,omitempty"`
}