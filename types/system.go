package types

// SystemInfo represents the response from GET /system
type SystemInfo struct {
	// API version (e.g., "1.12.0")
	ApiVersion int `json:"ApiVersion"`

	// Version of the database schema
	DatabaseVersion int `json:"DatabaseVersion"`

	// Database backend plugin name (if any)
	DatabaseBackendPlugin string `json:"DatabaseBackendPlugin,omitempty"`

	// DICOM Application Entity Title
	DicomAet string `json:"DicomAet"`

	// DICOM port number
	DicomPort int `json:"DicomPort"`

	// Whether revision checking is enabled
	CheckRevisions bool `json:"CheckRevisions"`

	// HTTP port number
	HttpPort int `json:"HttpPort"`

	// Server name
	Name string `json:"Name"`

	// Whether plugins are enabled
	PluginsEnabled bool `json:"PluginsEnabled"`

	// Storage area plugin name (if any)
	StorageAreaPlugin string `json:"StorageAreaPlugin,omitempty"`

	// Version of Orthanc server
	Version string `json:"Version"`

	// In-memory database identifier (if applicable)
	InMemoryDatabaseIdentifier string `json:"InMemoryDatabaseIdentifier,omitempty"`

	// Maximum storage size in bytes (0 = unlimited)
	MaximumStorageSize int64 `json:"MaximumStorageSize,omitempty"`

	// Maximum number of patients (0 = unlimited)
	MaximumPatients int `json:"MaximumPatients,omitempty"`

	// Whether storage compression is enabled
	StorageCompression *bool `json:"StorageCompression,omitempty"`

	// Whether instance overwriting is enabled
	OverwriteInstances *bool `json:"OverwriteInstances,omitempty"`
}

type SystemStatistics struct {
	// Total number of DICOM instances
	CountInstances int `json:"CountInstances"`

	// Total number of unique patients
	CountPatients int `json:"CountPatients"`

	// Total number of series across all studies
	CountSeries int `json:"CountSeries"`

	// Total number of clinical studies
	CountStudies int `json:"CountStudies"`

	// Size on disk in bytes
	TotalDiskSize string `json:"TotalDiskSize"`

	// Size on disk converted to Megabytes
	TotalDiskSizeMB int `json:"TotalDiskSizeMB"`

	// Full size without compression in bytes
	TotalUncompressedSize string `json:"TotalUncompressedSize"`

	// Uncompressed size converted to Megabytes
	TotalUncompressedSizeMB int `json:"TotalUncompressedSizeMB"`
}

