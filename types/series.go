package types

// Series represents detailed information about a DICOM series
type Series struct {
	// Unique identifier of the series in Orthanc
	ID string `json:"ID"`

	// Whether the series is stable (no more instances expected)
	IsStable bool `json:"IsStable"`

	// Last update timestamp
	LastUpdate string `json:"LastUpdate"`

	// Main DICOM tags for the series
	MainDicomTags SeriesDicomTags `json:"MainDicomTags"`

	// Parent study ID
	ParentStudy string `json:"ParentStudy"`

	// List of instance IDs in this series
	Instances []string `json:"Instances"`

	// Type of resource (should be "Series")
	Type string `json:"Type"`

	// Expected number of instances (if known)
	ExpectedNumberOfInstances int `json:"ExpectedNumberOfInstances,omitempty"`

	// Status of the series
	Status string `json:"Status,omitempty"`
}

// SeriesDicomTags represents the main DICOM tags for a series
type SeriesDicomTags struct {
	// Series Instance UID
	SeriesInstanceUID string `json:"SeriesInstanceUID,omitempty"`

	// Series Description
	SeriesDescription string `json:"SeriesDescription,omitempty"`

	// Series Number
	SeriesNumber string `json:"SeriesNumber,omitempty"`

	// Series Date
	SeriesDate string `json:"SeriesDate,omitempty"`

	// Series Time
	SeriesTime string `json:"SeriesTime,omitempty"`

	// Modality (e.g., CT, MR, US, XR)
	Modality string `json:"Modality,omitempty"`

	// Body Part Examined
	BodyPartExamined string `json:"BodyPartExamined,omitempty"`

	// Protocol Name
	ProtocolName string `json:"ProtocolName,omitempty"`

	// Sequence Name
	SequenceName string `json:"SequenceName,omitempty"`

	// Cardiac Number of Images
	CardiacNumberOfImages string `json:"CardiacNumberOfImages,omitempty"`

	// Images in Acquisition
	ImagesInAcquisition string `json:"ImagesInAcquisition,omitempty"`

	// Number of Temporal Positions
	NumberOfTemporalPositions string `json:"NumberOfTemporalPositions,omitempty"`

	// Number of Slices
	NumberOfSlices string `json:"NumberOfSlices,omitempty"`

	// Number of Time Slices
	NumberOfTimeSlices string `json:"NumberOfTimeSlices,omitempty"`

	// Image Orientation Patient
	ImageOrientationPatient string `json:"ImageOrientationPatient,omitempty"`

	// Series Type
	SeriesType string `json:"SeriesType,omitempty"`

	// Acquisition Number
	AcquisitionNumber string `json:"AcquisitionNumber,omitempty"`

	// Contrast/Bolus Agent
	ContrastBolusAgent string `json:"ContrastBolusAgent,omitempty"`

	// Scanner Make
	Manufacturer string `json:"Manufacturer,omitempty"`

	// Operator's Name
	OperatorsName string `json:"OperatorsName,omitempty"`

	// Performed Procedure Step Description
	PerformedProcedureStepDescription string `json:"PerformedProcedureStepDescription,omitempty"`
}

// SeriesQueryParams represents query parameters for GET /series
type SeriesQueryParams struct {
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

// AnonymizeRequest represents a request to anonymize a series
type SeriesAnonymizeRequest struct {
	// If true, the REST API will return a Job ID and the job will be put in a queue
	Asynchronous bool `json:"Asynchronous,omitempty"`

	// DicomVersion to use for anonymization
	DicomVersion string `json:"DicomVersion,omitempty"`

	// Force operation even if it would create an invalid DICOM file
	Force bool `json:"Force,omitempty"`	

	// By default orthanc does not exclude the source series, use set this to falso to delete after the anonymization proccess
	KeepSource bool `json:"KeepSource,omitempty"`	

	// If true, ignore errors during the individual steps of the job.
	Permissive bool `json:"Permissive,omitempty"`

	// Defines the priority of the job (Only work on async mode, which isn't implemented)
	Priority int `json:"Priority,omitempty"`

	// Transcode the DICOM instance to the provided transfersyntax (https://orthanc.uclouvain.be/book/faq/transcoding.html)
	Transcode string `json:"Transcode,omitempty"`
}

// AnonymizeResponse represents a response to anonymize a series in a synchronous mode
type SeriesAnonymizeResponse struct {
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