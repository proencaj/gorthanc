package types

// DicomWebStudy represents a study returned by QIDO-RS
type DicomWebStudy struct {
	// Study Instance UID (0020,000D)
	StudyInstanceUID string `json:"0020000D,omitempty"`
	// Study Date (0008,0020)
	StudyDate string `json:"00080020,omitempty"`
	// Study Time (0008,0030)
	StudyTime string `json:"00080030,omitempty"`
	// Accession Number (0008,0050)
	AccessionNumber string `json:"00080050,omitempty"`
	// Modalities in Study (0008,0061)
	ModalitiesInStudy string `json:"00080061,omitempty"`
	// Referring Physician's Name (0008,0090)
	ReferringPhysicianName string `json:"00080090,omitempty"`
	// Patient's Name (0010,0010)
	PatientName string `json:"00100010,omitempty"`
	// Patient ID (0010,0020)
	PatientID string `json:"00100020,omitempty"`
	// Study ID (0020,0010)
	StudyID string `json:"00200010,omitempty"`
	// Number of Study Related Series (0020,1206)
	NumberOfStudyRelatedSeries string `json:"00201206,omitempty"`
	// Number of Study Related Instances (0020,1208)
	NumberOfStudyRelatedInstances string `json:"00201208,omitempty"`
}

// DicomWebSeries represents a series returned by QIDO-RS
type DicomWebSeries struct {
	// Series Instance UID (0020,000E)
	SeriesInstanceUID string `json:"0020000E,omitempty"`
	// Modality (0008,0060)
	Modality string `json:"00080060,omitempty"`
	// Series Number (0020,0011)
	SeriesNumber string `json:"00200011,omitempty"`
	// Series Description (0008,103E)
	SeriesDescription string `json:"0008103E,omitempty"`
	// Number of Series Related Instances (0020,1209)
	NumberOfSeriesRelatedInstances string `json:"00201209,omitempty"`
	// Performed Procedure Step Start Date (0040,0244)
	PerformedProcedureStepStartDate string `json:"00400244,omitempty"`
	// Performed Procedure Step Start Time (0040,0245)
	PerformedProcedureStepStartTime string `json:"00400245,omitempty"`
}

// DicomWebInstance represents an instance returned by QIDO-RS
type DicomWebInstance struct {
	// SOP Class UID (0008,0016)
	SOPClassUID string `json:"00080016,omitempty"`
	// SOP Instance UID (0008,0018)
	SOPInstanceUID string `json:"00080018,omitempty"`
	// Instance Number (0020,0013)
	InstanceNumber string `json:"00200013,omitempty"`
	// Rows (0028,0010)
	Rows string `json:"00280010,omitempty"`
	// Columns (0028,0011)
	Columns string `json:"00280011,omitempty"`
	// Bits Allocated (0028,0100)
	BitsAllocated string `json:"00280100,omitempty"`
	// Number of Frames (0028,0008)
	NumberOfFrames string `json:"00280008,omitempty"`
}

// QidoQueryParams represents common query parameters for QIDO-RS requests
type QidoQueryParams struct {
	// Limit the number of results
	Limit int
	// Skip the first N results
	Offset int
	// Include all DICOM fields in the response
	Includefield string
	// Filter by fuzzy matching
	FuzzyMatching bool
}

// QidoStudyQueryParams represents query parameters for study-level QIDO-RS
type QidoStudyQueryParams struct {
	QidoQueryParams
	// Filter by Study Instance UID
	StudyInstanceUID string
	// Filter by Patient ID
	PatientID string
	// Filter by Patient Name
	PatientName string
	// Filter by Accession Number
	AccessionNumber string
	// Filter by Study Date (YYYYMMDD or range YYYYMMDD-YYYYMMDD)
	StudyDate string
	// Filter by Modalities in Study
	ModalitiesInStudy string
}

// QidoSeriesQueryParams represents query parameters for series-level QIDO-RS
type QidoSeriesQueryParams struct {
	QidoQueryParams
	// Filter by Study Instance UID (required for study-level series query)
	StudyInstanceUID string
	// Filter by Series Instance UID
	SeriesInstanceUID string
	// Filter by Modality
	Modality string
	// Filter by Series Number
	SeriesNumber string
}

// QidoInstanceQueryParams represents query parameters for instance-level QIDO-RS
type QidoInstanceQueryParams struct {
	QidoQueryParams
	// Filter by Study Instance UID
	StudyInstanceUID string
	// Filter by Series Instance UID
	SeriesInstanceUID string
	// Filter by SOP Instance UID
	SOPInstanceUID string
	// Filter by SOP Class UID
	SOPClassUID string
}

// WadoUriParams represents query parameters for WADO-URI requests
type WadoUriParams struct {
	// Request type, must be "WADO"
	RequestType string
	// Study Instance UID (required)
	StudyUID string
	// Series Instance UID (required)
	SeriesUID string
	// SOP Instance UID (required)
	ObjectUID string
	// Content type (e.g., "application/dicom", "image/jpeg", "image/png")
	ContentType string
	// Transfer syntax for DICOM objects
	TransferSyntax string
	// Anonymize the response
	Anonymize string
	// Frame number for multi-frame objects
	FrameNumber int
	// Image quality (1-100) for lossy compression
	ImageQuality int
	// Window center for image rendering
	WindowCenter string
	// Window width for image rendering
	WindowWidth string
	// Rows for image scaling
	Rows int
	// Columns for image scaling
	Columns int
	// Region to extract (format: xmin,ymin,xmax,ymax as fractions 0.0-1.0)
	Region string
}

// WadoRsRenderedParams represents parameters for WADO-RS rendered endpoints
type WadoRsRenderedParams struct {
	// Accept header value (e.g., "image/jpeg", "image/png")
	Accept string
	// Window center for image rendering
	WindowCenter string
	// Window width for image rendering
	WindowWidth string
	// Quality (1-100) for lossy compression
	Quality int
	// Viewport size (format: rows,columns)
	Viewport string
}
