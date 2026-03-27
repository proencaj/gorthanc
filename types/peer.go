package types

import "encoding/json"

// Peer represents an Orthanc peer configuration
type Peer struct {
	// URL of the remote Orthanc peer
	URL string `json:"Url"`

	// Username for authentication
	Username string `json:"Username,omitempty"`

	// Password for authentication (null when redacted by Orthanc)
	Password *string `json:"Password,omitempty"`

	// Whether PKCS#11 is enabled
	Pkcs11 bool `json:"Pkcs11"`

	// Connection timeout in seconds
	Timeout int `json:"Timeout"`

	// Custom HTTP headers for the connection.
	// Orthanc returns [] when empty and {} when populated.
	HttpHeaders FlexibleMap `json:"HttpHeaders"`
}

// FlexibleMap handles JSON fields that can be either an empty array [] or an object {}.
// Orthanc returns [] for empty maps in some endpoints.
type FlexibleMap map[string]string

func (f *FlexibleMap) UnmarshalJSON(data []byte) error {
	// Handle empty array
	if string(data) == "[]" || string(data) == "null" {
		*f = nil
		return nil
	}

	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*f = m
	return nil
}

// PeerCreateRequest represents a request to create or update an Orthanc peer
type PeerCreateRequest struct {
	// URL of the root of the REST API of the remote Orthanc peer
	URL string `json:"Url"`

	// Username for authentication
	Username string `json:"Username,omitempty"`

	// Password for authentication
	Password string `json:"Password,omitempty"`

	// SSL certificate file path
	CertificateFile string `json:"CertificateFile,omitempty"`

	// SSL certificate key file path
	CertificateKeyFile string `json:"CertificateKeyFile,omitempty"`

	// SSL certificate key password
	CertificateKeyPassword string `json:"CertificateKeyPassword,omitempty"`

	// Custom HTTP headers for the connection
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// PeerStoreRequest represents a request to send resources to an Orthanc peer
type PeerStoreRequest struct {
	// List of Orthanc identifiers of DICOM resources to send
	Resources []string `json:"Resources"`

	// Whether to compress DICOM instances using gzip before sending
	Compress *bool `json:"Compress,omitempty"`

	// Whether to run synchronously (default: true)
	Synchronous *bool `json:"Synchronous,omitempty"`

	// Whether to ignore errors during individual steps
	Permissive *bool `json:"Permissive,omitempty"`

	// Job priority in asynchronous mode (higher = more priority)
	Priority int `json:"Priority,omitempty"`

	// Transcode to this transfer syntax before sending
	Transcode string `json:"Transcode,omitempty"`
}

// PeerStoreResult represents the result of sending resources to an Orthanc peer
type PeerStoreResult struct {
	// Description of the operation
	Description string `json:"Description"`

	// Parent resources sent
	ParentResources []string `json:"ParentResources,omitempty"`

	// Number of instances sent
	InstancesCount int `json:"InstancesCount,omitempty"`

	// Number of failed instances
	FailedInstancesCount int `json:"FailedInstancesCount,omitempty"`
}
