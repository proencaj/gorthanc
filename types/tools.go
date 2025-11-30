package types

// ResourceLevel represents the level of DICOM resources
type ResourceLevel string

const (
	ResourceLevelPatient  ResourceLevel = "Patient"
	ResourceLevelStudy    ResourceLevel = "Study"
	ResourceLevelSeries   ResourceLevel = "Series"
	ResourceLevelInstance ResourceLevel = "Instance"
)

// LogLevel represents the logging verbosity level in Orthanc
type LogLevel string

const (
	// LogLevelDefault shows only WARNING and ERROR messages
	LogLevelDefault LogLevel = "default"

	// LogLevelVerbose adds INFO level messages
	LogLevelVerbose LogLevel = "verbose"

	// LogLevelTrace includes detailed TRACE level messages for debugging
	LogLevelTrace LogLevel = "trace"
)

// ToolsFindRequest represents a request to find DICOM resources
type ToolsFindRequest struct {
	// Level specifies the resource level (Patient, Study, Series, Instance)
	Level ResourceLevel `json:"Level"`

	// Query contains DICOM tag criteria for searching
	Query map[string]string `json:"Query"`

	// Expand returns expanded information about the resources (optional)
	Expand *bool `json:"Expand,omitempty"`

	// Limit restricts the number of results (optional)
	Limit *int `json:"Limit,omitempty"`

	// Since returns results starting from this index (optional, available in newer versions)
	Since *int `json:"Since,omitempty"`

	// RequestedTags specifies which tags to include in the response (optional, Orthanc 1.11.0+)
	RequestedTags []string `json:"RequestedTags,omitempty"`

	// Labels filters resources by labels (optional, Orthanc 1.12.0+)
	Labels []string `json:"Labels,omitempty"`

	// LabelsConstraint specifies how to apply label filters (optional, Orthanc 1.12.0+)
	LabelsConstraint string `json:"LabelsConstraint,omitempty"`

	// OrderBy specifies result ordering (optional, Orthanc 1.12.5+)
	OrderBy []OrderByEntry `json:"OrderBy,omitempty"`
}

// OrderByEntry represents an ordering criterion
type OrderByEntry struct {
	// Type specifies what to order by (e.g., "DicomTag", "Metadata")
	Type string `json:"Type"`

	// Key specifies the tag or metadata key
	Key string `json:"Key,omitempty"`

	// Direction specifies sort direction ("ASC" or "DESC")
	Direction string `json:"Direction,omitempty"`
}

// ToolsFindExpandedResource represents an expanded resource from /tools/find
type ToolsFindExpandedResource struct {
	// ID is the Orthanc identifier of the resource
	ID string `json:"ID"`

	// Type is the resource type (Patient, Study, Series, Instance)
	Type string `json:"Type"`

	// IsStable indicates if the resource is stable
	IsStable bool `json:"IsStable,omitempty"`

	// LastUpdate is the last update timestamp
	LastUpdate string `json:"LastUpdate,omitempty"`

	// MainDicomTags contains the main DICOM tags
	MainDicomTags map[string]interface{} `json:"MainDicomTags,omitempty"`

	// PatientMainDicomTags contains patient-level DICOM tags
	PatientMainDicomTags map[string]interface{} `json:"PatientMainDicomTags,omitempty"`

	// Labels contains resource labels (Orthanc 1.12.0+)
	Labels []string `json:"Labels,omitempty"`

	// ParentPatient is the parent patient ID (for Study/Series/Instance)
	ParentPatient string `json:"ParentPatient,omitempty"`

	// ParentStudy is the parent study ID (for Series/Instance)
	ParentStudy string `json:"ParentStudy,omitempty"`

	// ParentSeries is the parent series ID (for Instance)
	ParentSeries string `json:"ParentSeries,omitempty"`

	// Studies contains child study IDs (for Patient)
	Studies []string `json:"Studies,omitempty"`

	// Series contains child series IDs (for Study)
	Series []string `json:"Series,omitempty"`

	// Instances contains child instance IDs (for Series)
	Instances []string `json:"Instances,omitempty"`

	// RequestedTags contains requested DICOM tags (Orthanc 1.11.0+)
	RequestedTags map[string]interface{} `json:"RequestedTags,omitempty"`
}
