package types

// Instance represents detailed information about a DICOM instance
type Instance struct {
	// Unique identifier of the instance in Orthanc
	ID string `json:"ID"`

	// Main DICOM tags for the instance
	MainDicomTags InstanceDicomTags `json:"MainDicomTags"`

	// Parent series ID
	ParentSeries string `json:"ParentSeries"`

	// Type of resource (should be "Instance")
	Type string `json:"Type"`

	// File size in bytes
	FileSize int64 `json:"FileSize"`

	// File UUID
	FileUuid string `json:"FileUuid"`

	// Index in series
	IndexInSeries int `json:"IndexInSeries,omitempty"`

	// Modified timestamp
	ModifiedFrom string `json:"ModifiedFrom,omitempty"`
}

// InstanceDicomTags represents the main DICOM tags for an instance
type InstanceDicomTags struct {
	// SOP Instance UID
	SOPInstanceUID string `json:"SOPInstanceUID,omitempty"`

	// Image Index
	ImageIndex string `json:"ImageIndex,omitempty"`

	// Instance Creation Date
	InstanceCreationDate string `json:"InstanceCreationDate,omitempty"`

	// Instance Creation Time
	InstanceCreationTime string `json:"InstanceCreationTime,omitempty"`

	// Instance Number
	InstanceNumber string `json:"InstanceNumber,omitempty"`

	// Acquisition Number
	AcquisitionNumber string `json:"AcquisitionNumber,omitempty"`

	// Image Position Patient
	ImagePositionPatient string `json:"ImagePositionPatient,omitempty"`

	// Image Orientation Patient
	ImageOrientationPatient string `json:"ImageOrientationPatient,omitempty"`

	// Frame of Reference UID
	FrameOfReferenceUID string `json:"FrameOfReferenceUID,omitempty"`

	// Slice Location
	SliceLocation string `json:"SliceLocation,omitempty"`

	// Slice Thickness
	SliceThickness string `json:"SliceThickness,omitempty"`

	// Samples Per Pixel
	SamplesPerPixel string `json:"SamplesPerPixel,omitempty"`

	// Rows
	Rows string `json:"Rows,omitempty"`

	// Columns
	Columns string `json:"Columns,omitempty"`

	// Bits Allocated
	BitsAllocated string `json:"BitsAllocated,omitempty"`

	// Bits Stored
	BitsStored string `json:"BitsStored,omitempty"`

	// High Bit
	HighBit string `json:"HighBit,omitempty"`

	// Pixel Representation
	PixelRepresentation string `json:"PixelRepresentation,omitempty"`

	// Photometric Interpretation
	PhotometricInterpretation string `json:"PhotometricInterpretation,omitempty"`

	// Image Type
	ImageType string `json:"ImageType,omitempty"`

	// Rescale Intercept
	RescaleIntercept string `json:"RescaleIntercept,omitempty"`

	// Rescale Slope
	RescaleSlope string `json:"RescaleSlope,omitempty"`

	// Window Center
	WindowCenter string `json:"WindowCenter,omitempty"`

	// Window Width
	WindowWidth string `json:"WindowWidth,omitempty"`

	// Content Date
	ContentDate string `json:"ContentDate,omitempty"`

	// Content Time
	ContentTime string `json:"ContentTime,omitempty"`

	// Temporal Position Identifier
	TemporalPositionIdentifier string `json:"TemporalPositionIdentifier,omitempty"`

	// Number of Frames
	NumberOfFrames string `json:"NumberOfFrames,omitempty"`
}

// InstanceQueryParams represents query parameters for GET /instances
type InstancesQueryParams struct {
	// TODO: In future, support expand, short, full and requested-tags
	// Return results starting from this index
	Since int

	// Maximum number of results to return
	Limit int
}

// SimplifiedInstance represents a simplified DICOM instance metadata
type SimplifiedInstance struct {
	// Main DICOM tags in simplified format
	MainDicomTags map[string]string `json:"MainDicomTags,omitempty"`

	// Patient main DICOM tags
	PatientMainDicomTags map[string]string `json:"PatientMainDicomTags,omitempty"`

	// Study main DICOM tags
	StudyMainDicomTags map[string]string `json:"StudyMainDicomTags,omitempty"`

	// Series main DICOM tags
	SeriesMainDicomTags map[string]string `json:"SeriesMainDicomTags,omitempty"`
}

// InstanceHeader represents a DICOM tag header
type InstanceHeader struct {
	Name  string      `json:"Name"`
	Type  string      `json:"Type"`
	Value interface{} `json:"Value"`
}

type GetInstanceTagsQueryParams struct { 
	Short bool 

	Simplify bool 

	Whole bool
}

// AnonymizeRequest represents a request to anonymize a instance
type InstancesAnonymizeRequest struct {
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

// AnonymizeResponse represents a response to anonymize a instance in a synchronous mode
type InstanceAnonymizeResponse struct {
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