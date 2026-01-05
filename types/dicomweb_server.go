package types

// DicomWebServer represents a DICOMweb server configuration
// Note: Orthanc returns boolean values as string numbers ("0" or "1")
type DicomWebServer struct {
	// URL of the remote DICOMweb server
	Url string `json:"Url"`

	// Username for authentication (optional)
	Username string `json:"Username,omitempty"`

	// Password for authentication (optional)
	Password string `json:"Password,omitempty"`

	// Whether the server supports DELETE operations ("0" = false, "1" = true)
	HasDelete string `json:"HasDelete,omitempty"`

	// Whether to use chunked transfers ("0" = false, "1" = true)
	// Set to "0" if target is Orthanc <= 1.5.6
	ChunkedTransfers string `json:"ChunkedTransfers,omitempty"`

	// Whether the server supports WADO-RS universal transfer syntax ("0" = false, "1" = true)
	// Set to "0" if target is Orthanc DICOMweb plugin <= 1.0
	HasWadoRsUniversalTransferSyntax string `json:"HasWadoRsUniversalTransferSyntax,omitempty"`
}

// DicomWebServerCreateRequest represents a request to create or update a DICOMweb server
type DicomWebServerCreateRequest struct {
	// URL of the remote DICOMweb server
	Url string `json:"Url"`

	// Username for authentication (optional)
	Username string `json:"Username,omitempty"`

	// Password for authentication (optional)
	Password string `json:"Password,omitempty"`

	// Whether the server supports DELETE operations (optional, default: false)
	HasDelete *bool `json:"HasDelete,omitempty"`

	// Whether to use chunked transfers (optional, default: true)
	// Set to false if target is Orthanc <= 1.5.6
	ChunkedTransfers *bool `json:"ChunkedTransfers,omitempty"`

	// Whether the server supports WADO-RS universal transfer syntax (optional, default: true)
	// Set to false if target is Orthanc DICOMweb plugin <= 1.0
	HasWadoRsUniversalTransferSyntax *bool `json:"HasWadoRsUniversalTransferSyntax,omitempty"`
}
