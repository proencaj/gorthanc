package types

// Modality represents a DICOM modality configuration
type Modality struct {
	// Application Entity Title (AET) of the remote modality
	AET string `json:"AET"`

	// Host/IP address of the remote modality
	Host string `json:"Host"`

	// Port number of the remote modality
	Port int `json:"Port"`

	// Manufacturer name (optional)
	Manufacturer string `json:"Manufacturer,omitempty"`

	// Whether to allow echo requests
	AllowEcho bool `json:"AllowEcho,omitempty"`

	// Whether to allow C-FIND requests
	AllowFind bool `json:"AllowFind,omitempty"`

	// Whether to allow C-GET requests
	AllowGet bool `json:"AllowGet,omitempty"`

	// Whether to allow C-MOVE requests
	AllowMove bool `json:"AllowMove,omitempty"`

	// Whether to allow C-STORE requests
	AllowStore bool `json:"AllowStore,omitempty"`

	// Timeout for DICOM operations (in seconds)
	Timeout int `json:"Timeout,omitempty"`
}

// ModalityCreateRequest represents a request to create or update a modality
type ModalityCreateRequest struct {
	// Application Entity Title (AET) of the remote modality
	AET string `json:"AET"`

	// Host/IP address of the remote modality
	Host string `json:"Host"`

	// Port number of the remote modality
	Port int `json:"Port"`

	// Manufacturer name (optional)
	Manufacturer string `json:"Manufacturer,omitempty"`

	// Whether to allow echo requests (default: true)
	AllowEcho *bool `json:"AllowEcho,omitempty"`

	// Whether to allow C-FIND requests (default: true)
	AllowFind *bool `json:"AllowFind,omitempty"`

	// Whether to allow C-GET requests (default: true)
	AllowGet *bool `json:"AllowGet,omitempty"`

	// Whether to allow C-MOVE requests (default: true)
	AllowMove *bool `json:"AllowMove,omitempty"`

	// Whether to allow C-STORE requests (default: true)
	AllowStore *bool `json:"AllowStore,omitempty"`

	// Timeout for DICOM operations in seconds (default: 10)
	Timeout int `json:"Timeout,omitempty"`
}

// ModalityEchoResult represents the result of a C-ECHO operation
type ModalityEchoResult struct {
	// Whether the echo was successful
	Success bool `json:"Success"`

	// Error message if unsuccessful
	ErrorMessage string `json:"ErrorMessage,omitempty"`

	// Response time in milliseconds
	ResponseTime int `json:"ResponseTime,omitempty"`
}

// ModalityStoreRequest represents a request to store resources to a modality
type ModalityStoreRequest struct {
	// List of resource IDs to send
	Resources []string `json:"Resources,omitempty"`

	// Whether to synchronously wait for the transfer
	Synchronous bool `json:"Synchronous,omitempty"`

	// Local AET to use for the transfer
	LocalAet string `json:"LocalAet,omitempty"`

	// Remote AET (if different from modality's configured AET)
	RemoteAet string `json:"RemoteAet,omitempty"`

	// Timeout in seconds
	Timeout int `json:"Timeout,omitempty"`

	// Move Originator AET (for C-MOVE operations)
	MoveOriginatorAet string `json:"MoveOriginatorAet,omitempty"`

	// Move Originator ID (for C-MOVE operations)
	MoveOriginatorID int `json:"MoveOriginatorID,omitempty"`

	Permissive int `json:"Permissive,omitempty"`

	StorageCommitment int `json:"StorageCommitment,omitempty"`
}

// ModalityStoreResult represents the result of a C-STORE operation
type ModalityStoreResult struct {
	// Description of the operation
	Description string `json:"Description"`

	// Local AET used
	LocalAet string `json:"LocalAet"`

	// Remote AET used
	RemoteAet string `json:"RemoteAet"`

	// Parent resources sent
	ParentResources []string `json:"ParentResources,omitempty"`

	// Number of instances sent
	InstancesCount int `json:"InstancesCount,omitempty"`

	// Number of failed instances
	FailedInstancesCount int `json:"FailedInstancesCount,omitempty"`
}

// ModalityFindRequest represents a C-FIND query request
type ModalityFindRequest struct {
	// Level of the query (Patient, Study, Series, Instance)
	Level string `json:"Level"`

	// Query parameters (DICOM tags and values)
	Query map[string]string `json:"Query"`

	// Whether to normalize the query
	Normalize bool `json:"Normalize,omitempty"`

	// Timeout in seconds
	Timeout int `json:"Timeout,omitempty"`
}

// ModalityFindResult represents a single result from a C-FIND query
// This is a flexible map that contains DICOM tags returned by the query
type ModalityFindResult map[string]interface{}

// ModalityQueryResponse represents the response when creating a C-FIND query
type ModalityQueryResponse struct {
	// Query ID that can be used to retrieve results
	ID string `json:"ID"`

	// Path to the query resource
	Path string `json:"Path"`
}

// ModalityMoveRequest represents a C-MOVE request
type ModalityMoveRequest struct {
	// Level of the query (Patient, Study, Series, Instance)
	Level string `json:"Level"`

	Limit int `json:"Limit,omitempty"`

	// Target AET where resources should be moved
	TargetAet string `json:"TargetAet,omitempty"`

	// Resource ID to move (alternative to Query)
	// Resources []string `json:"Resources,omitempty"`
	Resources []map[string]interface{} `json:"Resources,omitempty"`

	Priority int `json:"Priority,omitempty"`

	Permissive bool `json:"Permissive,omitempty"`
	
	Asynchronous bool `json:"Asynchronous,omitempty"`

	// Timeout in seconds
	Timeout int `json:"Timeout,omitempty"`
}

// ModalityMoveResult represents the result of a C-MOVE operation
type ModalityMoveResult struct {
	// Description of the operation
	Description string `json:"Description"`

	// Local AET used
	LocalAet string `json:"LocalAet"`

	// Remote AET used
	RemoteAet string `json:"RemoteAet"`

	// Target AET where resources are being moved
	TargetAet string `json:"TargetAet"`

	// Query parameters used (array of DICOM tag/value pairs)
	Query []map[string]string `json:"Query,omitempty"`
}

// ModalityGetRequest represents a C-GET request
type ModalityGetRequest struct {
	// Level of the query (Patient, Study, Series, Instance)
	Level string `json:"Level"`

	// Resource ID to get (alternative to Query)
	Resources []map[string]interface{} `json:"Resources,omitempty"`

	// Timeout in seconds
	Timeout int `json:"Timeout,omitempty"`

	// If true, ignore errors during the individual steps of the job.
	Permissive bool `json:"Permissive,omitempty"`
	
	// If true, run the job in asynchronous mode,
	Asynchronous bool `json:"Asynchronous,omitempty"`
}
