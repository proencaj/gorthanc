package types

// Study represents detailed information about a DICOM study
type Study struct {
	// Unique identifier of the study in Orthanc
	ID string `json:"ID"`

	// Whether the study is stable (no more instances expected)
	IsStable bool `json:"IsStable"`

	// Last update timestamp
	LastUpdate string `json:"LastUpdate"`

	// Main DICOM tags for the study
	MainDicomTags MainDicomTags `json:"MainDicomTags"`

	// Parent patient ID
	ParentPatient string `json:"ParentPatient"`

	// List of series IDs in this study
	Series []string `json:"Series"`

	// Type of resource (should be "Study")
	Type string `json:"Type"`

	// Patient main DICOM tags (when expanded)
	PatientMainDicomTags MainDicomTags `json:"PatientMainDicomTags,omitempty"`
}

// MainDicomTags represents the main DICOM tags for a resource
type MainDicomTags struct {
	// Study Instance UID
	StudyInstanceUID string `json:"StudyInstanceUID,omitempty"`

	// Study Description
	StudyDescription string `json:"StudyDescription,omitempty"`

	// Study Date
	StudyDate string `json:"StudyDate,omitempty"`

	// Study Time
	StudyTime string `json:"StudyTime,omitempty"`

	// Study ID
	StudyID string `json:"StudyID,omitempty"`

	// Accession Number
	AccessionNumber string `json:"AccessionNumber,omitempty"`

	// Referring Physician Name
	ReferringPhysicianName string `json:"ReferringPhysicianName,omitempty"`

	// Institution Name
	InstitutionName string `json:"InstitutionName,omitempty"`

	// Request Attributes Sequence
	RequestAttributesSequence string `json:"RequestAttributesSequence,omitempty"`

	// Patient Name (for patient-level tags)
	PatientName string `json:"PatientName,omitempty"`

	// Patient ID
	PatientID string `json:"PatientID,omitempty"`

	// Patient Birth Date
	PatientBirthDate string `json:"PatientBirthDate,omitempty"`

	// Patient Sex
	PatientSex string `json:"PatientSex,omitempty"`
}

// StudiesQueryParams represents query parameters for GET /studies
type StudiesQueryParams struct {
	// Expand the response to include full details
	Expand bool

	// Return results starting from this index
	Since int

	// Maximum number of results to return
	Limit int

	// Show DICOM tags in full format
	Full bool

	// Show DICOM tags in hexadecimal format
	Short bool
}


// ModifyRequest represents a request to modify DICOM tags
type ModifyRequest struct {
	// DICOM tags to replace with new values
	Replace map[string]string `json:"Replace,omitempty"`

	// DICOM tags to remove
	Remove []string `json:"Remove,omitempty"`

	// DICOM tags to keep unchanged
	Keep []string `json:"Keep,omitempty"`

	// Whether to keep private tags
	KeepPrivateTags *bool `json:"KeepPrivateTags,omitempty"`

	// Force operation even if it would create an invalid DICOM file
	Force *bool `json:"Force,omitempty"`
}

// AnonymizeRequest represents a request to anonymize a study
type StudyAnonymizeRequest struct {
	// If true, the REST API will return a Job ID and the job will be put in a queue
	Asynchronous *bool `json:"Asynchronous,omitempty"`

	// DicomVersion to use for anonymization
	DicomVersion string `json:"DicomVersion,omitempty"`

	// Force operation even if it would create an invalid DICOM file
	Force *bool `json:"Force,omitempty"`

	// By default orthanc does not exclude the source study, use set this to falso to delete after the anonymization proccess
	KeepSource *bool `json:"KeepSource,omitempty"`

	// If true, ignore errors during the individual steps of the job.
	Permissive *bool `json:"Permissive,omitempty"`

	// Defines the priority of the job (Only work on async mode, which isn't implemented)
	Priority int `json:"Priority,omitempty"`

	// Transcode the DICOM instance to the provided transfersyntax (https://orthanc.uclouvain.be/book/faq/transcoding.html)
	Transcode string `json:"Transcode,omitempty"`
}

// AnonymizeResponse represents a response to anonymize a study in a synchronous mode
type StudyAnonymizeResponse struct {
	// Orthanc Study ID or the new study
	ID string `json:"ID,omitempty"`

	// Anonymized instance count
	InstancesCount int `json:"InstancesCount,omitempty"`

	// Anonymized failed instance count
	FailedInstancesCount int `json:"FailedInstancesCount,omitempty"`

	// If true, means the resource was anonymized
	IsAnonymization bool `json:"IsAnonymization,omitempty"`

	// Parent resource (Maybe deleted depending on KeepSource)
	ParentResources []string `json:"ParentResources,omitempty"`

	// Path to access the new study
	Path string `json:"Path,omitempty"`

	// Orthanc Patient ID or the new study
	PatientID string `json:"PatientID,omitempty"`

	// Type of the resource, can be "Study", "Series", "Instance" or "Patient"
	Type string `json:"Type,omitempty"`
}

// ModifyResponse represents the response from modify/anonymize operations
type UploadDicomFileResponse struct {
	// ID of the newly created resource
	ID string `json:"ID"`

	// Path to the newly created resource
	Path string `json:"Path"`

	// Type of the resource (Study, Series, etc.)
	Status string `json:"Status"`

	// Type of the resource (Study, Series, etc.)
	ParentStudy string `json:"ParentStudy"`

	// Type of the resource (Study, Series, etc.)
	ParentSeries string `json:"ParientSeries"`

	// Type of the resource (Study, Series, etc.)
	ParentPatient string `json:"ParentPatient"`
}

// ModifyResponse represents the response from modify/anonymize operations
type ModifyResponse struct {
	// ID of the newly created resource
	ID string `json:"ID"`

	// Path to the newly created resource
	Path string `json:"Path"`

	// Type of the resource (Study, Series, etc.)
	Type string `json:"Type"`
}

// Statistics represents statistics about a study/series/instance
type Statistics struct {
	// Total disk size in bytes
	DiskSize string `json:"DiskSize"`

	// Disk size in megabytes
	DiskSizeMB int `json:"DiskSizeMB"`

	// Number of instances
	CountInstances int `json:"CountInstances,omitempty"`

	// Number of series
	CountSeries int `json:"CountSeries,omitempty"`

	// Uncompressed size in bytes
	UncompressedSize string `json:"UncompressedSize,omitempty"`

	// Uncompressed size in megabytes
	UncompressedSizeMB int `json:"UncompressedSizeMB,omitempty"`
}
